package repository

import (
	"go-gin-blog-api/global"
	"go-gin-blog-api/models"
)

var Tag = newTagRepository()

func newTagRepository() *tagRepository {
	return &tagRepository{}
}

type tagRepository struct {
}

func (t *tagRepository) Create(tag *models.Tag) error {
	return global.DB.Create(&tag).Error
}
