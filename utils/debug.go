package utils

import (
	"os"

	"github.com/Sirupsen/logrus"
)

func EnableDebug() {
	os.Setenv("DEBUG", "1")
	logrus.SetLevel(logrus.DebugLevel)
}
