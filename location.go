package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// All I care is the name of the location
type locationResults struct {
	Name string `json:"name"`
}

type locationsObj struct {
	locations []locationResults
}

func getLocationNames() locationsObj {
	var locResults []locationResults
	locationsNumber := getInfo(Location).Count
	locationsRange := makerange(1, locationsNumber)
	locationWithIdsURL := fmt.Sprintf("%s%s", Location, sliceToString(locationsRange))

	locationData, _ := getReq(locationWithIdsURL)
	err := json.Unmarshal(locationData, &locResults)
	if err != nil {
		log.Fatal(err.Error())
	}
	return locationsObj{locations: locResults}
}

// this approach has a better performance
func (loc *locationsObj) countChar(char string) int {
	var count int
	for _, v := range loc.locations {
		count += strings.Count(v.Name, char)
	}
	return count
}
