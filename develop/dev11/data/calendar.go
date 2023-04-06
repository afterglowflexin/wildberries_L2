package data

import "time"

type Event struct {
	date    time.Time `json:"date"`
	user_ID int       `json:"user_id"`
	event   string
}

type Calendar struct {
	events map[time.Time][]Event
}

var calendar = Calendar{
	map[time.Time][]Event{},
}

func GetEvent(date time.Time) ([]Event, error) {
	return calendar.events[date], nil
}

func CreateEvent(event Event) error {
	calendar.events[event.date] = append(calendar.events[event.date], event)
	return nil
}

func UpdateEvent(event Event) error {

	calendar.events[event.date][0] = event

	return nil
}

func DeleteEvent(event Event) error {
	//
	return nil
}
