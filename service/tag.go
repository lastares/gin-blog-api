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
	var tag models.Tag
	tag.SetTagName(tagName)
	tag.SetTagStatus(tagStatus)

	err := repository.Tag.Create(&tag)
	if err != nil {
		return response.TAG_CREATED_FAILED
	}
	return http.StatusOK;
}


// 标签更新
func (t *tagService) Update(tagId, tagStatus int, tagName string) int {
	// 密等
	tag := repository.Tag.Get(tagId)
	if tag.Id == 0 {
		return response.TAG_NOT_FOUND
	}

	// 更新
	tag.SetTagName(tagName)
	tag.SetTagStatus(tagStatus)

	err := repository.Tag.Update(&tag)
	if err != nil {
		return response.TAG_UPDATE_FAILED
	}
	return response.Ok;
}

// 标签删除
func (t *tagService) Delete(tagId int) int {
	// 密等
	tag := repository.Tag.Get(tagId)
	if tag.Id == 0 {
		return response.TAG_NOT_FOUND
	}

	// 删除
	err := repository.Tag.Delete(tag)
	if err != nil {
		return response.TAG_DELETE_FAILED
	}
	return response.Ok;
}

// 标签列表
func (t *tagService) TagList(page, pageSize int, tagName string) ([]models.Tag, int) {
	offset := (page - 1) * pageSize
	tags, total := repository.Tag.GetTags(offset, pageSize, tagName)
	return tags, total
}
