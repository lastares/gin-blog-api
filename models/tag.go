package models

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/unknwon/com"
	"go-gin-blog-api/global"
	"go-gin-blog-api/util"
	"os"
)

// Tag example
type Tag struct {
	ID        int           `json:"id" uri:"id" gglobal:"primary_key;AUTO_INCREMENT" label:"主键" validate:"required,gt=0" example:"10"`
	TagName   string        `valid:"email" json:"tagName" gglobal:"type:varchar(32);not null;default:\"\"" example:"php"`
	TagStatus int           `json:"tagStatus" gglobal:"type:tinyint(4);not null;default:1" example:"10"`
	CreatedAt util.JSONTime `json:"createdAt" gglobal:"type:datetime;null"`
	UpdatedAt util.JSONTime `json:"updatedAt" gglobal:"type:datetime;null"`
}

//func (tag *Tag) GetValidateError(name string) string {
//	val := err.Type.Name()
//	fmt.Println("val: ", val)
//	if val == "ID" {
//		switch val {
//		case "gt":
//			return "请输入手机号码"
//		}
//	}
//	return "参数错误"
//}

func GetTags(page int, tagName string, tagStatus int) (tags []Tag, count int) {
	query := global.DB.Where("tag_status = ?", tagStatus)
	pageSize := com.StrTo(os.Getenv("PAGESIZE")).MustInt()
	offset := (page - 1) * pageSize

	// 搜索
	if tagName != "" {
		query = query.Where("name like ?", "%"+tagName+"%")
	}

	// 获取tag数据
	err1 := query.Offset(offset).Limit(pageSize).Find(&tags).Error

	// 获取tag总数
	err2 := query.Model(&Tag{}).Count(&count).Error
	if err1 != nil || err2 != nil {
		return []Tag{}, 0
	}
	return
}

func AddTag(tagName string, tagStatus int) bool {
	global.DB.Create(&Tag{
		TagName:   tagName,
		TagStatus: tagStatus,
	})
	return true
}

func UpdateTag(tagId int, tag map[string]interface{}) {
	global.DB.Model(&Tag{}).Where("id = ?", tagId).Updates(tag)
}

func Get(tagId int) bool {
	err := global.DB.First(&Tag{}, tagId).Error
	if err != nil {
		return false
	}

	return true
}

func DeleteTag(tagId int) bool {
	err := global.DB.Where("id = ?", tagId).Delete(&Tag{}).Error
	if err != nil {
		return false
	}
	return true
}
