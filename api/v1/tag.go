package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	http2 "go-gin-blog-api/http"
	"go-gin-blog-api/models"
	"net/http"
	"strings"
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
func GetTags(c *gin.Context) {
	// 获取标签列表query参数
	tagName := strings.ToLower(c.DefaultQuery("name", ""))
	page := com.StrTo(c.DefaultQuery("page", "1")).MustInt()
	tagStatus := com.StrTo(c.DefaultQuery("tagStatus", "1")).MustInt()
	data := make(map[string]interface{})

	var count int
	data["list"], count = models.GetTags(page, tagName, tagStatus)
	data["meta"] = Meta{count}

	// 返回标签列表json数据
	c.JSON(http.StatusOK, http2.Success{http.StatusOK, data})
}

// 添加文章标签
func AddTag(c *gin.Context) {
	tagName := c.PostForm("tagName")
	tagStatus := com.StrTo(c.PostForm("tagStatus")).MustInt()
	models.AddTag(tagName, tagStatus)
	c.JSON(http.StatusOK, http2.Ok{http.StatusOK})
}

// 编辑文章标签
func UpdateTag(c *gin.Context) {
	// 接受参数
	tagId := com.StrTo(c.PostForm("tagId")).MustInt()
	tagName := c.PostForm("tagName")
	tagStatus := com.StrTo(c.PostForm("tagStatus")).MustInt()

	// 更新数据map
	data := make(map[string]interface{})
	data["tagName"] = tagName
	data["tagStatus"] = tagStatus
	// 更新
	models.UpdateTag(tagId, data)
	c.JSON(http.StatusOK, http2.Ok{http.StatusOK})
}

// 删除标签
func DeleteTag(c *gin.Context) {
	tagId := com.StrTo(c.Param("id")).MustInt()
	if !models.Get(tagId) {
		c.JSON(http.StatusOK, http2.Failed{http2.ErrorCode, http2.Translate("Tag does not exist.")})
		return
	}

	if !models.DeleteTag(tagId) {
		c.JSON(http.StatusOK, http2.Failed{http2.ErrorCode, http2.Translate("Deleted failed.")})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
	})
}
