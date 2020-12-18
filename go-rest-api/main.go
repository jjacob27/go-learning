package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type event struct {
	ID      string `json:"ID"`
	Message string `json:"Message"`
}

type allEvents []event

var currentEvents = allEvents{
	{
		ID:      "1",
		Message: "first message",
	},
}

func createEvent(rw http.ResponseWriter, req *http.Request) {
	var newEvent event
	reqBody, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(reqBody, &newEvent)

	for _, singleEvent := range currentEvents {
		if singleEvent.ID == newEvent.ID {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	currentEvents = append(currentEvents, newEvent)
	rw.WriteHeader(http.StatusCreated)
}

func updateEvent(rw http.ResponseWriter, req *http.Request) {
	var updateEvent event
	reqBody, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(reqBody, &updateEvent)

	for i, singleEvent := range currentEvents {
		if singleEvent.ID == updateEvent.ID {
			singleEvent.Message = updateEvent.Message
			currentEvents = append(currentEvents[:i], singleEvent)
			rw.WriteHeader(http.StatusOK)
			return
		}
	}
	rw.WriteHeader(http.StatusBadRequest)

}

func delete(rw http.ResponseWriter, req *http.Request) {
	eventID := mux.Vars(req)["id"]

	for i, singleEvent := range currentEvents {
		if singleEvent.ID == eventID {
			currentEvents = append(currentEvents[:i], currentEvents[i+1:]...)
			rw.WriteHeader(http.StatusOK)
			return
		}
	}

	rw.WriteHeader(http.StatusBadRequest)
}

func getOne(rw http.ResponseWriter, req *http.Request) {
	eventID := mux.Vars(req)["id"]

	for _, singleEvent := range currentEvents {
		if singleEvent.ID == eventID {
			json.NewEncoder(rw).Encode(singleEvent)
		}
	}
}

func getAll(rw http.ResponseWriter, req *http.Request) {
	json.NewEncoder(rw).Encode(currentEvents)
}

func home(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "Welcome to my first rest in Go")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	router.HandleFunc("/event", createEvent).Methods("POST")
	router.HandleFunc("/event", updateEvent).Methods("PATCH")
	router.HandleFunc("/events", getAll).Methods("GET")
	router.HandleFunc("/events/{id}", getOne).Methods("GET")
	router.HandleFunc("/events/{id}", delete).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
