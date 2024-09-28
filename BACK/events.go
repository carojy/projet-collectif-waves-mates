package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"projet-collectif-waves-mates/BACK/assets"
)

func getAllEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(assets.Events)
}


func getOneEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for _, singleEvent := range assets.Events.Records {
		if singleEvent.ID == eventID {
			json.NewEncoder(w).Encode(singleEvent)
			return
		}
	}
}
