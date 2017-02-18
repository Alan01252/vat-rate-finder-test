package logger

import (
	"os"

	"flag"

	log "github.com/Sirupsen/logrus"
)

var requestLogger *log.Entry

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	if flag.Lookup("test.v") != nil {
		log.SetLevel(log.ErrorLevel)
	} else {
		log.SetLevel(log.DebugLevel)
	}
	requestLogger = log.WithFields(log.Fields{"package": "vat"})
}

func GetLogger(packageName string) *log.Entry {
	requestLogger = requestLogger.WithFields(log.Fields{"package": packageName})
	return requestLogger
}
