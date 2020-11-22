package model

type Route struct {
	Address string `json:"address"`
	Hops []Hop `json:"hops"`
}

type Hop struct {
	Number 	 string		`json:"number"`
	Url      string 	`json:"url"`
	IP       string   	`json:"IP"`
	Location Location 	`json:"location"`
}

type Location struct {
	Latitude	float64	`json:"lat"`
	Longitude	float64	`json:"lon"`
}

func ClearHopsWithoutLocation(route Route) Route {
	var new_hops Route
	for _, h := range route.Hops {
		if h.Location.Longitude != 0.0 {
			new_hops.Hops = append(new_hops.Hops, h)
		}
	}
	return new_hops
}
