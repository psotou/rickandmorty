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

var (
	count int
	char  string
)

func charCounter(resource string) CharCounterResults {
	switch resource {
	case "location":
		char = "l"
		rng := makeRange(1, getInfo(Location).Count)
		count = getLocations(rng).countChar(char)
	case "epsidode":
		char = "e"
		// rng := makeRange(1, getInfo(Episode).Count)
		count = getEpisodes().countChar(char)
	case "character":
		char = "c"
		rng := makeRange(1, getInfo(Character).Count)
		count = getCharacters(rng).countChar(char)
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
	res := []CharCounterResults{}
	charCounterResources := []string{"location", "epsidode", "character"}
	for _, resource := range charCounterResources {
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
