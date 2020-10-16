package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aeden/traceroute"
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

	fmt.Println("Starting router2...")
	return http.ListenAndServe(":8080", router)
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Zgeg")
}

func traceRoute(w http.ResponseWriter, r *http.Request) {
	var newRequest model.Request

	fmt.Println("alo")

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error reading request")
	}

	json.Unmarshal(reqBody, &newRequest)

	out, err := traceroute.Traceroute(newRequest.Address, new(traceroute.TracerouteOptions))
	if err != nil {
		fmt.Fprintf(w, "Error during traceroute"+err.Error())
	} else {
		if len(out.Hops) == 0 {
			fmt.Fprintf(w, "TestTraceroute failed. Expected at least one hop")
		}
	}
	for _, hop := range out.Hops {
		printHop(hop)
	}
	fmt.Println()

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newRequest)
}

func printHop(hop traceroute.TracerouteHop) {
	fmt.Printf("%-3d %v (%v)  %v\n", hop.TTL, hop.HostOrAddressString(), hop.AddressString(), hop.ElapsedTime)
}
