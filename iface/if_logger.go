package iface

type Logger interface {
	Log(prefix, format string, v ...interface{})

	// Log level
	Debug(format string, v ...interface{})
	Info(format string, v ...interface{})
	Warn(format string, v ...interface{})
	Error(format string, v ...interface{})
	Panic(format string, v ...interface{})
}
