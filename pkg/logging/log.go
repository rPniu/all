package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

var (
	F                  *os.File
	DefaultPrefix      = ""
	DefaultCallerDepth = 2
	logger             *log.Logger
	logPrefix          = ""
	levelFlags         = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func init() {
	filePath := getLogFileFullPath()
	F = openLogFile(filePath)
	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}
func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v...)
}
func Debugf(format string, a ...any) {
	Fatal(fmt.Sprintf(format, a...))
}
func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v...)
}
func Infof(format string, a ...any) {
	Fatal(fmt.Sprintf(format, a...))
}
func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v...)
}
func Warnf(format string, a ...any) {
	Fatal(fmt.Sprintf(format, a...))
}
func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v...)
}
func Errorf(format string, a ...any) {
	Fatal(fmt.Sprintf(format, a...))
}
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v...)
}
func Fatalf(format string, a ...any) {
	Fatal(fmt.Sprintf(format, a...))
}
func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}
	logger.SetPrefix(logPrefix)
}
