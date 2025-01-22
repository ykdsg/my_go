package main

import (
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

/*
 * 基础日志轮转配置
 */
func init() {
	path := "./log.log"

	/* 日志轮转相关函数
	`WithLinkName` 为最新的日志建立软连接
	`WithRotationTime` 设置日志分割的时间，隔多久分割一次
	WithMaxAge 和 WithRotationCount二者只能设置一个
	  `WithMaxAge` 设置文件清理前的最长保存时间
	  `WithRotationCount` 设置文件清理前最多保存的个数
	*/
	writer, _ := rotatelogs.New(
		path+".%Y%m%d",
		// rotatelogs.WithLinkName(path), // 生成软链，指向最新日志文件
		rotatelogs.WithRotationCount(10),
		rotatelogs.WithRotationTime(time.Hour*24), // 日志切割时间间隔
	)

	log.SetOutput(writer)
	log.SetFormatter(&logrus.JSONFormatter{})
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
