package model

type Route struct {
	Hops []Hop `json:"hops"`
}

type Hop struct {
	Number 	 string	`json:"number"`
	Url      string `json:"url"`
	IP       string `json:"IP"`
	Location string `json:"location"`
}
