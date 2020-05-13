package service

import (
	"go-gin-blog-api/models"
	"go-gin-blog-api/repository"
	"go-gin-blog-api/response"
	"net/http"
)

var Tag = newTagService()

func newTagService() *tagService {
	return &tagService{}
}

type tagService struct {
}

func (t *tagService) Create(tagName string, tagStatus int) int {
	tag := &models.Tag{
		TagName: tagName,
		TagStatus: tagStatus,
	}
	err := repository.Tag.Create(tag)
	if err != nil {
		return response.TAG_CREATED_FAILED
	}
	return http.StatusOK;
}
