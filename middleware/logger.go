package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {
	// 写入文件
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	// 实例化
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 设置输出
	logger.Out = src

	// 设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	// 设置 rotatelogs
	currentPath, _ := os.Getwd()
	logFileName := currentPath + "/log/api.log"
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		 logFileName + ".%Y%m%d.log",

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7 * 24 * time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 新增 Hook
	logger.AddHook(lfHook)

	return func(c *gin.Context) {
		//body, _ := ioutil.ReadAll(c.Request.Body)
		//params := string(body)

		// 开始时间
		startTime := time.Now()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 请求参数
		//params := c.GetPostFormArray()


		// 日志格式
		logger.WithFields(logrus.Fields{
			"code":         statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"method":       reqMethod,
			"uri":          reqUri,
			//"json_params": params,
		}).Info()

		// 处理请求
		c.Next()
	}
}
