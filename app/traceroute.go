package app

import (
	"bytes"
	"os/exec"
	"strings"
	"fmt"

	"github.com/zgegonline/capitrain-api/model"
)

func Traceroute(address string) (model.Route, error) {
	cmd := exec.Command("traceroute", address)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		return model.Route{}, err
	}
	return parse(out.String()), err
}

func parse(out string) model.Route {
	var route model.Route
	lines := strings.Split(out, "\n")
	fmt.Println("number of lines", len(lines))
	for i, line := range lines {					// loop over lines
		if i != 0 && i != (len(lines) - 1) {		// first and last lines contains nothing usefull
			field := strings.Split(strings.TrimSpace(line), " ")		// trim then split line 
			
			hop := model.Hop {
				Number: field[0],
				Url:	field[2],					
				IP:		strings.Trim(field[3], "()"),					// remove parenthesis surrounding IP 
			}
			route.Hops = append(route.Hops, hop)
		}
	}
	return route
}