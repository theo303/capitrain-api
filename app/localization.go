package app

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"io/ioutil"
	
	"github.com/zgegonline/capitrain-api/model"
)
	
const URL_IPWHOIS = "http://ipwhois.app/json/"
const URL_IPAPI = "http://ip-api.com/json/"

func FillLocations(route model.Route) model.Route {
	for i, hop := range route.Hops {
		if hop.Number == "1" {
			hop.IP = "90.59.71.87"
		}
		l := getLocation(hop.IP)
		route.Hops[i].Location = l
	}
	return route
}

func getLocation(ip string) (model.Location) {
	fmt.Println(URL_IPAPI + ip)
	
	resp, err := http.Get(URL_IPAPI + ip)
	
	if err != nil {
		log.Fatalln(err)
        return model.Location {Latitude: 0.0, Longitude: 0.0}
	}

	defer resp.Body.Close()
    bodyBytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bodyBytes))
    // Convert response body to location struct
	
	var location model.Location					
	
	json.Unmarshal(bodyBytes, &location)

	fmt.Println(location)
	
	return location
}
