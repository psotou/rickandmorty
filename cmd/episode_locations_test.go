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
	tests := struct {
		name string
		want []EpisodeLocations
	}{
		name: "",
		want: episodeLocationsGenerator(),
	}

	rangeEpiIds := []string{"1", "2", "3", "4", "5"}

	t.Run(tests.name, func(t *testing.T) {
		if got := episodeLocationsResult(rangeEpiIds).Results; !reflect.DeepEqual(got, tests.want) {
			t.Errorf("episodeLocationsResult() = %v, want %v", got, tests.want)
		}
	})
}

func episodeLocationsGenerator() []EpisodeLocations {
	epiRes := EpiLocations{}
	jsonFile, _ := os.Open("fixtures/episode_location.json")
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	err := json.Unmarshal(byteValue, &epiRes)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return epiRes.Results
}
