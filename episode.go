package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// iEpisodes interface defines the methods associated with the EpisodeObj struct
type iEpisodes interface {
	countChar(string) int
	// characterIdsPerEpisode() CharIdsEpisodeObj
	characterIdsPerEpisode() []EpisodeWithCharIds
}

type EpisodeObj struct {
	episodes []EpisodeResults
}

// EpisodeResults cointains the necessary information to map all
// locations used per episode
type EpisodeResults struct {
	Id         int      `json:"id"`
	Name       string   `json:"name"`
	Episode    string   `json:"episode"`
	Characters []string `json:"characters"`
}

// getEpisodes embeds EpisodeObj struct and indirectly implements
// the iEpisodes interface. This approach allows for the use of a syntax
// like getEpisodes().countChar() declared in one line
func getEpisodes() iEpisodes {
	var episodeResults []EpisodeResults
	episodeNumber := getInfo(Episode).Count
	episodeRange := makeRange(1, episodeNumber)
	episodeWithIdsURL := fmt.Sprintf("%s%s", Episode, sliceToString(episodeRange))

	episodeData, _ := getReq(episodeWithIdsURL)
	err := json.Unmarshal(episodeData, &episodeResults)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &EpisodeObj{episodes: episodeResults}
}

// countChar method implemented on the EpisodeObj struct
// Counts the ocurrence of a certain character in the EpisodeResult.Name field
func (c *EpisodeObj) countChar(char string) int {
	var count int
	for _, v := range c.episodes {
		count += strings.Count(v.Name, char)
	}
	return count
}

// EpisodeWithCharIds is similar to EpisodeResults struct
// however, it stores the ids of the characters instead of the endpoint
// of every character
type EpisodeWithCharIds struct {
	EpisodeName  string
	EpisodeCode  string
	CharacterIds []string
}

// characterIdsPerEpisode maps every episode with a slice containing
// the characters ids that appeared in said episode
// func (e *EpisodeObj) characterIdsPerEpisode() CharIdsEpisodeObj {
func (e *EpisodeObj) characterIdsPerEpisode() []EpisodeWithCharIds {
	charIdsSlc := make(map[string][]string)
	var charIds []EpisodeWithCharIds
	for _, epsds := range e.episodes {
		for _, chr := range epsds.Characters {
			idIndex := strings.LastIndex(chr, "/")
			charIdsSlc[epsds.Episode] = append(charIdsSlc[epsds.Episode], chr[idIndex+1:])
		}
		charIdsSingle := EpisodeWithCharIds{
			EpisodeName:  epsds.Name,
			EpisodeCode:  epsds.Episode,
			CharacterIds: charIdsSlc[epsds.Episode],
		}
		charIds = append(charIds, charIdsSingle)
	}
	// return CharIdsEpisodeObj{CharIds: charIds}
	return charIds
}

type EpisodeLocations struct {
	Name      string
	Episode   string   `json:"episode"`
	Locations []string `json:"locations"`
}

// episodeLocations takes all the locations (origins) per character and all the character ids per episode
// and returns the locations per episode (unique values)
func episodeLocations(locsPerCharId map[string]string, epiCharIds []EpisodeWithCharIds) []EpisodeLocations {
	episodeLocationMap := make(map[string][]string)
	var locEpi []EpisodeLocations
	var locEpiSingle EpisodeLocations
	for _, v := range epiCharIds {
		for _, vv := range v.CharacterIds {
			episodeLocationMap[v.EpisodeCode] = append(episodeLocationMap[v.EpisodeCode], locsPerCharId[vv])
			locEpiSingle = EpisodeLocations{Name: v.EpisodeName, Episode: v.EpisodeCode}

		}
		locEpiSingle.Locations = removeDuplicateStr(episodeLocationMap[v.EpisodeCode])
		locEpi = append(locEpi, locEpiSingle)
	}

	return locEpi
}
