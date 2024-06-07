package hook

import (
	"fmt"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

type SampleHook struct{

}

func (s *SampleHook) Levels() []logrus.Level{
	return []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel}
}

func (s *SampleHook) Fire(entry *logrus.Entry) error{
	fmt.Println("Ini hook", entry.Message, entry.Level)

	return nil
}

func TestLogger(t *testing.T){

	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.AddHook(&SampleHook{})

	file, _ := os.OpenFile("application.json", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0666)
	logger.SetOutput(file)

	logger.Trace("ini Trace")
	logger.WithField("username","farhan").Debug("ini Debug")
	logger.WithFields(logrus.Fields{
		"password": "1234556",
	}).Info("ini Info")
	logger.Warn("ini Warn")

}