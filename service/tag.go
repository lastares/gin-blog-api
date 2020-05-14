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


// 标签更新
func (t *tagService) Update(tagId, tagStatus int, tagName string) int {
	// 密等
	getError := repository.Tag.Get(tagId)
	if getError != nil {
		return response.TAG_NOT_FOUND
	}

	// 更新
	tag := &models.Tag{
		Id: tagId,
		TagName: tagName,
		TagStatus: tagStatus,
	}
	err := repository.Tag.Update(tag)
	if err != nil {
		return response.TAG_UPDATE_FAILED
	}
	return response.Ok;
}

// 标签删除
func (t *tagService) Delete(tagId int) int {
	// 密等
	getError := repository.Tag.Get(tagId)
	if getError != nil {
		return response.TAG_NOT_FOUND
	}

	// 删除
	err := repository.Tag.Delete(tagId)
	if err != nil {
		return response.TAG_DELETE_FAILED
	}
	return response.Ok;
}

// 标签列表
func (t *tagService) TagList(page, pageSize int, tagName string) []models.Tag {
	offset := (page - 1) * pageSize
	tags := repository.Tag.GetTags(offset, pageSize, tagName)
	return tags
}
