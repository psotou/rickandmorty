package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"reflect"
	"testing"
)

func Test_episodeLocationsResult(t *testing.T) {
	tests := []struct {
		name string
		// want EpiLocations
		want []EpisodeLocations
	}{
		{
			name: "",
			// want: []EpisodeLocations{},
			want: episodeLocationsGenerator(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := episodeLocationsResult().Results; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("episodeLocationsResult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func episodeLocationsGenerator() []EpisodeLocations {
	var epiRes EpiLocations
	jsonFile, _ := os.Open("fixtures/episode_location.json")
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	err := json.Unmarshal(byteValue, &epiRes)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return epiRes.Results
}
