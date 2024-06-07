package logginglevel

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogger(t *testing.T){

	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)

	logger.Trace("ini Trace")
	logger.Debug("ini Debug")
	logger.Info("ini Info")
	logger.Warn("ini Warn")
	logger.Panic("ini Panic")

}