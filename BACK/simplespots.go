package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"projet-collectif-waves-mates/BACK/assets"
)


//SUMMARY SPOTS LIST

type SimpleSpots struct {
	SurfBreak []string `json:"surfBreak"`
	Address string `json:"address"`
}

func getSimpleSpotsList(w http.ResponseWriter, r *http.Request) {
	var simpleSpotsList []SimpleSpots

	for _, singleEvent := range assets.Events.Records {
		spotsListInfos := SimpleSpots{
			SurfBreak: singleEvent.Fields.SurfBreak,
			Address: singleEvent.Fields.Address,
		}
		simpleSpotsList = append(simpleSpotsList, spotsListInfos)
	}
	json.NewEncoder(w).Encode(simpleSpotsList)
}


func getSimpleSpot(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for _, singleEvent := range assets.Events.Records {
		if singleEvent.ID == eventID {
			oneSpotInfos := SimpleSpots{
				SurfBreak: singleEvent.Fields.SurfBreak,
				Address: singleEvent.Fields.Address,
			}
			json.NewEncoder(w).Encode(oneSpotInfos)
			return
		}
	}

}

