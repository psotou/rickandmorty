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
	episodeLocationsSlc := []EpisodeLocations{}
	episodeLocations := EpisodeLocations{}
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

// episodeLocationsResult() only takes the range of ids of episodes and
// runs those against the whole universe of characters since there may be
// a lot of characters per episode
func episodeLocationsResult(rangeEpiIds []string) EpiLocations {
	start := time.Now()
	rangeCharIds := makeRange(1, getInfo(Character).Count)
	charIdAndLoc := getCharacters(rangeCharIds).locationName()

	epiAndCharIds := getEpisodes(rangeEpiIds).characterIds()
	epiLocations := episodeLocations(charIdAndLoc, epiAndCharIds)

	var intime bool
	elapsed := time.Since(start)
	if elapsed < time.Duration(3e9) {
		intime = true
	}

	return EpiLocations{
		ExerciseName: "Episode locations",
		Time:         fmt.Sprint(elapsed),
		InTime:       intime,
		Results:      epiLocations,
	}
}
