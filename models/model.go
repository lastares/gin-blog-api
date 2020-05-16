package models

import (
	"go-gin-blog-api/util"
)

// 文章
type Article struct {
	Id            int           `json:"id" gorm:"column:id;primary_key"`
	Title         string        `json:"title" gorm:"column:title"`
	Content       string        `json:"content" gorm:"column:content"`
	CurrentStatus int           `json:"currentStatus" gorm:"column:current_status"`
	Tags          []Tag  `json:"tags" gorm:"many2many:article_tag"`
	Comments      []Comment     `json:"comments" gorm:"foreignkey:ArticleId;association_foreignkey:ArticleId"`
	CreatedAt     util.JSONTime `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt     util.JSONTime `json:"updatedAt" gorm:"column:updated_at"`
}

type ArticleTag struct {
	Id        int           `json:"id" gorm:"column:id;primary_key"`
	ArticleId int           `json:"articleId" gorm:"column:article_id;index"`
	TagId     int           `json:"tagId" gorm:"column:tag_id;index"`
	CreatedAt util.JSONTime `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt util.JSONTime `json:"updatedAt" gorm:"column:updated_at"`
}

// Tag example
// 标签
type Tag struct {
	Id        int           `json:"id" gorm:"column:id;primary_key"`
	TagName   string        `json:"tagName" gorm:"column:tag_name"`
	TagStatus int           `json:"tagStatus" gorm:"column:tag_status"`
	Articles  []Article  `json:"articles" gorm:"many2many:article_tag;"`
	CreatedAt util.JSONTime `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt util.JSONTime `json:"updatedAt" gorm:"column:updated_at"`
}

// 文章评论
type Comment struct {
	Id             int           `json:"id" gorm:"column:id;primary_key"`
	CommentContent string        `json:"commentContent" gorm:"column:comment_content"`
	ArticleId      int           `json:"articleId" gorm:"column:article_id"`
	UserId         int           `json:"userId" gorm:"column:user_id"`
	CreatedAt      util.JSONTime `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt      util.JSONTime `json:"updatedAt" gorm:"column:updated_at"`
}
