package main

import "encoding/json"

type EpiLocations struct {
	ExerciseName string `json:"exercise_name"`
	// check the tyope of this thou
	Time    string                `json:"time"`
	InTime  bool                  `json:"in_time"`
	Results []EpiLocationsResults `json:"results"`
}

type EpiLocationsResults struct {
	Name      string   `json:"name"`
	Episode   string   `json:"episode"`
	Locations []string `json:"locations"`
}

func episodeLocations() ([]byte, error) {
	charNames := getCharacterNames()
	epiNames := getEpisodeNames()
	charIdsPerEpi := epiNames.characterIdsPerEpisode()
	locNames := charNames.charIdWithLocationName()
	epiLocations := charIdsPerEpi.locationPerEpisode(locNames)

	return json.Marshal(epiLocations)

}
