package logger

import (
	"os"

	log "github.com/Sirupsen/logrus"
)

var requestLogger *log.Entry

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	requestLogger = log.WithFields(log.Fields{"package": "vat"})
}

func GetLogger(packageName string) *log.Entry {
	requestLogger = requestLogger.WithFields(log.Fields{"package": packageName})
	return requestLogger
}
