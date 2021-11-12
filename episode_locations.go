package main

import (
	"fmt"
	"time"
)

type EpiLocations struct {
	ExerciseName string             `json:"exercise_name"`
	Time         string             `json:"time"`
	InTime       bool               `json:"in_time"`
	Results      []EpisodeLocations `json:"results"`
}

func episodeLocationsResult() EpiLocations {
	start := time.Now()
	charIdAndLoc := getCharacters().charIdWithLocationName()
	epiAndCharIds := getEpisodes().characterIdsPerEpisode()
	epiLocations := episodeLocations(charIdAndLoc, epiAndCharIds)

	var intime bool
	elapsed := time.Since(start)
	if elapsed < time.Duration(3*1e6) {
		intime = true
	}

	return EpiLocations{
		ExerciseName: "Episode locations",
		Time:         fmt.Sprint(elapsed),
		InTime:       intime,
		Results:      epiLocations,
	}
}
