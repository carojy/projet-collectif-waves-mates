package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	// "encoding/json"
	// "io/ioutil"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Salut la compagnie !")
	log.Println("homeLink est appelé")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/events", getEvents)
	router.HandleFunc("/simplespots", getSimpleSpotsList).Methods("GET")
	router.HandleFunc("/events/{id}", getOneEvent).Methods("GET")
	router.HandleFunc("/simplespots/{id}", getSimpleSpot).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

