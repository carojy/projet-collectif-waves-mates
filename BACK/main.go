package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello surfers !")
	log.Println("homeLink est appelé")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/events", getAllEvents)
	router.HandleFunc("/simplespots", getSimpleSpotsList).Methods("GET")
	router.HandleFunc("/events/{id}", getOneEvent).Methods("GET")
	router.HandleFunc("/simplespots/{id}", getSimpleSpot).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

