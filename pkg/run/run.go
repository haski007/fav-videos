package run

import "github.com/sirupsen/logrus"

type Args struct {
	LogLevel logrus.Level
}

func LogLevel(level string) logrus.Level {
	switch level {
	case "DEBUG":
		return logrus.DebugLevel
	case "INFO":
		return logrus.InfoLevel
	case "WARN":
		return logrus.WarnLevel
	case "ERROR":
		return logrus.ErrorLevel
	}

	logrus.Errorf(`unknown log level %q, using "ERROR" by default`, level)

	return logrus.ErrorLevel
}
