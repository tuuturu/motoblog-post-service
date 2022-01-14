package logrus

import (
	"github.com/sirupsen/logrus"
)

func (receiver *Logger) Debug(args ...interface{}) {
	receiver.entry.Debug(args...)
}

func (receiver *Logger) Debugf(format string, args ...interface{}) {
	receiver.entry.Debugf(format, args...)
}

func (receiver *Logger) Errorf(format string, args ...interface{}) {
	receiver.entry.Errorf(format, args...)
}

func (receiver *Logger) Infof(format string, args ...interface{}) {
	receiver.entry.Errorf(format, args...)
}

func New(level logrus.Level, component string, activity string) *Logger {
	ll := logrus.New()

	ll.SetFormatter(&logrus.JSONFormatter{})
	ll.Level = level

	return &Logger{
		entry: ll.WithFields(logrus.Fields{
			"component": component,
			"activity":  activity,
		}),
	}
}
