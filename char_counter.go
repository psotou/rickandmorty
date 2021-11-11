package main

import "encoding/json"

type CharCounter struct {
	ExerciseName string `json:"exercise_name"`
	// check the tyope of this thou
	Time    string               `json:"time"`
	InTime  bool                 `json:"in_time"`
	Results []CharCounterResults `json:"results"`
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
		locNames := GetLocationNames()
		count = locNames.CountChar(char)
	case "epsidode":
		char = "e"
		epiNames := getEpisodeNames()
		count = epiNames.countChar(char)
	case "character":
		char = "c"
		charNames := getCharacterNames()
		count = charNames.countChar(char)
	}

	res := CharCounterResults{
		Char:     char,
		Count:    count,
		Resource: resource,
	}

	return res
}

func charCounterResult() ([]byte, error) {
	charCounterResources := []string{"location", "epsidode", "character"}
	var res []CharCounterResults
	for _, resource := range charCounterResources {
		cc := charCounter(resource)
		res = append(res, cc)
	}
	return json.Marshal(&res)
}
