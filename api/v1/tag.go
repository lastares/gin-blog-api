package v1

import (
	"github.com/gin-gonic/gin"
	"go-gin-blog-api/form_request"
	"go-gin-blog-api/global"
	"go-gin-blog-api/models/form_validate"
	"go-gin-blog-api/response"
	"go-gin-blog-api/service"
	"net/http"
)

// @Summary 获取标签列表
// @Tags Tag
// @Produce  json
// @Router /tags [post]
// @Param page formData int false "当前页码"
// @Param pageSize formData int false "每页显示条数"
// @Success 200 {object} models.Tag
// @Failure 500 {object} http.Failed
func GetTags(c *gin.Context) {
	var tagList form_validate.TagListForm
	c.BindJSON(&tagList)
	err := global.Validate.Struct(tagList)
	if err != nil {
		form_request.ValidFailed(c, err)
		return
	}

	tags, total := service.Tag.TagList(
		tagList.Page,
		tagList.PageSize,
		tagList.TagName,
	)
	response.ResponseSuccessJson(c, tags, total)
}

// 添加文章标签
func TagCreate(c *gin.Context) {
	var tag form_validate.TagCreateModel
	c.BindJSON(&tag)
	err := global.Validate.Struct(tag)
	if err != nil {
		form_request.ValidFailed(c, err)
		return
	}

	if errorCode := service.Tag.Create(tag.TagName, tag.TagStatus); errorCode != http.StatusOK {
		response.InvalidOperation(c, errorCode)
		return
	}
	response.Success(c)
}

// 编辑文章标签
func TagUpdate(c *gin.Context) {
	// 字段校验
	var tag form_validate.TagUpdateModel
	c.BindJSON(&tag)
	err := global.Validate.Struct(tag)
	if err != nil {
		form_request.ValidFailed(c, err)
		return
	}

	// 更新
	errorCode := service.Tag.Update(
		tag.Id,
		tag.TagStatus,
		tag.TagName,
	)
	if errorCode != response.Ok {
		response.InvalidOperation(c, errorCode)
		return
	}
	response.Success(c)
}

// 删除标签
func TagDelete(c *gin.Context) {
	// 字段校验
	var tag form_validate.TagDeleteModel
	c.BindJSON(&tag)
	err := global.Validate.Struct(tag)
	if err != nil {
		form_request.ValidFailed(c, err)
		return
	}

	// 删除
	errorCode := service.Tag.Delete(tag.Id)
	if errorCode != response.Ok {
		response.InvalidOperation(c, errorCode)
		return
	}
	response.Success(c)
}
