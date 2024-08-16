package logger

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var log *logrus.Logger

func InitLogger(logFilePath string, logLevel string) error {
	log = logrus.New()

	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	multiWriter := io.MultiWriter(os.Stdout, file)
	log.SetOutput(multiWriter)
	log.SetFormatter(&logrus.JSONFormatter{})

	if logLevel == "debug" {
		log.SetLevel(logrus.DebugLevel)
	} else if logLevel == "info" {
		log.SetLevel(logrus.InfoLevel)
	} else if logLevel == "error" {
		log.SetLevel(logrus.ErrorLevel)
	} else {
		log.SetLevel(logrus.InfoLevel)
	}

	return nil
}

func GetLogger() *logrus.Logger {
	return log
}
