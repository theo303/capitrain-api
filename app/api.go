package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zgegonline/capitrain-api/model"
)

var Conf Configuration

func Start() {
	fmt.Println("Loading configuration")
	Conf = GetConfig()
	fmt.Println("Starting API...")
	log.Fatal(handleRequests())
}

func handleRequests() error {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/traceroute", traceRoute).Methods("POST")

	fmt.Println("Starting router on port : " + Conf.API_PORT + "...")
	return http.ListenAndServe(":" + Conf.API_PORT, router)
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Home")
}

func traceRoute(w http.ResponseWriter, r *http.Request) {
	var newRequest model.Request

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error reading request")
	}

	json.Unmarshal(reqBody, &newRequest)

	route, err := Traceroute(newRequest.Address)
	if err != nil {
		fmt.Fprintf(w, "Error during traceroute"+err.Error())
	} 

	route = FillLocations(route)

	valueString, _ := json.Marshal(route)
	Store(newRequest.Address, string(valueString))

	route = model.ClearHopsWithoutLocation(route)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(route)
}
