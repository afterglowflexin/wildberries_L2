package data

import "time"

type Event struct {
	date    time.Time `json:"date"`
	user_ID int       `json:"user_id"`
	event   string
}

type Calendar struct {
	events []Event
}

var calendar = Calendar{
	[]Event{},
}

func GetEvent(date time.Time) (Event, error) {
	return Event{}, nil
}

func CreateEvent(event Event) error {
	calendar.events = append(calendar.events, event)
	return nil
}

func UpdateEvent(event Event) error {
	var id int

	calendar.events[id] = event

	return nil
}

func DeleteEvent(event Event) error {
	//
	return nil
}
