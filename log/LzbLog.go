package log

import (
	"fmt"
	"log"
	"log-lzbagent/conf"
	"strings"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

//日志级别，error、warning、info
const (
	LevelError = iota
	LevelWarning
	LevelInfo
)

type Log struct {
	err  *log.Logger
	warn *log.Logger
	info *log.Logger

	lzbDepth int
	//日志级别
	level int
}
var lzbLog *Log
func createLogger() *Log {
	fileName := conf.GetConf().Local.FilePath
	maxsize := conf.GetConf().Local.LogMaxSize
	flag := log.LstdFlags

	jack := &lumberjack.Logger{
		Filename: fileName,
		MaxSize:  maxsize, // megabytes
	}

	logger := new(Log)

	logger.err = log.New(jack, "[E] ", flag)
	logger.warn = log.New(jack, "[W] ", flag)
	logger.info = log.New(jack, "[I] ", flag)
	logger.lzbDepth = 2
	//logger.SetLevel(LevelInformational)
	lzbLog = logger
	return lzbLog

}
func LogInit(){
	createLogger()
}
func (ll *Log) Error(format string, v ...interface{}) {
	ll.err.Output(ll.lzbDepth, fmt.Sprintf(format, v...))
}

func Error(f interface{}, v ...interface{}) {
	lzbLog.Error(formatLog(f, v...))
}
func (ll *Log) Warn(format string, v ...interface{}) {
	ll.warn.Output(ll.lzbDepth, fmt.Sprintf(format, v...))
}

func Warn(f interface{}, v ...interface{}) {
	lzbLog.Warn(formatLog(f, v...))
}
func (ll *Log) Info(format string, v ...interface{}) {
	ll.info.Output(ll.lzbDepth, fmt.Sprintf(format, v...))
}

func Info(f interface{}, v ...interface{}) {
	lzbLog.Info(formatLog(f, v...))
}

func formatLog(f interface{}, v ...interface{}) string {
	var msg string
	switch f.(type) {
	case string:
		msg = f.(string)
		if len(v) == 0 {
			return msg
		}
		if strings.Contains(msg, "%") && !strings.Contains(msg, "%%") {
			//format string
		} else {
			//do not contain format char
			msg += strings.Repeat(" %v", len(v))
		}
	default:
		msg = fmt.Sprint(f)
		if len(v) == 0 {
			return msg
		}
		msg += strings.Repeat(" %v", len(v))
	}
	return fmt.Sprintf(msg, v...)
}
