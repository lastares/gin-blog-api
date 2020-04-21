module go-gin-blog-api

go 1.14

require (
	github.com/gin-gonic/gin v1.6.2
	github.com/go-ini/ini v1.55.0
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/jinzhu/gorm v1.9.12
	github.com/unknwon/com v1.0.1
	gopkg.in/ini.v1 v1.55.0 // indirect
)

replace (
	./conf => /usr/local/study/go/go-gin-blog-api/conf
	./middleware => /usr/local/study/go/go-gin-blog-api/middleware
	./models => /usr/local/study/go/go-gin-blog-api/models
	./pkg/setting => /usr/local/study/go/go-gin-blog-api/pkg/setting
	./routers => /usr/local/study/go/go-gin-blog-api/routers
)
