package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// episodeResults cointains the necessary information to map all
// locations used per episode
type EpisodeResults struct {
	Id         int      `json:"id"`
	Name       string   `json:"name"`
	Episode    string   `json:"episode"`
	Characters []string `json:"characters"`
}

type EpisodeObj struct {
	episodes []EpisodeResults
}

// getEpisodeNames returns the names of all the episodes
func getEpisodeNames() EpisodeObj {
	var episodeResults []EpisodeResults
	episodeNumber := getInfo(Episode).Count
	episodeRange := makeRange(1, episodeNumber)
	episodeWithIdsURL := fmt.Sprintf("%s%s", Episode, sliceToString(episodeRange))

	episodeData, _ := getReq(episodeWithIdsURL)
	err := json.Unmarshal(episodeData, &episodeResults)
	if err != nil {
		log.Fatal(err.Error())
	}
	return EpisodeObj{episodes: episodeResults}
}

// countChar counts the number of ocurrences of a certain character
func (c *EpisodeObj) countChar(char string) int {
	var count int
	for _, v := range c.episodes {
		count += strings.Count(v.Name, char)
	}
	return count
}

// custom types for the characterIdsPerEpisode function
type charIdsPerEpisode map[string][]string

type CharIdsEpisode struct {
	EpisodeName  string
	EpisodeCode  string
	CharacterIds []string
}

// characterIdsPerEpisode maps every episode with a slice containing
// the characters ids that appeared in said episode
// func (e *EpisodeObj) characterIdsPerEpisode() map[string][]string {
func (e *EpisodeObj) characterIdsPerEpisode() CharIdsEpisodeObj {
	charIdsSlc := make(charIdsPerEpisode)
	var charIds []CharIdsEpisode
	for _, epsds := range e.episodes {
		for _, chr := range epsds.Characters {
			idIndex := strings.LastIndex(chr, "/")
			charIdsSlc[epsds.Episode] = append(charIdsSlc[epsds.Episode], chr[idIndex+1:])
		}
		charIdsSingle := CharIdsEpisode{EpisodeName: epsds.Name, EpisodeCode: epsds.Episode, CharacterIds: charIdsSlc[epsds.Episode]}
		charIds = append(charIds, charIdsSingle)
	}
	// return charIdsSlc
	return CharIdsEpisodeObj{CharIds: charIds}
}

// custom type for the locationPerEpisode function
type locationsPerCharId map[string]string

type CharIdsEpisodeObj struct {
	CharIds []CharIdsEpisode
}

type LocEpiObj struct {
	Name      string
	Episode   string   `json:"episode"`
	Locations []string `json:"locations"`
}

// locationPerEpisode takes all the character ids per episode and all the locations (origins) per character
// and returns a map with the episode code with all the locations per episode (unique values)
// func (c *CharIdsEpisodeObj) locationPerEpisode(locsPerCharId locationsPerCharId) map[string][]string {
func (c *CharIdsEpisodeObj) locationPerEpisode(locsPerCharId locationsPerCharId) []LocEpiObj {
	episodeLocationMap := make(map[string][]string)
	var locEpi []LocEpiObj
	var locEpiSingle LocEpiObj
	// for k, v := range c.characterIdsPerEpisode() {
	for _, v := range c.CharIds {
		for _, vv := range v.CharacterIds {
			episodeLocationMap[v.EpisodeCode] = append(episodeLocationMap[v.EpisodeCode], locsPerCharId[vv])
			locEpiSingle = LocEpiObj{Name: v.EpisodeName, Episode: v.EpisodeCode}

		}
		locEpiSingle.Locations = removeDuplicateStr(episodeLocationMap[v.EpisodeCode])
		// episodeLocationMap[k] = removeDuplicateStr(episodeLocationMap[k])
		// episodeLocationMap[v.EpisodeCode] = removeDuplicateStr(episodeLocationMap[v.EpisodeCode])
		locEpi = append(locEpi, locEpiSingle)
	}
	// return episodeLocationMap
	return locEpi
}
