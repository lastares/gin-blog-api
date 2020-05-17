package models

import (
	"go-gin-blog-api/util"
	"time"
)
type ArticleTag struct {
	Id        int `gorm:"column:id;primary_key"`
	ArticleId int `gorm:"column:article_id;index"`
	TagId     int `gorm:"column:tag_id;index"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

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