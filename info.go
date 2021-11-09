package main

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
)

type Info struct {
	Count int    `json:"count"`
	Pages int    `json:"pages"`
	Next  string `json:"next"`
	Prev  string `json:"prev"`
}

type Inf struct {
	Info Info `json:"info"`
}

func getInfo(endpoint string) Info {
	var info Inf
	infoData, _ := getReq(endpoint)
	err := json.Unmarshal(infoData, &info)
	if err != nil {
		log.Fatal(err.Error())
	}
	return info.Info
}

// makerange produces a slice of strings that ranges from 1 to
// the number of IDs returned by the info object of a certain endpoint
func makerange(min, max int) []string {
	strSlice := make([]string, max-min+1)
	for idx := range strSlice {
		strSlice[idx] = strconv.Itoa(min + idx)
	}
	return strSlice
}

// sliceToStringes poduces the string range allowed by the endpoint for querying
// multiple objects (used to return all the objects from one http call)
func sliceToString(slc []string) string { return strings.Join(slc, ",") }
