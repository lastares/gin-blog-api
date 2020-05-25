package docs

import "go-gin-blog-api/util"

// Ta example
type TagListDTO struct {
	Code int `json:"code"`
	Data []Tag `json:"data" example:"Tag"`
}

type Tag struct {
	Id        int           `json:"id" gorm:"column:id;primary_key"`
	TagName   string        `json:"tagName" gorm:"column:tag_name"`
	TagStatus int           `json:"tagStatus" gorm:"column:tag_status"`
	CreatedAt util.JSONTime `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt util.JSONTime `json:"updatedAt" gorm:"column:updated_at"`
}