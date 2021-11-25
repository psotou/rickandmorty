package main

import (
	"fmt"
	"time"
)

type CharCounter struct {
	ExerciseName string               `json:"exercise_name"`
	Time         string               `json:"time"`
	InTime       bool                 `json:"in_time"`
	Results      []CharCounterResults `json:"results"`
}

type CharCounterResults struct {
	Char     string `json:"char"`
	Count    int    `json:"count"`
	Resource string `json:"resource"`
}

type ResourceRange struct {
	Resource string
	Range    []string
}

func charCounter(resource ResourceRange) CharCounterResults {
	count := 0
	char := ""
	switch resource.Resource {
	case "location":
		char = "l"
		count = getLocations(resource.Range).countChar(char)
	case "episode":
		char = "e"
		count = getEpisodes(resource.Range).countChar(char)
	case "character":
		char = "c"
		count = getCharacters(resource.Range).countChar(char)
	}

	return CharCounterResults{
		Char:     char,
		Count:    count,
		Resource: resource.Resource,
	}
}

func charCounterResult(resourceIdsRangeSlc []ResourceRange) CharCounter {
	start := time.Now()
	res := []CharCounterResults{}

	for _, resource := range resourceIdsRangeSlc {
		res = append(res, charCounter(resource))
	}
	elapsed := time.Since(start)
	var intime bool
	if elapsed < time.Duration(3e9) {
		intime = true
	}

	return CharCounter{
		ExerciseName: "Char counter",
		Time:         fmt.Sprint(elapsed),
		InTime:       intime,
		Results:      res,
	}
}
