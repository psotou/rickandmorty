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

// makerange produces a slice of strings from a certain number range
func makerange(min, max int) []string {
	strSlice := make([]string, max-min+1)
	for idx := range strSlice {
		strSlice[idx] = strconv.Itoa(min + idx)
	}
	return strSlice
}

// takes a slice of string and turns its content into a string separated by commas
func sliceToString(slc []string) string { return strings.Join(slc, ",") }
