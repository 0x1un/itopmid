package support

import (
	"fmt"
	"runtime"
	"time"

	"github.com/0x1un/itopmid/util"
	"github.com/fatih/color"
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

func (self *ItopMidLogger) Log(prefix, format string, v ...interface{}) string {
	content := fmt.Sprintf(format+"\r\n", v...)
	// 在这里获取当前的调用行数
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "??unkown??"
		line = 0
	}
	// prefix + time of now + file path : current line + content
	consoLog := fmt.Sprintf(
		"%s%s %s:%d %s",
		prefix,
		util.FormatTimeToString(time.Now()),
		util.GetCurrentPath(file),
		line,
		content)
	return consoLog
}

func (self *ItopMidLogger) Debug(format string, v ...interface{}) {
	c := color.New(color.FgCyan, color.Bold)
	c.Print(self.Log(DEBUG, format, v...))
}

func (self *ItopMidLogger) Info(format string, v ...interface{}) {
	c := color.New(color.FgGreen, color.Bold)
	c.Print(self.Log(INFO, format, v...))
}

func (self *ItopMidLogger) Warn(format string, v ...interface{}) {
	c := color.New(color.FgYellow, color.Bold)
	c.Print(self.Log(WARN, format, v...))
}

func (self *ItopMidLogger) Error(format string, v ...interface{}) {
	c := color.New(color.FgRed, color.Bold)
	c.Print(self.Log(ERROR, format, v...))
}

func (self *ItopMidLogger) Panic(format string, v ...interface{}) {
	c := color.New(color.FgHiRed, color.Bold)
	c.Print(self.Log(PANIC, format, v...))
	panic(c.Sprintf(format, v...))
}
