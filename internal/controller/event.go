package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/nilsyadv/baston-eventos/internal/model"
	"github.com/nilsyadv/baston-eventos/internal/service"
	"github.com/nilsyadv/baston-eventos/internal/util"
	"github.com/nilsyadv/baston-eventos/internal/web"
)

type EventController struct{}

func NewEventController() *EventController {
	return &EventController{}
}

func (eventcont *EventController) AddEvent(w http.ResponseWriter, r *http.Request) {
	var event model.Event
	log.Println("Add Event Service Called.....")
	err := web.RequestParse(r, &event)
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}
	_, event.ID = util.CreateID()
	err = service.AddEvent(&event)
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}
	web.RespondJSON(w, http.StatusOK, map[string]interface{}{"message": "New Event Added Successfully"})
}

func (eventcont *EventController) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	var event model.Event
	err := web.RequestParse(r, &event)
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}

	params := mux.Vars(r)
	err = util.ValidateIDFormat(params["eventid"])
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}

	event.ID, err = util.ParseID(params["eventid"])
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}

	err = service.UpdateEvent(&event)
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}
	web.RespondJSON(w, http.StatusOK, map[string]interface{}{"message": "Event Updated Successfully"})
}

func (eventcont *EventController) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	var event model.Event
	params := mux.Vars(r)
	err := util.ValidateIDFormat(params["eventid"])
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}

	event.ID, err = util.ParseID(params["eventid"])
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}
	err = service.DeleteEvent(&event)
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}
	web.RespondJSON(w, http.StatusOK, map[string]interface{}{"message": "Event Deleted Successfully"})
}

func (eventcont *EventController) GetEvent(w http.ResponseWriter, r *http.Request) {
	var event model.Event
	params := mux.Vars(r)
	eventid, err := util.ParseID(params["eventid"])
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}

	err = service.GetEvent(&event, eventid)
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}
	web.RespondJSON(w, http.StatusOK, &event)
}

func (eventcont *EventController) GetEvents(w http.ResponseWriter, r *http.Request) {
	var events []model.Event
	err := service.GetEvents(&events)
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}
	web.RespondJSON(w, http.StatusOK, &events)
}
