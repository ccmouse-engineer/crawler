package helper

import "testing"

func TestLoggerFile(t *testing.T) {
	LoggerFile("test", "error", "filename: %s, filesize: %d", "content.txt", 512)
	LoggerFile("test", "info", "filename: %s, filesize: %d", "content.txt", 512)
}

func TestLoggerConsole(t *testing.T) {
	LoggerConsole("warn", "filename: %s, filesize: %d", "content.txt", 512)
	LoggerConsole("info", "filename: %s, filesize: %d", "content.txt", 512)
	LoggerConsole("fatal", "filename: %s, filesize: %d", "content.txt", 512)
}
