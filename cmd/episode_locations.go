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

type EpisodeLocations struct {
	Name      string   `json:"name"`
	Episode   string   `json:"episode"`
	Locations []string `json:"locations"`
}

// episodeLocations takes all the locations (origins) per character and all the character ids per episode
// and returns the locations per episode (unique values)
func episodeLocations(characterIdLocationMap map[string]string, epiCharIds []EpisodeWithCharIds) []EpisodeLocations {
	episodeLocationMap := make(map[string][]string)
	var episodeLocationsSlc []EpisodeLocations
	var episodeLocations EpisodeLocations
	for _, episode := range epiCharIds {
		for _, characterId := range episode.CharacterIds {
			episodeLocationMap[episode.EpisodeCode] = append(episodeLocationMap[episode.EpisodeCode], characterIdLocationMap[characterId])
			episodeLocations = EpisodeLocations{
				Name:    episode.EpisodeName,
				Episode: episode.EpisodeCode,
			}
		}
		episodeLocations.Locations = removeDuplicateStr(episodeLocationMap[episode.EpisodeCode])
		episodeLocationsSlc = append(episodeLocationsSlc, episodeLocations)
	}

	return episodeLocationsSlc
}

func episodeLocationsResult() EpiLocations {
	start := time.Now()
	charIdAndLoc := getCharacters().locationName()
	epiAndCharIds := getEpisodes().characterIds()
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
