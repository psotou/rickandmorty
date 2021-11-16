package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"reflect"
	"testing"
)

func Test_getEpisodes(t *testing.T) {
	tests := []struct {
		name string
		want iEpisodes
	}{
		{
			name: "returns the episode object",
			want: episodeInterfaceGenerator(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getEpisodes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getEpisodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEpisodeObj_countChar(t *testing.T) {
	type fields struct {
		episodes []EpisodeResults
	}
	type args struct {
		char string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "counts the ocurrence of a certain character in the field name",
			fields: fields{episodeObjGenerator().episodes},
			args:   args{"i"},
			want:   95,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &EpisodeObj{
				episodes: tt.fields.episodes,
			}
			if got := c.countChar(tt.args.char); got != tt.want {
				t.Errorf("EpisodeObj.countChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEpisodeObj_characterIds(t *testing.T) {
	type fields struct {
		episodes []EpisodeResults
	}
	tests := []struct {
		name   string
		fields fields
		want   []EpisodeWithCharIds
	}{
		{
			name:   "",
			fields: fields{episodeObjGenerator().episodes},
			want:   episodeCharIdsGenerator(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EpisodeObj{
				episodes: tt.fields.episodes,
			}
			if got := e.characterIds(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EpisodeObj.characterIds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func episodeInterfaceGenerator() iEpisodes {
	var epiRes []EpisodeResults
	jsonFile, _ := os.Open("fixtures/episode.json")
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	err := json.Unmarshal(byteValue, &epiRes)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return &EpisodeObj{episodes: epiRes}
}

func episodeObjGenerator() EpisodeObj {
	var epiRes []EpisodeResults
	jsonFile, _ := os.Open("fixtures/episode.json")
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	err := json.Unmarshal(byteValue, &epiRes)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return EpisodeObj{episodes: epiRes}
}

func episodeCharIdsGenerator() []EpisodeWithCharIds {
	var epiRes []EpisodeWithCharIds
	jsonFile, _ := os.Open("fixtures/character_ids.json")
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	err := json.Unmarshal(byteValue, &epiRes)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return epiRes
}
