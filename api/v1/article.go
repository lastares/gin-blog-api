package v1

import (
	"github.com/gin-gonic/gin"
	"go-gin-blog-api/form_request"
	"go-gin-blog-api/global"
	"go-gin-blog-api/models/form_validate"
	"go-gin-blog-api/response"
	"go-gin-blog-api/service"
)

// 添加文章
func ArticleCreate(c *gin.Context) {
	var article form_validate.ArticleCreateForm
	c.BindJSON(&article)
	err := global.Validate.Struct(article)
	if err != nil {
		form_request.ValidFailed(c, err)
		return
	}

	errorCode := service.Article.Create(article)
	if errorCode != response.Ok {
		response.InvalidOperation(c, errorCode)
		return
	}
	response.Success(c)
}

// 更新文章
func ArticleUpdate(c *gin.Context) {
	var article form_validate.ArticleUpdateForm
	c.BindJSON(&article)
	err := global.Validate.Struct(article)
	if err != nil {
		form_request.ValidFailed(c, err)
		return
	}
	errorCode := service.Article.Update(article)
	if errorCode != response.Ok {
		response.InvalidOperation(c, errorCode)
		return
	}
	response.Success(c)
}

// 文章列表
func ArticleList(c *gin.Context) {
	var articleList form_validate.ArticleListForm
	c.ShouldBindJSON(&articleList)
	err := global.Validate.Struct(articleList)
	if err != nil {
		form_request.ValidFailed(c, err)
		return
	}
	articles, total := service.Article.ArticleList(
		articleList.Page,
		articleList.PageSize,
		articleList.Content,
	)
	response.ResponseSuccessJson(c, articles, total)
}

// 文章删除
func ArticleDelete(c *gin.Context) {
	// 字段校验
	var article form_validate.ArticleDeleteForm
	c.BindJSON(&article)
	err := global.Validate.Struct(article)
	if err != nil {
		form_request.ValidFailed(c, err)
		return
	}

	// 删除
	errorCode := service.Article.Delete(article.Id)
	if errorCode != response.Ok {
		response.InvalidOperation(c, errorCode)
		return
	}
	response.Success(c)
}

