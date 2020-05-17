package models

import "go-gin-blog-api/util"


const (
	TAG_STATUS_NORMAL = 10
	TAG_STATUS_LOCK = 20
)


// Tag example
// 标签
type Tag struct {
	Id        int           `json:"id" gorm:"column:id;primary_key"`
	TagName   string        `json:"tagName" gorm:"column:tag_name"`
	TagStatus int           `json:"tagStatus" gorm:"column:tag_status"`
	Articles  []Article     `json:"articles" gorm:"many2many:article_tag;"`
	CreatedAt util.JSONTime `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt util.JSONTime `json:"updatedAt" gorm:"column:updated_at"`
}
