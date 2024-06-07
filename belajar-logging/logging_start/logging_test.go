package loggingstart

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogger(t *testing.T){

	logger := logrus.New()

	logger.Println("Ini logrus sebagai logging")

}