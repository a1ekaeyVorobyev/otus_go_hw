package calendar

import (
	"math/rand"
	"testing"
	"time"

	"github.com/a1ekaeyVorobyev/otus_go_hw/hw22/internal/calendar/event"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw22/internal/storage"
)

func TestNewCalendarHaveNoEvents(t *testing.T) {
	InFile := storage.InFile{}
	InFile.Init()
	InFile.Clear()
	calendar := Calendar{Storage: &InFile}

	events, err := calendar.GetAllEvents()
	if err != ErrNoEventsInStorage || len(events) != 0 {
		t.Error("In new storage exist events")
	}
}

func TestAddEventSuccess(t *testing.T) {
	InFile := storage.InFile{}
	InFile.Init()
	InFile.Clear()
	calendar := Calendar{Storage: &InFile}

	dateStart := time.Date(2020, 1, 1, 11, 0, 0, 0, time.UTC)
	dateEnd := time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC)
	for i := 0; i < 10; i++ {
		dateStart.AddDate(0, 0, 1)
		dateEnd.AddDate(0, 0, 1)
		title := "Event " + string(i)
		note := "This envets" + string(i) + " start:" + dateStart.Format(time.RFC3339) + " finish:" + dateEnd.Format(time.RFC3339)
		event, _ := event.CreateEvent(dateStart.Format(time.RFC3339), dateEnd.Format(time.RFC3339), title, note, 0, 0)
		err := calendar.AddEvent(event)
		if err != nil {
			t.Error("Can't add event")
		}
	}
	events, err := calendar.GetAllEvents()
	if err != nil || len(events) != 10 {
		t.Error("In storage have not 10 event")
	}
}

func TestDeleteEventSuccess(t *testing.T) {
	InFile := storage.InFile{}
	InFile.Init()
	InFile.Clear()
	calendar := Calendar{Storage: &InFile}

	dateStart := time.Date(2020, 1, 1, 11, 0, 0, 0, time.UTC)
	dateEnd := time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC)
	for i := 0; i < 10; i++ {
		dateStart.AddDate(0, 0, 1)
		dateEnd.AddDate(0, 0, 1)
		title := "Event " + string(i)
		note := "This envets" + string(i) + " start:" + dateStart.Format(time.RFC3339) + " finish:" + dateEnd.Format(time.RFC3339)
		event, _ := event.CreateEvent(dateStart.Format(time.RFC3339), dateEnd.Format(time.RFC3339), title, note, 0, 0)
		err := calendar.AddEvent(event)
		if err != nil {
			t.Error("Can't add event ")
		}
	}
	events, _ := calendar.GetAllEvents()
	for _, v := range events {
		err := calendar.DeleteEvent(v.Id)
		if err != nil {
			t.Error("Can't delete event")
		}
	}

	if calendar.CountRecord() != 0 {
		t.Error("In storage have events")
	}
}

func TestAddDateIntervalBusy(t *testing.T) {
	var err error
	inMemory := storage.InFile{}
	inMemory.Init()
	inMemory.Clear()
	calendar := Calendar{Storage: &inMemory}
	dateStart := time.Date(2020, 2, 1, 11, 0, 0, 0, time.UTC)
	//dateEnd := time.Date(2020, 1, 2, 11, 0, 0, 0, time.UTC)
	event1, _ := event.CreateEvent(dateStart.Format(time.RFC3339), "", "Event 1", "Start event", 1, event.EnumTypeDuration.Day)

	event2, _ := event.CreateEvent(dateStart.Format(time.RFC3339), "", "Event 1", "Start event", 1, event.EnumTypeDuration.Hour)
	event3, _ := event.CreateEvent("2020-01-02T15:00:00Z", "", "Event 1", "Start event", 1, event.EnumTypeDuration.Month)

	err = calendar.AddEvent(event1)
	if err != nil {
		t.Error("Error on add events")
	}
	err = calendar.AddEvent(event2)
	if err != ErrBusy {
		t.Error("Add not return error for busy interval")
	}

	err = calendar.AddEvent(event3)

	if err != ErrBusy {
		t.Error("Add not return error for busy interval")
	}

}

func TestGetEvent(t *testing.T) {
	var err error
	InFile := storage.InFile{}
	InFile.Init()
	InFile.Clear()
	calendar := Calendar{Storage: &InFile}
	dateStart := time.Date(2020, 1, 1, 11, 0, 0, 0, time.UTC)
	dateEnd := time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC)
	title := "Event 1"
	note := "This envets1 start:" + dateStart.Format(time.RFC3339) + " finish:" + dateEnd.Format(time.RFC3339)
	event, _ := event.CreateEvent(dateStart.Format(time.RFC3339), dateEnd.Format(time.RFC3339), title, note, 0, 0)
	_ = calendar.AddEvent(event)

	getEvent, err := calendar.GetEvent(0)
	if err != nil {
		t.Error("Get error, for exist event")
	}

	if getEvent.StartTime != event.StartTime ||
		getEvent.EndTime != event.EndTime ||
		getEvent.Title != event.Title ||
		getEvent.Note != event.Note ||
		getEvent.Duration != event.Duration ||
		getEvent.TypeDuration != event.TypeDuration {
		t.Error("Event in storage not ident")
	}
}

func TestEditEvent(t *testing.T) {
	InFile := storage.InFile{}
	InFile.Init()
	InFile.Clear()
	calendar := Calendar{Storage: &InFile}

	dateStart := time.Date(2020, 1, 1, 11, 0, 0, 0, time.UTC)
	dateEnd := time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC)
	for i := 0; i < 10; i++ {
		dateStart.AddDate(0, 0, 1)
		dateEnd.AddDate(0, 0, 1)
		title := "Event " + string(i)
		note := "This envets" + string(i) + " start:" + dateStart.Format(time.RFC3339) + " finish:" + dateEnd.Format(time.RFC3339)
		event, _ := event.CreateEvent(dateStart.Format(time.RFC3339), dateEnd.Format(time.RFC3339), title, note, 0, 0)
		_ = calendar.AddEvent(event)
	}
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(10)
	editEvent, _ := InFile.Get(r)

	editEvent.StartTime, _ = time.Parse(time.RFC3339, "2006-01-02T15:10:00Z")
	editEvent.EndTime, _ = time.Parse(time.RFC3339, "2006-01-02T15:20:00Z")
	editEvent.Title = "Title1"
	editEvent.Note = "Note1"

	err := calendar.EditEvent(editEvent)
	if err != nil {
		t.Error(" expected error after edit")
	}

	eventFromStorageAfterEdit, _ := calendar.GetEvent(r)
	if eventFromStorageAfterEdit != editEvent {
		t.Error("Edit Event not id Event  after edit")
	}
}
