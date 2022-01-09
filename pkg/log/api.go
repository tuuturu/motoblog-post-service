package log

import (
	logrusWrapper "github.com/deifyed/post-service/pkg/log/logrus"
	"github.com/sirupsen/logrus"
)

func GetLogger(component string, activity string) Logger {
	level, _ := logrus.ParseLevel("info")

	logger := logrusWrapper.New(level, component, activity)

	return logger
}

func (receiver Level) String() string {
	return string(receiver)
}
