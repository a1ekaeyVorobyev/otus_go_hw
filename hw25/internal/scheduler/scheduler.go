package scheduler

import (
	"fmt"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/calendar/event"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/pkg"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/rabbitmq"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"sync"
	"time"
)

type Config struct {
	CheckInSeconds int `yaml:"checkInSeconds"`
	NotifyInMinute int `yaml:"notifyInMinute"`
}

type Scheduler struct {
	Store     pkg.Scheduler
	Logger    *logrus.Logger
	Config    Config
	ConfigRMQ rabbitmq.Config
	Done      chan bool
}


func (s *Scheduler) Run() {
	var wg sync.WaitGroup
	fmt.Println("run Scheduler")
	ticker := time.NewTicker(time.Duration(s.Config.CheckInSeconds) * time.Second)
out:
	for {
		select {
		case <-s.Done:
			break out
		case <-ticker.C:
			fmt.Println("run ticker")
			wg.Add(2)
			go func() {
				defer wg.Done()
				s.sendEventsToQueue()
			}()
			go func() {
				defer wg.Done()
				s.markEvent()
			}()
		}
	}
}

func (s *Scheduler) sendEventsToQueue() {
	//defer s.Wg.Done()
	dateFinish := time.Now().Add(time.Duration(s.Config.NotifyInMinute) * time.Minute)
	events, err := s.Store.GetEventSending(dateFinish)
	if err != nil {
		s.Logger.Error("Fail to get events by scheduler:", err.Error())
	}
	if events == nil {
		return
	}
	r, err := rabbitmq.NewRMQ(s.ConfigRMQ, s.Logger)
	if err != nil {
		s.Logger.Error("Fail to create new RabbitMQ by scheduler", err.Error())
		return
	}
	s.Logger.Infoln("Start scheduler")
	for _, e := range events {
		d, err := yaml.Marshal(&e)
		if err != nil {
			s.Logger.Error("Fail to marshal event by scheduler", err.Error())
		}
		err = r.Send(d)
		if err != nil {
			s.Logger.Error("Fail to send message by RabbitMQ from scheduler", err.Error())
			continue
		}
		err = s.Store.MarkEventSentToQueue(e.Id)
		if err != nil {
			s.Logger.Error("Fail to mark event sending  by RabbitMQ from scheduler", err.Error())
			continue
		}
	}
}

func (s *Scheduler) markEvent() {
	//defer s.Wg.Done()
	r, err := rabbitmq.NewRMQ(s.ConfigRMQ, s.Logger)
	if err != nil {
		s.Logger.Error("Fail to create new RabbitMQ by scheduler", err.Error())
		return
	}
out:
	for {
		msgs, ok, err := r.GetChanel(s.ConfigRMQ.Queue2)
		if !ok {
			break out
		}
		if err != nil {
			s.Logger.Error("Fail to send message by RabbitMQ", err.Error())
		}
		e := event.Event{}
		yaml.Unmarshal(msgs.Body, &e)
		err = s.Store.MarkEventSentToSubScribe(e.Id)
		if err != nil {
			s.Logger.Error("Fail to mark event sending  by RabbitMQ from storage", err.Error())
			continue

		}
	}
}

func (s *Scheduler) ShutDown() {
	s.Done <- true
}
