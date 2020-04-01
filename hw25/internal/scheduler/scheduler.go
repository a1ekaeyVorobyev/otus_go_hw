package scheduler

import (
	"github.com/sirupsen/logrus"
	"time"
)

type Config struct {
	CheckInSeconds  int `yaml:"checkInSeconds"`
	NotifyInSeconds int `yaml:"notifyInSeconds"`
}

type Interface interface {
	GetEventSending() error
	MarkEventSentToQueue(int) error
	MarkEventSentToSubScribe(int)error
}

func Run(done chan bool, config Config,logger *logrus.Logger){
	ticker := time.NewTicker(time.Duration(config.CheckInSeconds) * time.Second)
out:
	for {
			select {
			case <-done:
				break out
			case  <-ticker.C:
				go sendEventsToQueue(logger)
				go markEvent(logger)
		}
	}
}

func sendEventsToQueue(logger *logrus.Logger){
	logger.Infoln("Start scheduler")
}

func markEvent(logger *logrus.Logger){
	logger.Infoln("Start scheduler")
}