package v1

import (
	"github.com/gin-gonic/gin"
	"go-gin-blog-api/form_request"
	"go-gin-blog-api/global"
	"go-gin-blog-api/models/request"
	"go-gin-blog-api/response"
	"go-gin-blog-api/service"
	"net/http"
)

type Meta struct {
	Total int
}

// @Summary 获取标签列表
// @Tags Tag
// @Produce  json
// @Router /tags [post]
// @Param page formData int false "当前页码"
// @Param pageSize formData int false "每页显示条数"
// @Success 200 {object} models.Tag
// @Failure 500 {object} http.Failed
//func GetTags(c *gin.Context) {
//	// 获取标签列表query参数
//	tagName := strings.ToLower(c.DefaultQuery("name", ""))
//	page := com.StrTo(c.DefaultQuery("page", "1")).MustInt()
//	tagStatus := com.StrTo(c.DefaultQuery("tagStatus", "1")).MustInt()
//	data := make(map[string]interface{})
//
//	var count int
//	data["list"], count = models.GetTags(page, tagName, tagStatus)
//	data["meta"] = Meta{count}
//
//	// 返回标签列表json数据
//	c.JSON(http.StatusOK, http2.Success{http.StatusOK, data})
//}

// 添加文章标签
func CreateTag(c *gin.Context) {
	var tag request.TagCreateModel
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
//func Update(c *gin.Context) {
//	// 字段校验
//	var tag models.Tag
//	c.BindJSON(&tag)
//	err := global.Validate.Struct(tag)
//	if err != nil {
//		for _, err := range err.(validator.ValidationErrors) {
//			c.JSON(http.StatusOK, http2.Failed{http2.ErrorCode, err.Translate(global.Translator)})
//			return
//		}
//	}

	//tagService.(tag)
	// 判断标签是否存在
	//if models.GetTagById(tag) == false {
	//	c.JSON(http.StatusOK, http2.Failed{
	//		http2.ErrorCode,
	//		http2.Translate("The tag tag not found."),
	//	})
	//	return
	//}
	//
	//// 更新标签
	//if models.UpdateTag(tag) == false {
	//	c.JSON(http.StatusOK, http2.Failed{
	//		http2.ErrorCode,
	//		http2.Translate("Updated failed."),
	//	})
	//	return
	//}
	//c.JSON(http.StatusOK, http2.Ok{http.StatusOK})
//}

// 删除标签
//func DeleteTag(c *gin.Context) {
//	tagId := com.StrTo(c.Param("id")).MustInt()
//	if !models.Get(tagId) {
//		c.JSON(http.StatusOK, http2.Failed{http2.ErrorCode, http2.Translate("Tag does not exist.")})
//		return
//	}
//
//	if !models.DeleteTag(tagId) {
//		c.JSON(http.StatusOK, http2.Failed{http2.ErrorCode, http2.Translate("Deleted failed.")})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"code": http.StatusOK,
//	})
//}
