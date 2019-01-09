package handler

import (
	"config/model"
	"config/server_model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

func AddUserHandler(c *gin.Context)  {
	var req server_model.AddUpdateUser
	if err := c.BindJSON(&req); err != nil {
		log.Error("AddUserHandler params parse err", "err", err.Error())
		c.JSON(400, &gin.H{
			"errmsg":"t001",
			"code":400,
		})
		return
	}
	if err := req.CheckParams(); err != nil {
		c.JSON(400, &gin.H{
			"errmsg": err.Error(),
			"code":400,
		})
		return
	}
	user := &model.UserInfo{}
	if err := user.CreateUser(&req); err != nil {
		c.JSON(400, &gin.H{
			"errmsg":err.Error(),
			"code":400,
		})
		return
	}
	c.JSON(200, &gin.H{
		"info":"create success",
		"code":200,
	})
}

func Logger() gin.HandlerFunc {
	logClient := log.New()
	var logPath = "/var/log/katy"
	//if util.PathExists(logPath) {
	//	os.MkdirAll(logPath, os.ModePerm)
	//}
	fileName := path.Join(logPath, "gin-api.log")
	//禁止logrus的输出
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err!= nil{
		fmt.Println("err", err)
	}
	logClient.Out = src
	logClient.SetLevel(log.DebugLevel)
	//apiLogPath := "gin-api.log"
	logWriter, err := rotatelogs.New(
		fileName+".%Y-%m-%d-%H-%M.log",
		rotatelogs.WithLinkName(fileName), // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour), // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	writeMap := lfshook.WriterMap{
		log.InfoLevel:  logWriter,
		log.FatalLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &log.JSONFormatter{})
	logClient.AddHook(lfHook)


	return func (c *gin.Context) {
		// 开始时间
		start := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		end := time.Now()
		//执行时间
		latency := end.Sub(start)

		path := c.Request.URL.Path

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		logClient.Infof("| %3d | %13v | %15s | %s  %s |",
			statusCode,
			latency,
			clientIP,
			method, path,
		)
	}
}