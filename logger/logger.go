package logger

import (
	"config/util"
	"fmt"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

func init() {
	// 每天的日志拆分
	ConfigLocalFilesystemLogger("/var/log/katy", "service-log", time.Hour*24*365, time.Hour * 24)
	//InitLogger("/var/log/katy", "service-log")
}

// config logrus log to local filesystem, with file rotation
func ConfigLocalFilesystemLogger(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) {
	if util.PathExists(logPath) {
		os.MkdirAll(logPath, os.ModePerm)
	}
	baseLogPath := path.Join(logPath, logFileName)
	fileName := baseLogPath + "-" + time.Now().Format("2006-01-02") + ".txt"
	_, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println("文件打开出错", err.Error())
		return
	}
	writer, err := rotatelogs.New(
		fileName,
		//rotatelogs.WithLinkName(baseLogPaht), // 生成软链，指向最新日志文件
		//
		//rotatelogs.WithMaxAge(maxAge),        // 文件最大保存时间
		//rotatelogs.WithRotationCount(365),  // 最多存365个文件

		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	if err != nil {
		log.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer, // 为不同级别设置不同的输出目的
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, &log.JSONFormatter{})
	log.AddHook(lfHook)
}

func InitLogger(logPath, logPrefix string) () {
	if util.PathExists(logPath) {
		os.MkdirAll(logPath, os.ModePerm)
	}
	baseLogPath := path.Join(logPath, logPrefix)
	fileName := baseLogPath + "-" + time.Now().Format("2006-01-02") + ".txt"
	logClient := log.New()
	//禁止logrus的输出
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println("err", err)
	}
	logClient.Out = src
	logClient.SetLevel(log.DebugLevel)
	//apiLogPath := "service.log"
	logWriter, err := rotatelogs.New(
		//baseLogPath+".%Y-%m-%d-%H-%M.log",
		fileName,
		rotatelogs.WithLinkName(baseLogPath),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	writeMap := lfshook.WriterMap{
		log.DebugLevel: logWriter, // 为不同级别设置不同的输出目的
		log.InfoLevel:  logWriter,
		log.WarnLevel:  logWriter,
		log.ErrorLevel: logWriter,
		log.FatalLevel: logWriter,
		log.PanicLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &log.JSONFormatter{})
	logClient.AddHook(lfHook)
}
