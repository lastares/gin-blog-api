package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go-gin-blog-api/api/v1"
	"go-gin-blog-api/middleware"
	"gopkg.in/gomail.v2"
	"log"
	"net/http"
	"strconv"
)

func InitRouter(engine *gin.Engine) *gin.Engine {
	engine.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, world",
		})
	})

	engine.POST("/send/mail", func(c *gin.Context) {
		err :=SendMail([]string{"songyaofeng@aliyun.com", "1615730914@qq.com"}, "go语法发送邮件测试", "go语法发送邮件测试, 这是一侧测试邮件，请勿回复")

		if err != nil {
			log.Println(err)
			fmt.Println("send fail")
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "发送成功",
		})
	})

	//engine.GET("/auth", v1.GetAuth)
	apiV1 := engine.Group("/api/v1")
	apiV1.Use(middleware.Logger())
	{
		// 标签
		apiV1.POST("/tag/list", v1.GetTags)
		apiV1.POST("/tag/create", v1.TagCreate)
		apiV1.POST("/tag/update", v1.TagUpdate)
		apiV1.POST("/tag/delete", v1.TagDelete)

		// 文章
		apiV1.POST("/article/create", v1.ArticleCreate)
		apiV1.POST("/article/update", v1.ArticleUpdate)
		apiV1.POST("/article/list", v1.ArticleList)
		apiV1.POST("/article/delete", v1.ArticleDelete)

	}
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return engine
}

func SendMail(mailTo []string, subject string, body string) error {
	//定义邮箱服务器连接信息，如果是网易邮箱 pass填密码，qq邮箱填授权码

	//mailConn := map[string]string{
	//  "user": "xxx@163.com",
	//  "pass": "your password",
	//  "host": "smtp.163.com",
	//  "port": "465",
	//}

	mailConn := map[string]string{
		"user": "862761213@qq.com",
		"pass": "zqtigngoxngmbehi",
		"host": "smtp.qq.com",
		"port": "465",
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()

	m.SetHeader("From",  m.FormatAddress(mailConn["user"], "go语言测试发送邮件")) //这种方式可以添加别名，即“XX官方”
	m.SetHeader("To", mailTo...)    //发送给多个用户
	m.SetHeader("Subject", subject) //设置邮件主题
	m.SetBody("text/html", body)    //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	return err

}
