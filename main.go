package main

import (
	"WordCount/Utilities"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

// HealthCheck api to test application
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

// createEvent is a post API
// Returns top 10 word count with status, 200
// Returns 400 status for bad invalid requests
func createEvent(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil || len(reqBody) == 0 {
		_, err := fmt.Fprintf(w, "Invalid data Sent!")
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	output := Utilities.Repetition(string(reqBody))
	json.NewEncoder(w).Encode(output)
}

func main() {
	port := ":8080"
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", HealthCheck)
	router.HandleFunc("/event", createEvent).Methods("POST")
	fmt.Println(" Server started and Listening on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
