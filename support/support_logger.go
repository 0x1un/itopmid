package support

import (
	"fmt"
	"runtime"
	"time"

	"github.com/0x1un/itopmid/util"
)

const (
	DEBUG = "[DEBUG] "
	WARN  = "[WARN] "
	INFO  = "[INFO] "
	ERROR = "[ERROR] "
	PANIC = "[PANIC] "
)

type ItopMidLogger struct {
	//	goLogger *log.Logger
}

func (self *ItopMidLogger) Log(prefix, format string, v ...interface{}) {
	content := fmt.Sprintf(format+"\r\n", v...)
	// 在这里获取当前的调用行数
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "??unkown??"
		line = 0
	}
	// prefix + time of now + file path : current line + content
	fmt.Printf(
		"%s%s %s:%d %s",
		prefix,
		util.FormatTimeToString(time.Now()),
		util.GetCurrentPath(file),
		line,
		content)
}

func (self *ItopMidLogger) Debug(format string, v ...interface{}) {
	self.Log(DEBUG, format, v...)
}

func (self *ItopMidLogger) Info(format string, v ...interface{}) {
	self.Log(INFO, format, v...)
}

func (self *ItopMidLogger) Warn(format string, v ...interface{}) {
	self.Log(WARN, format, v...)
}

func (self *ItopMidLogger) Error(format string, v ...interface{}) {
	self.Log(ERROR, format, v...)
}

func (self *ItopMidLogger) Panic(format string, v ...interface{}) {
	self.Log(PANIC, format, v...)
}
