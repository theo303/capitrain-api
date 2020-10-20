package app

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"io/ioutil"
	
	"github.com/zgegonline/capitrain-api/model"
)
	
const URL = "http://ipwhois.app/json/"

func FillLocations(route model.Route) model.Route {
	for i, hop := range route.Hops {
		l := getLocation(hop.IP)
		route.Hops[i].Location = l
	}
	return route
}

func getLocation(ip string) (model.Location) {
	resp, err := http.Get(URL + ip)
	if err != nil {
		log.Fatalln(err)
        return model.Location {Latitude: "Error", Longitude: "Error"}
	}

	defer resp.Body.Close()
    bodyBytes, _ := ioutil.ReadAll(resp.Body)

    // Convert response body to location struct
    var location model.Location
	json.Unmarshal(bodyBytes, &location)

	fmt.Println(location)
	
	return location
}
