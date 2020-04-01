package rabbitmq

import (
	"fmt"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/calendar/event"
	"gopkg.in/yaml.v2"
	"testing"
	"time"
)

func TestSendAndReciveMessage(t *testing.T) {
	config := Config{
		User:     "guest",
		Pass:     "guest",
		HostPort: "192.168.1.124:5672",
		Timeout:  10,
		Queue1:   "sendEvent",
	}

	r, err := NewRMQ(config, nil)
	if err != nil {
		t.Error("Fail to create new RabbitMQ", err.Error())
	}
	eSource := make(map[int]event.Event)
	//send 10 event to RabbitMQ
	dateStart := time.Date(2020, 1, 1, 11, 0, 0, 0, time.UTC)
	dateEnd := time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC)
	for i := 0; i < 10; i++ {
		dateStart.AddDate(0, 0, 1)
		dateEnd.AddDate(0, 0, 1)
		title := fmt.Sprintf("Event %v", i)
		note := fmt.Sprintf("Event %v  start:%v finish: %v", i, dateStart.Format(time.RFC3339), dateEnd.Format(time.RFC3339))
		e, _ := event.CreateEvent(dateStart.Format(time.RFC3339), dateEnd.Format(time.RFC3339), title, note, 0, 0)
		d, err := yaml.Marshal(&e)
		eSource[i] = e
		if err != nil {
			t.Error("Fail to marshal event ", err.Error())
		}
		err = r.Send(d)
		if err != nil {
			t.Error("Fail to send message by RabbitMQ", err.Error())
		}
	}
	eQueue := make(map[int]event.Event)
	for i := 0; i < 10; i++ {
		msgs, ok, err := r.ch.Get(config.Queue1, true)
		if ok {

			if err != nil {
				t.Error("Fail to send message by RabbitMQ", err.Error())
			}
			e := event.Event{}
			yaml.Unmarshal(msgs.Body, &e)
			eQueue[i] = e
		}
	}
	//Check recive message by RabbitMQ
	for i := 0; i < 10; i++ {
		if eQueue[i].StartTime != eSource[i].StartTime ||
			eQueue[i].EndTime != eSource[i].EndTime ||
			eQueue[i].Title != eSource[i].Title ||
			eQueue[i].Note != eSource[i].Note ||
			eQueue[i].Duration != eSource[i].Duration ||
			eQueue[i].TypeDuration != eSource[i].TypeDuration {
			t.Error("Event in storage")
		}
	}
}
