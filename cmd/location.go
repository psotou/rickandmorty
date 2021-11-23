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
func getLocations(locationsRange []string) iLocation {
	locResults := []LocationResults{}
	// locationEndpointMultipleIds returns the ids in range to fetch multiple episodes
	// See https://rickandmortyapi.com/documentation/#get-multiple-locations
	locationEndpointMultipleIds := fmt.Sprintf("%s%s", Location, sliceToString(locationsRange))

	locationData, _ := getReq(locationEndpointMultipleIds)
	err := json.Unmarshal(locationData, &locResults)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &LocationsObj{locations: locResults}
}

// countChar method implemented on the LocationsObj struct
// Counts the ocurrence of a certain character in the LocationResults.Name field
func (loc *LocationsObj) countChar(char string) int {
	count := 0
	for _, location := range loc.locations {
		count += strings.Count(location.Name, char)
	}
	return count
}
