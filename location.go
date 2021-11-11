package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// All I care is the name of the location
type LocationResults struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type LocationsObj struct {
	locations []LocationResults
}

func GetLocationNames() LocationsObj {
	var locResults []LocationResults
	locationsNumber := getInfo(Location).Count
	locationsRange := makeRange(1, locationsNumber)
	locationWithIdsURL := fmt.Sprintf("%s%s", Location, sliceToString(locationsRange))

	locationData, _ := getReq(locationWithIdsURL)
	err := json.Unmarshal(locationData, &locResults)
	if err != nil {
		log.Fatal(err.Error())
	}
	return LocationsObj{locations: locResults}
}

// this approach has a better performance
func (loc *LocationsObj) CountChar(char string) int {
	var count int
	for _, v := range loc.locations {
		count += strings.Count(v.Name, char)
	}
	return count
}
