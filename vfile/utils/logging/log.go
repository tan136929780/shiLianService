package logging

import (
	"bufio"
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"time"
)

var Logger = log.New()
var SlowLogger = log.New()
var AccessLogger = log.New()
var MetaDataLogger = log.New()
var UploadLogger = log.New()
var DownloadLogger = log.New()

var devNull *os.File

func InitAllLogger() {
	//正式环境关闭标准输出
	if viper.GetString("run_mode") == "prod" {
		var err error
		devNull, err = os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			fmt.Println("Open Src File err", err)
		} else {
			Logger.SetOutput(bufio.NewWriter(devNull))
		}
	}
	initLogger()
	initAccessLogger()
	initSlowLogger()
	initMetaDataLogger()
	initUploadLogger()
	initDownloadLogger()
}

func initLogger() {
	//正式环境关闭标准输出
	if devNull != nil {
		Logger.SetOutput(bufio.NewWriter(devNull))
	}
	//设置日志级别
	Logger.SetLevel(log.DebugLevel)

	//设置日志格式
	Logger.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 设置 rotatelogs
	logWriter, _ := rotatelogs.New(
		getLogFileFullPath()+".%Y%m%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(getLogFileFullPath()),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		log.InfoLevel:  logWriter,
		log.FatalLevel: logWriter,
		log.DebugLevel: logWriter,
		log.WarnLevel:  logWriter,
		log.ErrorLevel: logWriter,
		log.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 新增 Hook
	Logger.AddHook(lfHook)
}

func initSlowLogger() {
	if devNull != nil {
		SlowLogger.SetOutput(bufio.NewWriter(devNull))
	}
	//设置日志级别
	SlowLogger.SetLevel(log.DebugLevel)

	//设置日志格式
	SlowLogger.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 设置 rotatelogs
	logWriter, _ := rotatelogs.New(
		getSlowLogFileFullPath()+".%Y%m%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(getSlowLogFileFullPath()),

		// 设置最大保存时间
		rotatelogs.WithMaxAge(24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		log.InfoLevel:  logWriter,
		log.FatalLevel: logWriter,
		log.DebugLevel: logWriter,
		log.WarnLevel:  logWriter,
		log.ErrorLevel: logWriter,
		log.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 新增 Hook
	SlowLogger.AddHook(lfHook)
}

func initAccessLogger() {
	if devNull != nil {
		AccessLogger.SetOutput(bufio.NewWriter(devNull))
	}
	//设置日志级别
	AccessLogger.SetLevel(log.DebugLevel)

	//设置日志格式
	AccessLogger.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 设置 rotatelogs
	logWriter, _ := rotatelogs.New(
		getAccessLogFileFullPath()+".%Y%m%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(getAccessLogFileFullPath()),

		// 设置最大保存时间
		rotatelogs.WithMaxAge(3*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		log.InfoLevel:  logWriter,
		log.FatalLevel: logWriter,
		log.DebugLevel: logWriter,
		log.WarnLevel:  logWriter,
		log.ErrorLevel: logWriter,
		log.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 新增 Hook
	AccessLogger.AddHook(lfHook)
}

func initMetaDataLogger() {
	if devNull != nil {
		MetaDataLogger.SetOutput(bufio.NewWriter(devNull))
	}
	//设置日志级别
	MetaDataLogger.SetLevel(log.DebugLevel)

	//设置日志格式
	MetaDataLogger.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 设置 rotatelogs
	logWriter, _ := rotatelogs.New(
		getMetaDataLogFileFullPath()+".%Y%m%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(getMetaDataLogFileFullPath()),

		// 设置最大保存时间
		rotatelogs.WithMaxAge(3*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		log.InfoLevel:  logWriter,
		log.FatalLevel: logWriter,
		log.DebugLevel: logWriter,
		log.WarnLevel:  logWriter,
		log.ErrorLevel: logWriter,
		log.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 新增 Hook
	MetaDataLogger.AddHook(lfHook)
}

func initUploadLogger() {
	if devNull != nil {
		UploadLogger.SetOutput(bufio.NewWriter(devNull))
	}
	//设置日志级别
	UploadLogger.SetLevel(log.DebugLevel)

	//设置日志格式
	UploadLogger.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 设置 rotatelogs
	logWriter, _ := rotatelogs.New(
		getUploadLogFileFullPath()+".%Y%m%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(getUploadLogFileFullPath()),

		// 设置最大保存时间
		rotatelogs.WithMaxAge(3*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		log.InfoLevel:  logWriter,
		log.FatalLevel: logWriter,
		log.DebugLevel: logWriter,
		log.WarnLevel:  logWriter,
		log.ErrorLevel: logWriter,
		log.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 新增 Hook
	UploadLogger.AddHook(lfHook)
}

func initDownloadLogger() {
	if devNull != nil {
		DownloadLogger.SetOutput(bufio.NewWriter(devNull))
	}
	//设置日志级别
	DownloadLogger.SetLevel(log.DebugLevel)

	//设置日志格式
	DownloadLogger.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 设置 rotatelogs
	logWriter, _ := rotatelogs.New(
		getDownloadLogFileFullPath()+".%Y%m%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(getDownloadLogFileFullPath()),

		// 设置最大保存时间
		rotatelogs.WithMaxAge(3*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		log.InfoLevel:  logWriter,
		log.FatalLevel: logWriter,
		log.DebugLevel: logWriter,
		log.WarnLevel:  logWriter,
		log.ErrorLevel: logWriter,
		log.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 新增 Hook
	DownloadLogger.AddHook(lfHook)
}
