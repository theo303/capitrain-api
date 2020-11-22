package app

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"io/ioutil"
	
	"github.com/zgegonline/capitrain-api/model"
)
	
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

func getLocation(ip string) model.Location {
	if (Conf.REDIS_PORT == "-1") {		//DB disabled
		return getLocationFromAPI(ip)
	}


	val, _ := Get(ip + REDIS_LOCATION_SUFFIXE);
	if val == "" {
		fmt.Println("nil db")
		return getLocationFromAPI(ip)
	} else {
		fmt.Println(val)
		var location model.Location
		json.Unmarshal([]byte(val), &location)
		fmt.Println(location)
		return location
	}

}

func getLocationFromAPI(ip string) (model.Location) {
	resp, err := http.Get(URL_IPAPI + ip)
	
	if err != nil {
		log.Fatalln(err)
        return model.Location {Latitude: 0.0, Longitude: 0.0}
	}

	defer resp.Body.Close()
    bodyBytes, _ := ioutil.ReadAll(resp.Body)
	
	var location model.Location					
	
	json.Unmarshal(bodyBytes, &location)

	storeLocation(ip, location)
	
	return location
}

func storeLocation(ip string, location model.Location) {
	if (Conf.REDIS_PORT != "-1") {
		valueString, _ := json.Marshal(location)
		Store(ip + REDIS_LOCATION_SUFFIXE, string(valueString))
	}
}

