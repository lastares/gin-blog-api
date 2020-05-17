package models

import "go-gin-blog-api/util"

type ArticleAttachment struct {
	Id             int           `json:"id" gorm:"column:id;primary_key"`
	AttachmentPath string        `json:"attachment_path" gorm:"column:attachment_path"`
	ArticleId      int           `json:"articleId" gorm:"index"`
	CreatedAt      util.JSONTime `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt      util.JSONTime `json:"updatedAt" gorm:"column:updated_at"`
}
