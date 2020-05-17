package models

import "time"

type ArticleTag struct {
	Id        int `gorm:"column:id;primary_key"`
	ArticleId int `gorm:"column:article_id;index"`
	TagId     int `gorm:"column:tag_id;index"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
