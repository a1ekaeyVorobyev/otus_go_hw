package scheduler

import (
	"fmt"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/calendar/event"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/rabbitmq"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/storage"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	_ "gopkg.in/yaml.v2"
	"testing"
	"time"
)

var conf = storage.Config{
	Server:         "127.0.0.1:5432",
	User:           "postgres",
	Pass:           "postgres",
	Database:       "calendar",
	TimeoutConnect: 10,
}

func TestAddGetAllGetDel(t *testing.T) {
	//var wg sync.WaitGroup
	pg := storage.Postgres{Config: conf, Logger: &logrus.Logger{}}
	pg.Init()
	s := Scheduler{
		Store:     &pg,
		Logger:    &logrus.Logger{},
		Config:    Config{},
		ConfigRMQ: rabbitmq.Config{},
		Done:      nil,
	}

	s.Config = Config{
		CheckInSeconds: 10,
		NotifyInMinute: 10,
	}

	s.ConfigRMQ = rabbitmq.Config{
		User:     "guest",
		Pass:     "guest",
		HostPort: "192.168.1.124:5672",
		Timeout:  10,
		Queue1:   "sendEventTest",
		Queue2:   "reciveEventTest",
	}

	//add 10 events
	dateStart := time.Now()
	dateEnd := time.Now().Add(time.Duration(5) * time.Minute)
	for i := 0; i < 10; i++ {
		dateStart.Add(time.Duration(6) * time.Minute)
		dateEnd.Add(time.Duration(6) * time.Minute)
		title := fmt.Sprintf("Event %v", i)
		note := fmt.Sprintf("Event %v  start:%v finish: %v", i, dateStart.Format(time.RFC3339), dateEnd.Format(time.RFC3339))
		event, _ := event.CreateEvent(dateStart.Format(time.RFC3339), dateEnd.Format(time.RFC3339), title, note, 0, 0)
		err := pg.Add(event)
		if err != nil {
			t.Error("Fail to add event to storage", err.Error())
		}
	}
	dateFinish := time.Now().Add(time.Duration(s.Config.NotifyInMinute) * time.Minute)
	events, err := s.Store.GetEventSending(dateFinish)
	if err != nil {
		t.Error("Fail get event RabbitMQ", err.Error())
	}
	s.sendEventsToQueue()
	//work emulation sender
	r, err := rabbitmq.NewRMQ(s.ConfigRMQ, s.Logger)
	if err != nil {
		t.Error("Fail to create new RabbitMQ by scheduler", err.Error())
	}
	i := 0
out:
	for {
		msgs, ok, err := r.GetChanel(s.ConfigRMQ.Queue1)
		if !ok {
			break out
		}
		if err != nil {
			t.Error("Fail to send message by RabbitMQ", err.Error())
		}
		e := event.Event{}
		yaml.Unmarshal(msgs.Body, &e)
		i++
	}
	if i != len(events) {
		t.Error("Fail count message to send message by RabbitMQ")
	}
	//send nessage from 2 queue
	if err != nil {
		t.Error("Fail to create new RabbitMQ by scheduler", err.Error())
	}
	s.Logger.Infoln("Start scheduler")
	for _, e := range events {
		d, err := yaml.Marshal(&e)
		if err != nil {
			t.Error("Fail to marshal event by scheduler", err.Error())
		}
		err = r.Send2(d)
		if err != nil {
			t.Error("Fail to send message by RabbitMQ from scheduler", err.Error())
		}
	}
	//сheck
	s.markEvent()
	for _, e := range events {
		ev, err := pg.Get(e.Id)
		if err != nil {
			t.Error("Fail get event RabbitMQ", err.Error())
		}
		if ev.Issending != 2 {
			t.Error("Fail get event RabbitMQ, Issending", ev.Issending)
		}
	}
}
