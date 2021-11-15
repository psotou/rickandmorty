package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// iLocation interface defines the methods associated with the LocationsObj struct
type iLocation interface {
	countChar(string) int
}

// LocationsObj is the struct type that implemts the iLocation interface
// Contains a collection of the object LocationResults
type LocationsObj struct {
	locations []LocationResults
}
type LocationResults struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// getLocations embeds LocationsObj struct and indirectly implements
// the iLocation interface. This approach allows for the use of a syntax
// like getLocations().countChar() declared in one line
func getLocations() iLocation {
	var locResults []LocationResults
	locationsNumber := getInfo(Location).Count
	locationsRange := makeRange(1, locationsNumber)
	locationWithIdsURL := fmt.Sprintf("%s%s", Location, sliceToString(locationsRange))

	locationData, _ := getReq(locationWithIdsURL)
	err := json.Unmarshal(locationData, &locResults)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &LocationsObj{locations: locResults}
}

// countChar method implemented on the LocationsObj struct
// Counts the ocurrence of a certain character in the LocationResults.Name field
func (loc *LocationsObj) countChar(char string) int {
	var count int
	for _, v := range loc.locations {
		count += strings.Count(v.Name, char)
	}
	return count
}
