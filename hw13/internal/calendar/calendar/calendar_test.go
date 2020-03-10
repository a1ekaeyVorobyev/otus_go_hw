package calendar

import (
	"math/rand"
	"testing"
	"time"

	"github.com/a1ekaeyVorobyev/otus_go_hw/hw13/internal/calendar/event"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw13/internal/storage"
)

func TestNewCalendarHaveNoEvents(t *testing.T) {
	InFile := storage.InFile{}
	InFile.Init()
	InFile.Clear()
	calendar := Calendar{Storage:&InFile}

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
			t.Error("Can't add event to storage")
		}
	}
	events, err := calendar.GetAllEvents()
	if err != nil || len(events) != 10 {
		t.Error("In storage not 10 event")
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
			t.Error("Can't add event to storage")
		}
	}
	events, _ := calendar.GetAllEvents()
	for _, v := range events {
		err := calendar.DelEvent(v.Id)
		if err != nil {
			t.Error("Can't del event from storage")
		}
	}

	if calendar.CountRecord() != 0 {
		t.Error("In storage exist events")
	}
}

func TestAddDateIntervalBusy(t *testing.T) {
	var err error
	inMemory := storage.InFile{}
	inMemory.Init()
	calendar := Calendar{Storage: &inMemory}

	event1, _ := event.CreateEvent("2006-01-02T15:00:00Z", "2006-01-02T16:00:00Z", "Event 1", "Some Desc1", 0, 0)
	event2, _ := event.CreateEvent("2006-01-02T16:00:00Z", "2006-01-02T17:00:00Z", "Event 2", "Some Desc2", 0, 0)
	event3, _ := event.CreateEvent("2006-01-02T18:00:00Z", "2006-01-02T19:00:00Z", "Event 3", "Some Desc3", 0, 0)
	err = calendar.AddEvent(event1)
	err = calendar.AddEvent(event2)
	err = calendar.AddEvent(event3)
	if err != nil {
		t.Error("Error on add not intersection events")
	}

	event4, _ := event.CreateEvent("2006-01-02T16:10:00Z", "2006-01-02T16:20:00Z", "Event 4", "Some Desc4")
	err = calendar.AddEvent(event4)
	if err != ErrDateBusy {
		t.Error("Add not return error for busy interval")
	}

	event5, _ := event.CreateEvent("2006-01-02T10:10:00Z", "2006-01-02T22:00:00Z", "Event 5", "Some Desc5")
	err = calendar.AddEvent(event5)
	if err != ErrDateBusy {
		t.Error("Add not return error for busy interval")
	}

	event6, _ := event.CreateEvent("2006-01-02T17:10:00Z", "2006-01-02T18:10:00Z", "Event 6", "Some Desc6")
	err = calendar.AddEvent(event6)
	if err != ErrDateBusy {
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
		t.Error("Got not expected error on edit")
	}

	eventFromStorageAfterEdit, _ := calendar.GetEvent(r)
	if eventFromStorageAfterEdit != editEvent {
		t.Error("Edit Event not ident Event in storage after edit")
	}
}
