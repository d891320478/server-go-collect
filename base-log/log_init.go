package baselog

import (
	"os/exec"

	"github.com/jeanphorn/log4go"
)

var rotateSize = 100 * 1024 * 1024
var format = "%D %T %L %S - %M"
var maxBackup = 7
var infoLog = "info"
var errorLog = "error"

func InitLog(appName string) {
	// 创建日志目录
	logPath := "/data/logs/" + appName
	createLogCmd := exec.Command("mkdir", "-p", logPath)
	createLogCmd.Output()
	// 输出到控制台,级别为CRITICAL
	log4go.AddFilter("stdout", log4go.CRITICAL, log4go.NewConsoleLogWriter())
	// info
	infoLogWriter := log4go.NewFileLogWriter(logPath+"/info.log", true, true)
	infoLogWriter.SetRotateSize(rotateSize)
	infoLogWriter.SetRotateMaxBackup(maxBackup)
	infoLogWriter.SetFormat(format)
	log4go.AddFilter(infoLog, log4go.INFO, infoLogWriter)
	// error
	errorLogWriter := log4go.NewFileLogWriter(logPath+"/error.log", true, true)
	errorLogWriter.SetRotateSize(rotateSize)
	errorLogWriter.SetRotateMaxBackup(maxBackup)
	errorLogWriter.SetFormat(format)
	log4go.AddFilter(errorLog, log4go.ERROR, errorLogWriter)
}

func InfoLog() *log4go.Filter {
	return log4go.LOGGER(infoLog)
}

func ErrorLog() *log4go.Filter {
	return log4go.LOGGER(errorLog)
}
