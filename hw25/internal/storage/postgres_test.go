package storage

import (
	"fmt"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/calendar/event"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

var conf = Config{
	Server:         "127.0.0.1:5432",
	User:           "postgres",
	Pass:           "postgres",
	Database:       "calendar",
	TimeoutConnect: 10,
}


func TestAddGetAllGetDel(t *testing.T) {
	pg := Postgres{Config: conf, Logger: &logrus.Logger{}}
	err := pg.Init()
	if err != nil {
		t.Error("Fail to init PG")
	}
	defer pg.Shutdown()
	//add record
	dateStart := time.Date(2020, 1, 1, 11, 0, 0, 0, time.UTC)
	dateEnd := time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC)
	for i := 0; i < 10; i++ {
		dateStart.AddDate(0, 0, 1)
		dateEnd.AddDate(0, 0, 1)
		title := fmt.Sprintf("Event %v", i)
		note := fmt.Sprintf("Event %v  start:%v finish: %v", i, dateStart.Format(time.RFC3339), dateEnd.Format(time.RFC3339))
		event, _ := event.CreateEvent(dateStart.Format(time.RFC3339), dateEnd.Format(time.RFC3339), title, note, 0, 0)
		err = pg.Add(event)
		if err != nil {
			t.Error("Fail to add event to storage", err.Error())
		}
	}

	eventsFromPg, err := pg.GetAll()
	if err != nil {
		t.Error("Fail to get all events in storage")
	}
	cnt := pg.CountRecord()
	if cnt < 1 {
		t.Error("Fail to get count events in storage")
	}
	//compare
	for _, v := range eventsFromPg {
		e, err := pg.Get(v.Id)
		if err != nil {
			t.Error("Fail to get record in storage id:", v.Id)
			continue
		}
		if e.Title != v.Title || e.TypeDuration != v.TypeDuration || e.EndTime != v.EndTime ||
			e.Duration != v.Duration || e.StartTime != v.StartTime || e.Note != v.Note {
			t.Error("Fail to compare record in storage id:", v.Id)
		}
	}
	//edit record
	for _, v := range eventsFromPg {
		v.Title = fmt.Sprintf("Новый - %v", v.Title)
		err := pg.Edit(v)
		if err != nil {
			t.Error("Fail to edit record in storage id:", v.Id)
		}
	}
	//delete all record
	for _, v := range eventsFromPg {
		err := pg.Delete(v.Id)
		if err != nil {
			t.Error("Fail to get record in storage id:", v.Id)

		}
	}
	cnt = pg.CountRecord()
	if cnt != 0 {
		t.Error("Fail to get count events in storage")
	}

}
