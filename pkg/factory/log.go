package factory

import log "github.com/sirupsen/logrus"

func InitLog(level log.Level) {
	log.SetLevel(level)
	log.SetFormatter(&log.JSONFormatter{})
}
