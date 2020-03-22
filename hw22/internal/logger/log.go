package logger

import (
	"fmt"
	"os"

	"github.com/a1ekaeyVorobyev/otus_go_hw/hw22/internal/config"
	"github.com/sirupsen/logrus"
)

func GetLogger(c config.Config) (logrus.Logger, *os.File) {
	logger := logrus.Logger{}
	switch c.LogLevel {
	case "info":
		logger.SetLevel(logrus.InfoLevel)
	case "debug":
		fallthrough
	default:
		logger.SetLevel(logrus.DebugLevel)
	}
	formatter := logrus.JSONFormatter{}
	logger.SetFormatter(&formatter)
	if c.FileName != "" {
		f, err := os.OpenFile(c.FileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			fmt.Printf("error opening file: %v", err)
		}
		logger.SetOutput(f)
		return logger, f
	} else {
		logger.SetOutput(os.Stdout)
	}
	return logger, nil
}
