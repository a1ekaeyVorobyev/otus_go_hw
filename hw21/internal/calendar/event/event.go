package event

import (
	"fmt"
	"time"
)

type Event struct {
	Id          	int 		`yaml:"Id"`
	StartTime   	time.Time 	`yaml:"StartTime"`
	EndTime     	time.Time 	`yaml:"EndTime"`
	Duration		int 		`yaml:"Duration"`
	TypeDuration 	int 		`yaml:"TypeDuration"`
	Title       	string 		`yaml:"Title"`
	Note 			string 		`yaml:"Note"`
}

type AliasDuration = int

type listTypeDuration struct {
	Min AliasDuration
	Hour AliasDuration
	Day AliasDuration
	Month AliasDuration
}

var EnumTypeDuration = &listTypeDuration{
	Min : 1,
	Hour: 2,
	Day: 3,
	Month: 4,
}

func  getEndTime(start time.Time,Duration int,TypeDuration int )(time.Time,error) {
	switch TypeDuration {
	case 1:
		return start.Add(time.Minute * time.Duration(Duration)),nil
	case 2:
		return start.Add(time.Hour * time.Duration(Duration)),nil
	case 3:
		return start.AddDate(0, 0, Duration),nil
	case 4:
		return start.AddDate(0, Duration, 0),nil
	}
	return time.Time{},fmt.Errorf("Не верно указан тип")
}

func CreateEvent(startTime, endTime, title, Note string,Duration int,TypeDuration int) (Event, error) {
	e := Event{}
	sTime, err := time.Parse(time.RFC3339, startTime)
	if err != nil {
		return e, err
	}
	e.StartTime = sTime

	if endTime == "" {
		e.EndTime, err = getEndTime(sTime,Duration,TypeDuration)
		if err != nil {
			return e, err
		}
	}else
	{
		eTime, err := time.Parse(time.RFC3339, endTime)
		if err != nil {
			return e, err
		}
		e.EndTime = eTime
	}

	e.Duration = Duration
	e.TypeDuration = TypeDuration
	e.Title = title
	e.Note = Note

	return e, nil
}
