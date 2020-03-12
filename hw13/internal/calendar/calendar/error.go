package calendar

type BusinessError string

func (e BusinessError) Error() string {
	return string(e)
}

const (
	ErrBusy              = BusinessError("Date and Time  already busy by another event")
	ErrNoEventsInStorage = BusinessError("No events in storage")
)
