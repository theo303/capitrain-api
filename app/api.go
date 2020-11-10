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

func Start() {
	fmt.Println("Starting API...")
	log.Fatal(handleRequests())
}

func handleRequests() error {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/traceroute", traceRoute).Methods("POST")

	fmt.Println("Starting router...")
	return http.ListenAndServe(":8080", router)
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

	fmt.Println("Address = " + newRequest.Address)

	out, err := Traceroute(newRequest.Address)
	if err != nil {
		fmt.Fprintf(w, "Error during traceroute"+err.Error())
	} 

	out = FillLocations(out)

	valueString, _ := json.Marshal(out)
	Store(newRequest.Address, string(valueString))

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(out)
}
