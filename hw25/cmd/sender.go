package main

import (
	"flag"
	"fmt"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/calendar/event"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/config"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/logger"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/rabbitmq"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"os"
	"os/signal"
	"syscall"
)

type sender struct {
	rmq    *rabbitmq.RMQ
	logger *logrus.Logger
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	var configFile string
	flag.StringVar(&configFile, "config", "config/config.yaml", "Config file")
	flag.Parse()
	if configFile == "" {
		_, _ = fmt.Fprint(os.Stderr, "don't config file")
		os.Exit(2)
	}
	conf, err := config.ReadFromFile(configFile)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err)
		os.Exit(2)
	}
	fmt.Println(conf)

	logger, f := logger.GetLogger(conf.Log)
	if f != nil {
		defer f.Close()
	}

	r, err := rabbitmq.NewRMQ(conf.Rmq, &logger)
	s := sender{r, &logger}
	if err != nil {
		logger.Error("Fail to create new RabbitMQ by scheduler", err.Error())
	}

	msgCh, err := r.GetMsgsCh()
	if err != nil {
		logger.Error("Fail to create new chanel RabbitMQ ", err.Error())
	}

exit:
	for {
		select {
		case message, ok := <-msgCh:
			if ok {
				go s.processMessage(message.Body)
			}
		case c := <-sigs:
			logger.Infof("Got signal: %v. Exit.", c)
			break exit
		}
	}
}

func (s *sender) processMessage(message []byte) {
	e := event.Event{}
	err := yaml.Unmarshal(message, &e)
	if err != nil {
		s.logger.Error("Error Marshal in sender")
	}
	fmt.Println("You get next event:", e)
	err = s.rmq.Send2(message)
	if err != nil {
		s.logger.Error("Fail to send message by RabbitMQ from scheduler", err.Error())
	}
}
