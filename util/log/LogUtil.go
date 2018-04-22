package log

import (
	"github.com/astaxie/beego/logs"
)

var consoleLog *logs.BeeLogger;
var fileInfoLog *logs.BeeLogger;
var warnInfoLog *logs.BeeLogger;
var errInfoLog *logs.BeeLogger;


/**
	console级别是只在控制台输出的日志
 */
func Console(info string){
	consoleLog.Info(info)
}

/**
	Info级别会在日志文件记录的日志
 */
func Info(info string){
	fileInfoLog.Info(info)
}

/**
	Warn级别会在日志文件记录的日志,都是警告
 */
func Warn(info string){
	warnInfoLog.Warn(info)
}

/**
	Error级别会在日志文件记录的日志,都是报错
 */
func Error(info string){
	errInfoLog.Error(info)
}


func init() {
	consoleLog = logs.NewLogger()
	fileInfoLog = logs.NewLogger()
	warnInfoLog = logs.NewLogger()
	errInfoLog = logs.NewLogger()

	//输出log的文件名称以及行号
	consoleLog.EnableFuncCallDepth(true)
	fileInfoLog.EnableFuncCallDepth(true)
	warnInfoLog.EnableFuncCallDepth(true)
	errInfoLog.EnableFuncCallDepth(true)
	//封装了调用 log 包,那么需要设置 SetLogFuncCallDepth,默认是 2
	consoleLog.SetLogFuncCallDepth(3)
	fileInfoLog.SetLogFuncCallDepth(3)
	warnInfoLog.SetLogFuncCallDepth(3)
	errInfoLog.SetLogFuncCallDepth(3)
	//设置异步输出
	fileInfoLog.Async()//普通日志以及警告日志用异步来输出，提高效率，报错日志不能异步，防止丢失
	warnInfoLog.Async()
	//设置缓冲Chan大小，异步才需要缓冲
	fileInfoLog.Async(1e3)
	warnInfoLog.Async(1e3)

	//设置日志输出文件
	consoleLog.SetLogger(logs.AdapterConsole)
	fileInfoLog.SetLogger(logs.AdapterFile,`{"filename":"outfile/loginfo.log","maxdays":30}`)
	warnInfoLog.SetLogger(logs.AdapterFile,`{"filename":"outfile/logWarn.log","maxdays":90}`)
	errInfoLog.SetLogger(logs.AdapterFile,`{"filename":"outfile/logError.log","maxdays":365}`)
}