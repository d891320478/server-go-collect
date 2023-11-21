package baselog

import (
	"fmt"
	"os/exec"

	"github.com/jeanphorn/log4go"
)

func InitLog(appName string) {
	// 创建日志目录
	logPath := "/data/logs/" + appName
	createLogCmd := exec.Command("mkdir", "-p", logPath)
	createLogOut, err := createLogCmd.Output()
	fmt.Println(string(createLogOut))
	fmt.Println(err)
	// 输出到控制台,级别为CRITICAL
	log4go.AddFilter("stdout", log4go.CRITICAL, log4go.NewConsoleLogWriter())
	// info
	infoLogWriter := log4go.NewFileLogWriter(logPath+"/error.log", true, true)
	infoLogWriter.SetRotateSize(100 * 1024 * 1024)
	infoLogWriter.SetRotateMaxBackup(7)
	infoLogWriter.SetFormat("%D %T %L %S - %M")
	log4go.AddFilter("info", log4go.INFO, infoLogWriter)
	// error
	errorLogWriter := log4go.NewFileLogWriter(logPath+"/error.log", true, true)
	errorLogWriter.SetRotateSize(100 * 1024 * 1024)
	errorLogWriter.SetRotateMaxBackup(7)
	errorLogWriter.SetFormat("%D %T %L %S - %M")
	log4go.AddFilter("error", log4go.ERROR, errorLogWriter)
}

func InfoLog() *log4go.Filter {
	return log4go.LOGGER("file")
}

func ErrorLog() *log4go.Filter {
	return log4go.LOGGER("error")
}
