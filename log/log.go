package log

import (
	"github.com/astaxie/beego/logs"
)

var logger *logs.BeeLogger

func init() {
	logger = logs.NewLogger(10000)
	logger.Async()
	logger.SetLogger("console", "")
	logger.EnableFuncCallDepth(true)
}

func Logger() *logs.BeeLogger {
	return logger
}
