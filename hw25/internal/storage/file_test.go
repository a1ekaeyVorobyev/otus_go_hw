package storage

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/calendar/event"
)

func TestSaveAndLoadFile(t *testing.T) {
	//InFile := InFile{}
	InFile,_ := NewStorage()
	fmt.Print("create storage")
	InFile.Clear()
	dateStart := time.Date(2020, 1, 1, 11, 0, 0, 0, time.UTC)
	dateEnd := time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC)
	for i := 0; i < 10; i++ {
		dateStart.AddDate(0, 0, 1)
		dateEnd.AddDate(0, 0, 1)
		title := "Event " + string(i)
		note := "This envets" + string(i) + " start:" + dateStart.Format(time.RFC3339) + " finish:" + dateEnd.Format(time.RFC3339)
		event, _ := event.CreateEvent(dateStart.Format(time.RFC3339), dateEnd.Format(time.RFC3339), title, note, 0, 0)
		err := InFile.Add(event)
		if err != nil {
			t.Error("Can't add event")
		}
	}
	events, err := InFile.GetAll()
	len := len(events)
	if err != nil || len != 10 {
		t.Error("In storage not 10 event")
	}
	InFile.SaveEvents()
	InFile.Clear()
	InFile.loadEvents()
	if InFile.CountRecord() != len {
		t.Error("Number of records are not equal")
	}

}

func TestEmptyStorageHaveNoEvents(t *testing.T) {
	InFile,err := NewStorage()
	InFile.Clear()
	events, err := InFile.GetAll()
	if err != nil || len(events) != 0 {
		t.Error("In new storage have events")
	}
}

func TestAddEventSuccess(t *testing.T) {
	InFile,err := NewStorage()
	InFile.Clear()
	event, _ := event.CreateEvent("2020-01-02T11:00:00Z", "2020-01-02T12:00:00Z", "Event 1", "Start event", 0, 0)
	err = InFile.Add(event)
	if err != nil {
		t.Error("Can't add event ")
	}

	events, err := InFile.GetAll()
	if err != nil || len(events) != 1 {
		t.Error("In storage not 1 event")
	}
}

func TestDeleteEventSuccess(t *testing.T) {
	InFile,_ := NewStorage()
	InFile.Clear()
	dateStart := time.Date(2020, 1, 1, 11, 0, 0, 0, time.UTC)
	dateEnd := time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC)
	for i := 0; i < 10; i++ {
		dateStart.AddDate(0, 0, 1)
		dateEnd.AddDate(0, 0, 1)
		title := "Event " + string(i)
		note := "This envets" + string(i) + " start:" + dateStart.Format(time.RFC3339) + " finish:" + dateEnd.Format(time.RFC3339)
		event, _ := event.CreateEvent(dateStart.Format(time.RFC3339), dateEnd.Format(time.RFC3339), title, note, 0, 0)
		err := InFile.Add(event)
		if err != nil {
			t.Error("Can't add event to storage")
		}
	}
	events, _ := InFile.GetAll()
	for _, v := range events {
		err := InFile.Delete(v.Id)
		if err != nil {
			t.Error("Can't delete event")
		}
	}
	if InFile.CountRecord() != 0 {
		t.Error("In storage have events")
	}
}

func TestIslBusy(t *testing.T) {
	InFile,_ := NewStorage()
	InFile.Clear()
	dateStart := time.Date(2020, 1, 1, 11, 0, 0, 0, time.UTC)
	dateEnd := time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC)
	for i := 0; i < 10; i++ {
		dateStart.AddDate(0, 0, 1)
		dateEnd.AddDate(0, 0, 1)
		title := "Event " + string(i)
		note := "This envets" + string(i) + " start:" + dateStart.Format(time.RFC3339) + " finish:" + dateEnd.Format(time.RFC3339)
		event, _ := event.CreateEvent(dateStart.Format(time.RFC3339), dateEnd.Format(time.RFC3339), title, note, 0, 0)
		err := InFile.Add(event)
		if err != nil {
			t.Error("Can't add event")
		}
	}
	events, _ := InFile.GetAll()
	for _, v := range events {
		ok, err := InFile.IsBusy(v)
		if ok {
			t.Error("This events have in memory")
		}
		if err != nil {
			t.Error(err.Error())
		}
	}
}

func TestGetEvent(t *testing.T) {
	InFile,err := NewStorage()
	InFile.Clear()
	dateStart := time.Date(2020, 1, 1, 11, 0, 0, 0, time.UTC)
	dateEnd := time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC)
	title := "Event 1"
	note := "This envets1 start:" + dateStart.Format(time.RFC3339) + " finish:" + dateEnd.Format(time.RFC3339)
	event, _ := event.CreateEvent(dateStart.Format(time.RFC3339), dateEnd.Format(time.RFC3339), title, note, 0, 0)
	err = InFile.Add(event)
	if err != nil {
		t.Error("Can't add event")
	}

	getEvent, err := InFile.Get(0)
	if err != nil {
		t.Error("Get error event")
	}

	if getEvent.StartTime != event.StartTime ||
		getEvent.EndTime != event.EndTime ||
		getEvent.Title != event.Title ||
		getEvent.Note != event.Note ||
		getEvent.Duration != event.Duration ||
		getEvent.TypeDuration != event.TypeDuration {
		t.Error("Event in storage")
	}
}

func TestEditEvent(t *testing.T) {
	InFile,err := NewStorage()
	InFile.Clear()
	dateStart := time.Date(2020, 1, 1, 11, 0, 0, 0, time.UTC)
	dateEnd := time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC)
	for i := 0; i < 10; i++ {
		dateStart.AddDate(0, 0, 1)
		dateEnd.AddDate(0, 0, 1)
		title := "Event " + string(i)
		note := "This envets" + string(i) + " start:" + dateStart.Format(time.RFC3339) + " finish:" + dateEnd.Format(time.RFC3339)
		event, _ := event.CreateEvent(dateStart.Format(time.RFC3339), dateEnd.Format(time.RFC3339), title, note, 0, 0)
		err := InFile.Add(event)
		if err != nil {
			t.Error("Can't add event")
		}
	}
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(10)
	editEvent, _ := InFile.Get(r)

	editEvent.StartTime, _ = time.Parse(time.RFC3339, "2020-01-02T16:15:00Z")
	editEvent.EndTime, _ = time.Parse(time.RFC3339, "2020-01-02T17:30:00Z")
	editEvent.Title = "testTitle"
	editEvent.Note = "testNote"

	err = InFile.Edit(editEvent)
	if err != nil {
		t.Error("Got not expected error on edit")
	}
	InFile.SaveEvents()
	InFile.Clear()
	InFile,err = NewStorage()
	eventAfterEdit, _ := InFile.Get(r)
	if eventAfterEdit.Id != editEvent.Id {
		t.Error("Edit Event not id Event after edit")
	}
}

func TestAddDurrationEventSuccess(t *testing.T) {
	InFile,err := NewStorage()
	InFile.Clear()
	dateStart := time.Date(2020, 1, 1, 11, 0, 0, 0, time.UTC)
	dateEnd := time.Date(2020, 1, 2, 11, 0, 0, 0, time.UTC)
	event, _ := event.CreateEvent(dateStart.Format(time.RFC3339), "", "Event 1", "Start event", 1, event.EnumTypeDuration.Day)
	err = InFile.Add(event)
	if err != nil {
		t.Error("Can't add event")
	}
	events, err := InFile.Get(0)
	if err != nil || events.EndTime != dateEnd {
		t.Error("Error calc dateEnd")
	}
}
