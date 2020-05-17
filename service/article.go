package service

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"go-gin-blog-api/models"
	"go-gin-blog-api/models/form_validate"
	"go-gin-blog-api/models/read_model"
	"go-gin-blog-api/repository"
	"go-gin-blog-api/response"
)

var Article = newArticleService()

func newArticleService() *articleService {
	return &articleService{}
}

type articleService struct {
}

// 添加文章（标签信息关联插入)
func (t *articleService) Create(articleCreateData form_validate.ArticleCreateForm) int {
	// 文章信息
	var article models.Article
	article.SetTitle(articleCreateData.Title)
	article.SetContent(articleCreateData.Content)
	article.SetCurrentStatus(articleCreateData.CurrentStatus)
	article.SetTags(articleCreateData.TagIds)
	article.SetAttachments(articleCreateData.Attachments, 0)
	err := repository.Article.Create(&article)
	if err != nil {
		return response.ARTICLE_CREATE_FAILED
	}
	return response.Ok;
}


// 文章更新
func (t *articleService) Update(articleUpdateData form_validate.ArticleUpdateForm) int {
	// 密等
	article := repository.Article.Get(articleUpdateData.Id)
	if article.Id == 0 {
		return response.ARTICLE_NOT_FOUND
	}

	// 文章信息
	article.SetTitle(articleUpdateData.Title)
	article.SetContent(articleUpdateData.Content)
	article.SetCurrentStatus(articleUpdateData.CurrentStatus)
	article.SetTags(articleUpdateData.TagIds)
	article.SetAttachments(articleUpdateData.Attachments, articleUpdateData.Id)

	// 替换文章与标签的对应关系
	err := repository.Article.ReplaceAssocTags(&article, article.GetTags())
	if err != nil {
		return response.ARTICLE_REPLACE_ASSOC_TAG_FALIED
	}

	// 替换文章与附件的对应关系ReplaceAssocAttachment
	logrus.Errorf("info: %v", article.ArticleAttachments)
	err = repository.Article.DeleteAssocAttachment(&article)
	fmt.Println("错误", err.Error())
	if err != nil {
		return response.ARTICLE_REPLACE_ASSOCIATION_ATTACHMENT_FAILED
	}

	err = repository.Article.Update(&article)
	if err != nil {
		return response.ARTICLE_UPDATE_FAILED
	}
	return response.Ok;
}

// 标签删除
func (t *articleService) Delete(articleId int) int {
	// 密等
	article := repository.Article.Get(articleId)
	if (article.Id == 0) {
		return response.ARTICLE_NOT_FOUND
	}

	// 清空文章与标签的级联关系(也就是该文章在那个中间表的数据)
	err := repository.Article.ClearAssocTags(article)
	if err != nil {
		return response.ARTICLE_CLEAR_ASSOCIATION_FAILED
	}

	// 删除文章
	err = repository.Article.Delete(article)
	if err != nil {
		return response.ARTICLE_DELETE_FAILED
	}
	return response.Ok;
}

// 文章列表
func (t *articleService) ArticleList(page, pageSize int, content string) ([]read_model.Article, int) {
	offset := (page - 1) * pageSize
	params := map[string]interface{}{
		"offset": offset,
		"pageSize": pageSize,
		"content": content,
	}
	tags, total := repository.Article.ArticleList(params)
	return tags, total
}
