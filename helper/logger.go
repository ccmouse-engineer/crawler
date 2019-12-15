package helper

import (
	"os"

	"github.com/sirupsen/logrus"
)

const LogFileExt = ".log"
const logFilePerm = 0666

var logger = logrus.New()

// LoggerFile写入日志到文件
// Param: level(e.g: panic, fatal, error, warn, warning, info, debug, trace)
// Param: filename(e.g: logrus)
// Param: format(e.g: "filename: %s, filesize: %d")
// Param: args(e.g: "test.txt", 512)
func LoggerFile(filename string, level string, format string, args ...interface{}) {
	file, err := OpenFile(filename, LogFileExt, logFilePerm)
	if err == nil {
		logger.Out = file
	} else {
		logger.Info("Failed to log to file, using default stderr")
	}
	defer file.Close()
	parseLevel, err := logrus.ParseLevel(level)
	if err != nil {
		panic(err)
	}
	logger.Logf(parseLevel, format, args...)
}

// LoggerFile输入到
// Param: level(e.g: panic, fatal, error, warn, warning, info, debug, trace)
// Param: format(e.g: "filename: %s, filesize: %d")
// Param: args(e.g: "test.txt", 512)
func LoggerConsole(level string, format string, args ...interface{}) {
	logger.Out = os.Stdout
	switch level {
	case "panic":
		logger.Panicf(format, args...)
	case "fatal":
		logger.Fatalf(format, args...)
	case "error":
		logger.Errorf(format, args...)
	case "warn", "warning":
		logger.Warnf(format, args...)
	case "info":
		logger.Infof(format, args...)
	case "debug":
		logger.Debugf(format, args...)
	case "trace":
		logger.Tracef(format, args...)
	}
}

// 打开文件
func OpenFile(filename string, ext string, perm uint32) (*os.File, error) {
	return os.OpenFile(filename+ext, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.FileMode(perm))
}
