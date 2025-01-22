package main

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"time"
)

var log *logrus.Logger

func ConfigLog(name string) *rotatelogs.RotateLogs {
	fileName := fmt.Sprintf("./log/%s.log", name)
	writer, err := rotatelogs.New(
		fileName+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(fileName),           // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(time.Second*60*3),     // 文件最大保存时间
		rotatelogs.WithRotationTime(time.Second*60), // 日志切割时间间隔
	)
	if err != nil {
		log.Fatal("Init log failed, err:", err)
	}
	return writer
}

func init() {
	log = logrus.New()
	infoLog := ConfigLog("info")
	errorLog := ConfigLog("error")
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: ConfigLog("debug"), // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  infoLog,
		logrus.WarnLevel:  infoLog,
		logrus.ErrorLevel: errorLog,
		logrus.FatalLevel: errorLog,
		logrus.PanicLevel: errorLog,
	}, &logrus.TextFormatter{DisableColors: true})
	log.AddHook(lfHook)
}

func main() {
	for i := 0; i < 10; i++ {
		log.Infof("test")
		// logrus不推荐使用冗长的消息来记录运行信息,它推荐使用Fields来进行精细化的、结构化的信息记录
		log.WithFields(logrus.Fields{
			"omg":    true,
			"number": 22}).Warn("The number is big")

		entry := log.WithFields(logrus.Fields{
			"omg":    true,
			"number": 100,
		})
		entry.Error("The number is so big")

		// log.Fatal() 输出日志后进程退出
		// log.Panic() 输出日志进程退出，并抛出异常
	}
}
