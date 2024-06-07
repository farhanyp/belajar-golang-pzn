package loggingaddfield

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogger(t *testing.T){

	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})

	file, _ := os.OpenFile("application.json", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0666)
	logger.SetOutput(file)

	logger.Trace("ini Trace")
	logger.WithField("username","farhan").Debug("ini Debug")
	logger.WithFields(logrus.Fields{
		"password": "1234556",
	}).Info("ini Info")
	logger.Warn("ini Warn")

}