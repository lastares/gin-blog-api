package models

import "time"

type Model struct {
	Id        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
