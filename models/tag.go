package models

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/unknwon/com"
	"go-gin-blog-api/global"
	"os"
)

// Tag example
type Tag struct {
	Model
	TagName   string `json:"tagName" validate:"required" label:"标签名称"`
	TagStatus int    `json:"tagStatus" validate:"required, oneof=10 20" label:"标签状态"`
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
