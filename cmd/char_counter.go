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

func charCounter(resource string) CharCounterResults {
	var (
		count int
		char  string
	)
	switch resource {
	case "location":
		char = "l"
		count = getLocations().countChar(char)
	case "epsidode":
		char = "e"
		count = getEpisodes().countChar(char)
	case "character":
		char = "c"
		count = getCharacters().countChar(char)
	}

	res := CharCounterResults{
		Char:     char,
		Count:    count,
		Resource: resource,
	}

	return res
}

func charCounterResult() CharCounter {
	start := time.Now()
	charCounterResources := []string{"location", "epsidode", "character"}
	var res []CharCounterResults
	for _, resource := range charCounterResources {
		res = append(res, charCounter(resource))
	}
	elapsed := time.Since(start)
	var intime bool
	if elapsed < time.Duration(3*1e9) {
		intime = true
	}

	return CharCounter{
		ExerciseName: "Char counter",
		Time:         fmt.Sprint(elapsed),
		InTime:       intime,
		Results:      res,
	}
}
