package pkg

import (
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/calendar/event"
	"time"
)

type Storage interface {
	Add(e event.Event) error
	Delete(id int) error
	Clear() error
	Get(id int) (event.Event, error)
	GetAll() ([]event.Event, error)
	Edit(event.Event) error
	IsBusy(event.Event) (bool, error)
	CountRecord() int
}

type Scheduler interface {
	GetEventSending(time.Time) ([]event.Event, error)
	MarkEventSentToQueue(int) error
	MarkEventSentToSubScribe(int) error
}
