package storage

import (
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw15/internal/calendar/event"
)

type Interface interface {
	Add(e event.Event) error
	Del(id int) error
	Clear() error
	Get(id int) (event.Event, error)
	GetAll() ([]event.Event, error)
	Edit(event.Event) error
	IsBusy(event.Event) (bool, error)
	CountRecord() int
}
