package repository

import (
	"fmt"
	"go-gin-blog-api/global"
	"go-gin-blog-api/models"
	"go-gin-blog-api/models/read_model"
)

var Article = newArticleRepository()

func newArticleRepository() *articleRepository {
	return &articleRepository{}
}

type articleRepository struct {
}

// 文章创建
func (t *articleRepository) Create(article *models.Article) error {
	fmt.Printf("%v", article.ArticleAttachments)
	return global.DB.Create(&article).Error
}

// 文章更新
func (t *articleRepository) Update(article *models.Article) error {
	return global.DB.Save(article).Error
}

// 替换标签级联关系
func (t *articleRepository) ReplaceAssocTags(article *models.Article, tags []models.Tag) error {
	return global.DB.Model(&article).Association("Tags").Replace(tags).Error
}

// 删除文章附件级联关系
func (t *articleRepository) DeleteAssocAttachment(articleId int) error {
	return global.DB.Delete(models.ArticleAttachment{}, "article_id = ?", articleId).Error
	//return global.DB.Model(&article).Association("ArticleAttachments").Replace(&article.ArticleAttachments).Error
}

// 根据主键获取标签
func (t *articleRepository) Get(id int) (article models.Article){
	global.DB.First(&article, id).Scan(&article)
	return
}

// 删除标签
func (t *articleRepository) Delete(article models.Article) error {
	return global.DB.Delete(&article).Error
}

// 清空级联关系
func (t *articleRepository) ClearAssocTags(article models.Article) error {
	return global.DB.Model(&article).Association("Tags").Clear().Error
}

// 文章列表
func (t *articleRepository) ArticleList(params map[string]interface{}) ([]read_model.Article, int) {
	var articles []read_model.Article
	global.DB.Model(articles).
		Preload("Tags").
		Preload("ArticleAttachments").
		Where("current_status = ?", models.ARTICLE_STATUS_NORMAL).
		Offset(params["offset"]).
		Limit(params["pageSize"]).
		Order("created_at desc").
		Order("id asc").
		Find(&articles)

	var total int
	global.DB.Model(articles).Where("current_status = ?", models.ARTICLE_STATUS_NORMAL).Model(&read_model.Article{}).Count(&total)
	return articles, total
}
