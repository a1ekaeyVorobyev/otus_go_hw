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
	ForceCloseInMinute	int `yaml:"forceCloseInMinute"`
}

type scheduler struct {
	store     pkg.Scheduler
	logger    *logrus.Logger
	config    Config
	configRMQ rabbitmq.Config
	done      chan bool
	sync.Mutex
	x  int
}

func NewScheduler(store pkg.Scheduler, logger *logrus.Logger,config Config,configRMQ rabbitmq.Config)(s *scheduler,err error){
	s = &scheduler{}
	s.config = config
	s.configRMQ = configRMQ
	s.logger = logger
	return s,nil
}

func (s *scheduler) Run() {
	fmt.Println("run Scheduler")
	ticker := time.NewTicker(time.Duration(s.config.CheckInSeconds) * time.Second)
out:
	for {
		select {
		case <-s.done:
			break out
		case <-ticker.C:
			s.Lock()
			if s.x ==0 {
				s.x= 2
				go s.sendEventsToQueue()
				go s.markEvent()
			}
			s.Unlock()
		}
	}
	ticker.Stop()
}

func (s *scheduler) sendEventsToQueue() {
	//defer s.Wg.Done()
	dateFinish := time.Now().Add(time.Duration(s.config.NotifyInMinute) * time.Minute)
	events, err := s.store.GetEventSending(dateFinish)
	if err != nil {
		s.logger.Error("Fail to get events by scheduler:", err.Error())
	}
	if events == nil {
		return
	}
	r, err := rabbitmq.NewRMQ(s.configRMQ, s.logger)
	if err != nil {
		s.logger.Error("Fail to create new RabbitMQ by scheduler", err.Error())
		return
	}
	s.logger.Infoln("Start scheduler")
	for _, e := range events {
		d, err := yaml.Marshal(&e)
		if err != nil {
			s.logger.Error("Fail to marshal event by scheduler", err.Error())
		}
		err = r.Send(d)
		if err != nil {
			s.logger.Error("Fail to send message by RabbitMQ from scheduler", err.Error())
			continue
		}
		err = s.store.MarkEventSentToQueue(e.Id)
		if err != nil {
			s.logger.Error("Fail to mark event sending  by RabbitMQ from scheduler", err.Error())
			continue
		}
	}
	s.Lock()
	s.x--
	s.Unlock()
}

func (s *scheduler) markEvent() {
	r, err := rabbitmq.NewRMQ(s.configRMQ, s.logger)
	if err != nil {
		s.logger.Error("Fail to create new RabbitMQ by scheduler", err.Error())
		return
	}
out:
	for {
		msgs, ok, err := r.GetChanel(s.configRMQ.Queue2)
		if !ok {
			break out
		}
		if err != nil {
			s.logger.Error("Fail to send message by RabbitMQ", err.Error())
		}
		e := event.Event{}
		yaml.Unmarshal(msgs.Body, &e)
		err = s.store.MarkEventSentToSubScribe(e.Id)
		if err != nil {
			s.logger.Error("Fail to mark event sending  by RabbitMQ from storage", err.Error())
			continue

		}
	}
	s.Lock()
	s.x--
	s.Unlock()
}

func (s *scheduler) ShutDown() {
	t := time.NewTimer(time.Minute * time.Duration(s.config.ForceCloseInMinute))
	select{
		case <-t.C:
			s.logger.Info("Forced close scheduler")
			return
	default:
		s.Lock()
		if s.x ==0 {
			s.done <- true
			return
		}
		s.Unlock()
		time.Sleep(time.Duration(5)*time.Second)
	}
}
