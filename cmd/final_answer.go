package main

import (
	"encoding/json"
	"log"
	"os"
)

type iFinalRes interface {
	writeJSON()
}

type FinalResult struct {
	FinalRes string
}

func finalResult() iFinalRes {
	//
	// char counter part
	//
	resourceRangeOfIds := []ResourceRange{
		{
			Resource: "location",
			Range:    makeRange(1, getInfo(Location).Count),
		},
		{
			Resource: "episode",
			Range:    makeRange(1, getInfo(Episode).Count),
		},
		{
			Resource: "character",
			Range:    makeRange(1, getInfo(Character).Count),
		},
	}
	charCount := charCounterResult(resourceRangeOfIds)

	//
	// location (origin) name part
	//
	rangeEpiIds := makeRange(1, getInfo(Episode).Count)
	locNames := episodeLocationsResult(rangeEpiIds)
	res := []interface{}{charCount, locNames}
	dataBytes, err := json.Marshal(res)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return &FinalResult{FinalRes: string(dataBytes)}
}

func (f *FinalResult) writeJSON() {
	// return os.WriteFile("result", data, 0644)
	file, _ := os.Create("result.json")
	defer file.Close()

	_, err := file.WriteString(f.FinalRes)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
