package log

type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Infof(format string, args ...interface{})
}

type Level string

const (
	LevelDebug = iota
	LevelError
	LevelInfo
)
