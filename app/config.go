package app

import (
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	API_PORT string
	REDIS_PORT string
}

func GetConfig() Configuration {
	config := Configuration{}
	err := gonfig.GetConf("config.json", &config)
	if err != nil {
		panic(err)
	}
	return config
}