package logging

import (
	"fmt"
)

var (
	GraylogSavePath   = "./log/"
	LogSavePath       = "./log/"
	LogSaveName       = "go.log"
	AccessLogSaveName = "go.access"
	SlowLogSaveName   = "go.slow"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

func getLogFileFullPath() string {
	return GraylogSavePath + LogSaveName
}

func getAccessLogFileFullPath() string {
	return LogSavePath + AccessLogSaveName
}

func getSlowLogFileFullPath() string {
	return LogSavePath + SlowLogSaveName
}
