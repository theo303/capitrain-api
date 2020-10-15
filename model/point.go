package model

type Hops struct {
	Hops []Hop `json:"hops"`
}

type Hop struct {
	Url      string `json:"url"`
	IP       string `json:"IP"`
	Location string `json:"location"`
}
