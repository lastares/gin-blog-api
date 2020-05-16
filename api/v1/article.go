package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-blog-api/form_request"
	"go-gin-blog-api/global"
	"go-gin-blog-api/models/form_validate"
	"go-gin-blog-api/response"
	"go-gin-blog-api/service"
	"net/http"
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

	fmt.Println(article)
	errorCode := service.Article.CreateArticle(
		article.Title,
		article.Content,
		article.CurrentStatus,
		article.TagIds,
	)
	if errorCode != http.StatusOK {
		response.InvalidOperation(c, errorCode)
		return
	}
	response.Success(c)
}

// 文章列表
func ArticleList(c *gin.Context) {
	var articleList form_validate.ArticleListForm
	c.BindJSON(&articleList)
	err := global.Validate.Struct(articleList)
	if err != nil {
		form_request.ValidFailed(c, err)
		return
	}
 fmt.Println(articleList)
	articles, total := service.Article.ArticleList(
		articleList.Page,
		articleList.PageSize,
		articleList.Content,
	)
	response.ResponseSuccessJson(c, articles, total)
}
