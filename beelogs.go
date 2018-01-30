package beelogger

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"os"
	"time"
)

// consoleLogs开发模式下日志
var consoleLogs *logs.BeeLogger

// FileLogs 生产环境下日志
var FileLogs *logs.BeeLogger

//运行方式
var runmode string

const (
	Emergency 
	Alert
	Critical
	Error
	Warning
	Notice
	Info
	Debug
	Trace
	
)
type Beelog struct {
	consoleLogs *logs.BeeLogger
	FileLogs  *logs.BeeLogger
	level string
	runmode string

}

func InitLogs(devmode string,logpath string,beelog Beelog)  *Beelog{
	/*命令行日志*/
	beelog.consoleLogs = logs.NewLogger(1)
	beelog.consoleLogs.SetLogger(logs.AdapterConsole)
	beelog.consoleLogs.EnableFuncCallDepth(true) //输出行号
	//beelog.consoleLogs.Async()                   //异步

	/* 文件日志*/
	beelog.FileLogs = logs.NewLogger(10000)
	beelog.level = "7"
	beelog.runmode = devmode
	if beelog.runmode == "" {
		beelog.runmode = "DEBUG"
	}
	timeNow := time.Now().Format("20060102")
	logPath := logpath+"/" + timeNow
	file, _ := PathExists(logPath)

	if file == false  {
		err := os.MkdirAll(logPath, os.ModePerm)
		if err != nil {
			fmt.Println("日志文件目录无权限创建", err)
			panic("err")
		}
	}
	
	FileLogsConfig := `{
		"filename":"`+ logpath +`/` + timeNow + `/beelog.log",
		"level":` + beelog.level + `,
		"daily":true,
		"maxdays":10
	}`
	beelog.FileLogs.SetLogger(logs.AdapterMultiFile, FileLogsConfig)
	beelog.FileLogs.EnableFuncCallDepth(true) //输出行号
	//beelog.FileLogs.Async()                   //异步



	return &beelog
}
func (this *Beelog) Fatal(v interface{}, r ...interface{}) {
	this.toLog("emergency", v)
}
func (this *Beelog) Emergency(v interface{}, r ...interface{}) {
	this.toLog("emergency", v)
}
func (this *Beelog) Alert(v interface{}, r ...interface{}) {
	this.toLog("alert", v)
}

func (this *Beelog) Critical(v interface{}, r ...interface{}) {
	this.toLog("critical", v)
}
func (this *Beelog) Error(v interface{},r ...interface{}) {
	this.toLog("error", v)

}
func (this *Beelog) Warning(v interface{}, r ...interface{}) {
	this.toLog("warning", v)

}
func (this *Beelog) Notice(v interface{},r ...interface{}) {
	this.toLog("notice", v)

}
func (this *Beelog) Info(v interface{},r ...interface{}) {
	this.toLog("Info", v)

}
func (this *Beelog) Debug(v interface{},r ...interface{}) {
	this.toLog("debug", v)

}

func (this *Beelog) Trace(v interface{},r ...interface{}) {
	this.toLog("trace", v)

}

//Log 输出日志
func (this *Beelog)  toLog(level, v interface{}) {
	format := "%s"

	if this.runmode == "DEBUG" {
		switch level {
		case "emergency":
			this.consoleLogs.Emergency(format, v)
		case "alert":
			this.consoleLogs.Alert(format, v)
		case "critical":
			this.consoleLogs.Critical(format, v)
		case "error":
			this.consoleLogs.Error(format, v)
		case "warning":
			this.consoleLogs.Warning(format, v)
		case "notice":
			this.consoleLogs.Notice(format, v)
		case "info":
			this.consoleLogs.Info(format, v)
		case "debug":
			this.consoleLogs.Debug(format, v)
		case "trace":
			this.consoleLogs.Trace(format, v)
		default:
			this.consoleLogs.Debug(format, v)
		}
	}

	switch level {
	case "emergency":
		this.FileLogs.Emergency(format, v)
	case "alert":
		this.FileLogs.Alert(format, v)
	case "critical":
		this.FileLogs.Critical(format, v)
	case "error":
		this.FileLogs.Error(format, v)
	case "warning":
		this.FileLogs.Warning(format, v)
	case "notice":
		this.FileLogs.Notice(format, v)
	case "info":
		this.FileLogs.Info(format, v)
	case "debug":
		this.FileLogs.Debug(format, v)
	case "trace":
		this.FileLogs.Trace(format, v)
	default:
		this.FileLogs.Debug(format, v)
	}

}
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main()  {
	logger := InitLogs("","logs",Beelog{})
	logger.Warning("warning")
}