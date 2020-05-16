package service

import (
	"fmt"
	"go-gin-blog-api/models"
	"go-gin-blog-api/models/read_model"
	"go-gin-blog-api/repository"
	"go-gin-blog-api/response"
	"net/http"
)

var Article = newArticleService()

func newArticleService() *articleService {
	return &articleService{}
}

type articleService struct {
}

// 添加文章（标签信息关联插入)
func (t *articleService) CreateArticle(title, content string, currentStatus int, tagIds []int) int {
	// 整理标签ID
	var articleTags []models.Tag
	for _, tagId := range tagIds {
		fmt.Println("tag_id: ", tagId)
		articleTags = append(articleTags, models.Tag{Id: tagId})
	}

	// 文章信息
	article := &models.Article{
		Title: title,
		Content: content,
		CurrentStatus: currentStatus,
		Tags: articleTags,
	}

	err := repository.Article.ArticleCreate(article)
	if err != nil {
		return response.ARTICLE_CREATED_FAILED
	}
	return http.StatusOK;
}


// 标签更新
//func (t *tagService) Update(tagId, tagStatus int, tagName string) int {
//	// 密等
//	getError := repository.Tag.Get(tagId)
//	if getError != nil {
//		return response.TAG_NOT_FOUND
//	}
//
//	// 更新
//	tag := &models.Tag{
//		Id: tagId,
//		TagName: tagName,
//		TagStatus: tagStatus,
//	}
//	err := repository.Tag.Update(tag)
//	if err != nil {
//		return response.TAG_UPDATE_FAILED
//	}
//	return response.Ok;
//}

// 标签删除
//func (t *tagService) Delete(tagId int) int {
//	// 密等
//	getError := repository.Tag.Get(tagId)
//	if getError != nil {
//		return response.TAG_NOT_FOUND
//	}
//
//	// 删除
//	err := repository.Tag.Delete(tagId)
//	if err != nil {
//		return response.TAG_DELETE_FAILED
//	}
//	return response.Ok;
//}

// 标签列表
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
