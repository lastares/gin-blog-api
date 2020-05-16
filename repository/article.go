package repository

import (
	"go-gin-blog-api/global"
	"go-gin-blog-api/models"
	"go-gin-blog-api/models/const_type"
	"go-gin-blog-api/models/read_model"
)

var Article = newArticleRepository()

func newArticleRepository() *articleRepository {
	return &articleRepository{}
}

type articleRepository struct {
}

// 文章创建
func (t *articleRepository) ArticleCreate(article *models.Article) error {
	return global.DB.Set("gorm:association_autoupdate", false).Create(&article).Error
}

// 标签更新
//func (t *articleRepository) Update(tag *models.Tag) error {
//	return global.DB.Save(tag).Error
//}
//
//// 根据主键获取标签
//func (t *articleRepository) Get(id int) error {
//	return global.DB.First(&models.Tag{}, id).Error
//}
//
//// 删除标签
//func (t *articleRepository) Delete(id int) error {
//	return global.DB.Where("id = ?", id).Delete(&models.Tag{}).Error
//}
//
func (t *articleRepository) ArticleList(params map[string]interface{}) ([]read_model.Article, int) {
	var articles []read_model.Article
	global.DB.Model(articles).
		Preload("Tags").
		Where("current_status = ?", const_type.ARTICLE_STATUS_NORMAL).
		Offset(params["offset"]).
		Limit(params["pageSize"]).
		Order("created_at desc").
		Order("id asc").
		Find(&articles)

	var total int
	global.DB.Model(articles).Where("current_status = ?", const_type.ARTICLE_STATUS_NORMAL).Model(&read_model.Article{}).Count(&total)
	//query := global.DB.Where("current_status = ?", const_type.ARTICLE_STATUS_NORMAL)
	//
	//// 搜索
	////if tagName != "" {
	////	query = query.Where("tag_name like ?", "%"+tagName+"%")
	////}
	//
	//// 获取tag数据
	//var articles []models.Article
	//query.Offset(params["offset"]).Limit(params["pageSize"]).Preload("Tags").Find(&articles)
	//
	//// 获取tag总数
	//var total int
	//query.Model(&models.Article{}).Count(&total)
	return articles, total
}
