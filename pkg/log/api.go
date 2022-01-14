package log

import (
	"os"

	"github.com/sirupsen/logrus"
	logrusWrapper "github.com/tuuturu/motoblog-post-service/pkg/log/logrus"
)

func GetLogger(component string, activity string) Logger {
	raw := os.Getenv("LOG_LEVEL")

	if raw == "" {
		raw = "info"
	}

	level, err := logrus.ParseLevel(raw)
	if err != nil {
		level = logrus.InfoLevel
	}

	logger := logrusWrapper.New(level, component, activity)

	return logger
}

func (receiver Level) String() string {
	return string(receiver)
}
