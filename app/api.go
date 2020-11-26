package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/theo303/capitrain-api/model"
)

var Conf Configuration
var PUBLIC_IP = "0.0.0.0"

func Start() {
	fmt.Println("Loading configuration")
	Conf = GetConfig()
	getPublicIP()
	fmt.Println("Starting API...")
	log.Fatal(handleRequests())
}

func getPublicIP() {
	url := "https://api.ipify.org?format=text"	// we are using a pulib IP API, we're using ipify here, below are some others
                                              // https://www.ipify.org
                                              // http://myexternalip.com
                                              // http://api.ident.me
                                              // http://whatismyipaddress.com/api
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("My IP is:%s\n", ip)
	PUBLIC_IP = string(ip)
}

func handleRequests() error {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/traceroute", traceRoute).Methods("POST")
	router.HandleFunc("/all-routes", getAllRoutes).Methods("GET")

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
		return
	}

	json.Unmarshal(reqBody, &newRequest)

	route, err := Traceroute(newRequest.Address)
	if err != nil {
		fmt.Fprintf(w, "Error during traceroute"+err.Error())
		return
	} 

	route = FillLocations(route)
	route = model.ClearHopsWithoutLocation(route)
	route.Address = newRequest.Address


	storeRoute(route)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(route)
}

func storeRoute(route model.Route) {
	if (Conf.REDIS_PORT != "-1") {
		valueString, _ := json.Marshal(route)
		Store(route.Address, string(valueString))
		AddToAddressList(route.Address)
	}
}

func getAllRoutes(w http.ResponseWriter, r *http.Request) {
	if (Conf.REDIS_PORT == "-1") {
		fmt.Fprintf(w, "REDIS is disabled (add valid port in config file)")
		return
	}
	addressList, _ := GetAddressList()			//retrieve address list in DB
	
	var routes []model.Route
	
	for _, address := range addressList {		//get route in DB for each address
		val, _ := Get(address)
		var route model.Route
		json.Unmarshal([]byte(val), &route)
		routes = append(routes, route)
	}

	json.NewEncoder(w).Encode(routes)
}
