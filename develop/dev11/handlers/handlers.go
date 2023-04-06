package handlers

import (
	"net/http"

	"dev11/data"
)

type CalendarHandler struct {
}

func (h CalendarHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func (h CalendarHandler) ListForDay(w http.ResponseWriter, r *http.Request) {

	event, err := data.GetEvent()

	if err != nil {

	}

	data.ToJSON(event, w)
}

func (h CalendarHandler) ListForWeek(w http.ResponseWriter, r *http.Request) {

	event, err := data.GetEvent()

	if err != nil {

	}

	data.ToJSON(event, w)
}

func (h CalendarHandler) ListForMonth(w http.ResponseWriter, r *http.Request) {

	event, err := data.GetEvent()

	if err != nil {

	}

	data.ToJSON(event, w)
}

func (h CalendarHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	event := data.Event{}

	err := data.FromJSON(event, r.Body)
	if err != nil {

	}

	err = data.CreateEvent(event)
	if err != nil {

	}
}

func (h CalendarHandler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	event := data.Event{}

	err := data.FromJSON(event, r.Body)
	if err != nil {

	}

	err = data.UpdateEvent(event)
	if err != nil {

	}
}
