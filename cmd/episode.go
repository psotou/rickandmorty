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
	characterIds() []EpisodeWithCharIds
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
	// episodeEndpointMultipleIds returns the ids in range to fetch multiple episodes
	// See https://rickandmortyapi.com/documentation/#get-multiple-episodes
	episodeEndpointMultipleIds := fmt.Sprintf("%s%s", Episode, sliceToString(episodeRange))

	episodeData, _ := getReq(episodeEndpointMultipleIds)
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
	for _, episode := range c.episodes {
		count += strings.Count(episode.Name, char)
	}
	return count
}

// EpisodeWithCharIds is similar to EpisodeResults struct
// however, it stores the ids of the characters instead of the endpoint
type EpisodeWithCharIds struct {
	EpisodeName  string
	EpisodeCode  string
	CharacterIds []string
}

// characterIds maps every episode with a slice containing
// the characters ids that appeared in said episode
func (e *EpisodeObj) characterIds() []EpisodeWithCharIds {
	characterIdsMap := make(map[string][]string)
	var charIds []EpisodeWithCharIds
	for _, episode := range e.episodes {
		for _, character := range episode.Characters {
			idIndex := strings.LastIndex(character, "/")
			characterIdsMap[episode.Episode] = append(characterIdsMap[episode.Episode], character[idIndex+1:])
		}
		charIdsSingle := EpisodeWithCharIds{
			EpisodeName:  episode.Name,
			EpisodeCode:  episode.Episode,
			CharacterIds: characterIdsMap[episode.Episode],
		}
		charIds = append(charIds, charIdsSingle)
	}
	return charIds
}
