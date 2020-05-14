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
func (t *tagRepository) Get(id int) error {
	return global.DB.First(&models.Tag{}, id).Error
}

// 删除标签
func (t *tagRepository) Delete(id int) error {
	return global.DB.Where("id = ?", id).Delete(&models.Tag{}).Error
}

func (t *tagRepository) GetTags(offset, pageSize int, tagName string) (tags []models.Tag) {
	query := global.DB.Where("tag_status = ?", models.TAG_STATUS_NORMAL)
	//pageSize := com.StrTo(os.Getenv("PAGESIZE")).MustInt()
	//offset := (page - 1) * pageSize

	// 搜索
	if tagName != "" {
		query = query.Where("name like ?", "%"+tagName+"%")
	}

	// 获取tag数据
	query.Offset(offset).Limit(pageSize).Find(&tags)
	return

	// 获取tag总数
	//err2 := query.Model(&Tag{}).Count(&count).Error
	//if err1 != nil || err2 != nil {
	//	return []Tag{}, 0
	//}
	//return
}