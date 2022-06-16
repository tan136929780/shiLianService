package logging

import (
	"fmt"
)

var (
	LogSavePath         = "./log/"
	CommonLogSaveName   = "common.log"
	AccessLogSaveName   = "access.log"
	SlowLogSaveName     = "slow.log"
	MetaDataLogSaveName = "metaData.log"
	UploadLogSaveName   = "upload.log"
	DownloadLogSaveName = "download.log"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

func getLogFileFullPath() string {
	return LogSavePath + CommonLogSaveName
}

func getAccessLogFileFullPath() string {
	return LogSavePath + AccessLogSaveName
}

func getSlowLogFileFullPath() string {
	return LogSavePath + SlowLogSaveName
}

func getMetaDataLogFileFullPath() string {
	return LogSavePath + MetaDataLogSaveName
}

func getUploadLogFileFullPath() string {
	return LogSavePath + UploadLogSaveName
}

func getDownloadLogFileFullPath() string {
	return LogSavePath + DownloadLogSaveName
}
