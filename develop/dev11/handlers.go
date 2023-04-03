package handlers

import (
	"net/http"

	"github.com/afterglowflexin/wildberries_L2/develop/dev11/data"
)

type CalendarHandler struct {
}

func (h CalendarHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

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
