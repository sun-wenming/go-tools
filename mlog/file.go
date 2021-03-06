package mlog

import (
	"fmt"
	"time"
)

var (
	LogSavePath = "runtime/logs/"
	LogSaveName = "log"
	LogFileExt  = "log"
	TimeFormat  = "20060102"
)

// 获取日志的存放目录
func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

// 默认为 log20060102.log
func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		LogSaveName,
		time.Now().Format(TimeFormat),
		LogFileExt,
	)
}

// func openLogFile(filePath string) *os.File {
// 	_, err := os.Stat(filePath)
// 	switch {
// 	case os.IsNotExist(err):
// 		mkDir()
// 	case os.IsPermission(err):
// 		log.Fatalf("Permission :%v", err)
// 	}

// 	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		log.Fatalf("Fail to OpenFile :%v", err)
// 	}

// 	return handle
// }

// func mkDir() {
// 	// 返回与当前目录对应的根路径名
// 	dir, _ := os.Getwd()
// 	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
// 	if err != nil {
// 		panic(err)
// 	}
// }
