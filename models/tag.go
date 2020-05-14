package models

import (
	_ "github.com/joho/godotenv/autoload"
	"time"
)

const (
	TAG_STATUS_NORMAL = 10
	TAG_STATUS_LOCK = 20
)

// Tag example
type Tag struct {
	Id        int `gorm:"primary_key"`
	TagName   string `json:"tagName" validate:"required" label:"标签名称"`
	TagStatus int    `json:"tagStatus" validate:"required,oneof=10 20" label:"标签状态"`
	CreatedAt time.Time
	UpdatedAt time.Time
}