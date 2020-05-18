package repository

import (
	"go-gin-blog-api/global"
	"go-gin-blog-api/models"
)

var Tag = newTagRepository()

func newTagRepository() *tagRepository {
	return &tagRepository{}
}

type tagRepository struct {
}

// 标签创建
func (t *tagRepository) Create(tag *models.Tag) error {
	return global.DB.Create(&tag).Error
}

// 标签更新
func (t *tagRepository) Update(tag *models.Tag) error {
	return global.DB.Save(tag).Error
}

// 根据主键获取标签
func (t *tagRepository) Get(id int) (tag models.Tag) {
	global.DB.First(&tag, id).Scan(&tag)
	return
	//return global.DB.First(&models.Tag{}, id).Error
}

// 删除标签
func (t *tagRepository) Delete(id int) error {
	return global.DB.Where("id = ?", id).Delete(&models.Tag{}).Error
}

func (t *tagRepository) GetTags(offset, pageSize int, tagName string) (tags []models.Tag, total int) {
	query := global.DB.Where("tag_status = ?", models.TAG_STATUS_NORMAL)

	// 搜索
	if tagName != "" {
		query = query.Where("tag_name like ?", "%"+tagName+"%")
	}

	// 获取tag数据
	//var tags []models.Tag
	query.Offset(offset).Limit(pageSize).Find(&tags)

	// 获取tag总数
	//var total int
	query.Model(&models.Tag{}).Count(&total)
	return
}