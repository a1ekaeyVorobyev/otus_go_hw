package scheduler

import (
	"fmt"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/calendar/event"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/rabbitmq"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/storage"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	_ "gopkg.in/yaml.v2"
	"sync"
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
	fmt.Println("create config")
	var wg sync.WaitGroup
	//pg := storage.Postgres{Config: conf, Logger: &logrus.Logger{}}
	pg,err := storage.NewPG(conf,&logrus.Logger{})
	if err!=nil{
		t.Error("Fail create storage", err.Error())
	}
	conf := Config{
		CheckInSeconds: 10,
		NotifyInMinute: 10,
		ForceCloseInMinute : 5,
	}

	configRMQ := rabbitmq.Config{
		User:     "guest",
		Pass:     "guest",
		HostPort: "192.168.1.31:5672",
		Timeout:  10,
		Queue1:   "sendEventTest",
		Queue2:   "reciveEventTest",
	}

	s,err := NewScheduler(pg, &logrus.Logger{},conf,configRMQ)
	if err!=nil{
		t.Error("Fail create new scheduler", err.Error())
	}
	fmt.Println("add 10 events")
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
	dateFinish := time.Now().Add(time.Duration(s.config.NotifyInMinute) * time.Minute)
	fmt.Println("GetEventSending",dateFinish)
	events, err := s.store.GetEventSending(dateFinish)
	//events, err := pg.GetEventSending(dateFinish)
	if err != nil {
		t.Error("Fail get event RabbitMQ", err.Error())
	}
	fmt.Println("endEventsToQueue")
	wg.Add(1)
	go func() {
		defer wg.Done()
		s.sendEventsToQueue()
	}()
	wg.Wait()
	//work emulation sender
	r, err := rabbitmq.NewRMQ(s.configRMQ, s.logger)
	if err != nil {
		t.Error("Fail to create new RabbitMQ by scheduler", err.Error())
	}
	i := 0
out:
	for {
		msgs, ok, err := r.GetChanel(s.configRMQ.Queue1)
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
	s.logger.Infoln("Start scheduler")
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
