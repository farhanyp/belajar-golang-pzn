package loggingtofile

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogger(t *testing.T){

	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)

	file, _ := os.OpenFile("application.log", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0666)
	logger.SetOutput(file)

	logger.Trace("ini Trace")
	logger.Debug("ini Debug")
	logger.Info("ini Info")
	logger.Warn("ini Warn")

}