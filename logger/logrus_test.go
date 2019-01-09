package logger

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestLogrus(t *testing.T)  {
	//log.WithFields(log.Fields{
	//	"animal": "walrus",
	//	"size":   10,
	//}).Info("A group of walrus emerges from the ocean")
	//
	//log.WithFields(log.Fields{
	//	"omg":    true,
	//	"number": 122,
	//}).Warn("The group's number increased tremendously!")
	//
	//log.WithFields(log.Fields{
	//	"omg":    true,
	//	"number": 100,
	//}).Fatal("The ice breaks!")
	//ConfigLocalFilesystemLogger("/var/log/katy", "lg", time.Hour*24*365, time.Minute)

	//for {
	//	time.Sleep(time.Second*3)
		log.Info("hello ")
	//}
}
//
//// config logrus log to local filesystem, with file rotation
//func ConfigLocalFilesystemLogger(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) {
//	if util.PathExists(logPath) {
//		os.MkdirAll(logPath, os.ModePerm)
//	}
//	baseLogPaht := path.Join(logPath, logFileName)
//	writer, err := rotatelogs.New(
//		baseLogPaht+"-%Y%m%d%H%M",
//		//rotatelogs.WithLinkName(baseLogPaht), // 生成软链，指向最新日志文件
//		//
//		//rotatelogs.WithMaxAge(maxAge),        // 文件最大保存时间
//		// rotatelogs.WithRotationCount(365),  // 最多存365个文件
//
//		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
//	)
//	if err != nil {
//		log.Errorf("config local file system logger error. %+v", errors.WithStack(err))
//	}
//	lfHook := lfshook.NewHook(lfshook.WriterMap{
//		log.DebugLevel: writer, // 为不同级别设置不同的输出目的
//		log.InfoLevel:  writer,
//		log.WarnLevel:  writer,
//		log.ErrorLevel: writer,
//		log.FatalLevel: writer,
//		log.PanicLevel: writer,
//	}, &log.JSONFormatter{})
//	log.AddHook(lfHook)
//}