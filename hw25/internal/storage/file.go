package storage

import (
	"encoding/gob"
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/calendar/event"
)

const fileName = "Events.dat"

type InFile struct {
	sync.Mutex
	events map[int]event.Event
}

func NewStorage()(s *InFile, err error){
	s.events = make(map[int]event.Event)
	s.loadEvents()
	return s,nil
}


func fileExists(filename string) bool {

	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func (s *InFile) SaveEvents() error {
	s.Lock()
	defer s.Unlock()
	if fileExists(fileName) {
		err := os.Remove(fileName)
		if err != nil {
			return fmt.Errorf("cant delete file")
		}
	}
	f, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("cant open file")
	}
	defer f.Close()

	enc := gob.NewEncoder(f)
	if err := enc.Encode(s.events); err != nil {
		return fmt.Errorf("cant encode")
	}
	return nil
}

func (s *InFile) loadEvents() error {
	s.Lock()
	defer s.Unlock()
	if !fileExists(fileName) {
		return nil
	}
	f, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("cant open file")
	}
	defer f.Close()

	enc := gob.NewDecoder(f)
	if err := enc.Decode(&s.events); err != nil {
		return fmt.Errorf("cant decode")
	}
	return nil
}

func (s *InFile) Add(e event.Event) error {
	s.Lock()
	defer s.Unlock()
	e.Id = len(s.events)
	s.events[len(s.events)] = e
	return nil
}

func (s *InFile) Delete(id int) error {
	s.Lock()
	defer s.Unlock()
	delete(s.events, id)
	return nil
}

func (s *InFile) Clear() error {
	s.Lock()
	defer s.Unlock()
	s.events = make(map[int]event.Event)
	return nil
}

func (s *InFile) Get(id int) (event.Event, error) {
	event, exist := s.events[id]
	if !exist {
		return event, errors.New(fmt.Sprintf("Event with id: %d not found", id))
	}
	return event, nil
}

func (s *InFile) GetAll() ([]event.Event, error) {
	if s.CountRecord() == 0 {
		return make([]event.Event, 0), nil
	}
	events := make([]event.Event, 0, len(s.events))
	for _, e := range s.events {
		events = append(events, e)
	}
	return events, nil
}

func (s *InFile) Edit(e event.Event) error {
	s.Lock()
	defer s.Unlock()
	_, exist := s.events[e.Id]
	if !exist {
		return errors.New(fmt.Sprintf("Event with id: %d not found", e.Id))
	}
	s.events[e.Id] = e
	return nil
}

func (s *InFile) IsBusy(newEvent event.Event) (bool, error) {
	for id, Event := range s.events {
		if newEvent.Id == id && newEvent.Id != 0 {
			continue
		}
		if newEvent.StartTime.Before(Event.StartTime) && newEvent.EndTime.After(Event.EndTime) {
			return true, nil
		}
		if newEvent.EndTime.After(Event.StartTime) && newEvent.EndTime.Before(Event.EndTime) {
			return true, nil
		}
	}
	return false, nil
}

func (s *InFile) CountRecord() int {
	return len(s.events)
}
