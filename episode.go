package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// episodeResults cointains the necessary information to map all
// locations used per episode
type episodeResults struct {
	Id         int      `json:"id"`
	Name       string   `json:"name"`
	Episode    string   `json:"episode"`
	Characters []string `json:"characters"`
}

type episodeObj struct {
	episodes []episodeResults
}

func getEpisodeNames() episodeObj {
	var episodeResults []episodeResults
	episodeNumber := getInfo(Episode).Count
	episodeRange := makerange(1, episodeNumber)
	episodeWithIdsURL := fmt.Sprintf("%s%s", Episode, sliceToString(episodeRange))

	episodeData, _ := getReq(episodeWithIdsURL)
	err := json.Unmarshal(episodeData, &episodeResults)
	if err != nil {
		log.Fatal(err.Error())
	}
	return episodeObj{episodes: episodeResults}
}

// this approach has a better performance
func (c *episodeObj) countChar(char string) int {
	var count int
	for _, v := range c.episodes {
		count += strings.Count(v.Name, char)
	}
	return count
}

// type episodeCharIds map[string][]string

// characterOrigin maps every episode with the location (origin) of every character
// that appeared in the episode
func (c *episodeObj) characterIdsPerEpisode() map[string][]string {
	charIdsSlc := make(map[string][]string)
	// charIdsStr := make(map[string]string)
	for _, epsds := range c.episodes {
		for _, chr := range epsds.Characters {
			idIndex := strings.LastIndex(chr, "/")
			charIdsSlc[epsds.Episode] = append(charIdsSlc[epsds.Episode], chr[idIndex+1:])
		}
		// charIdsStr[epsds.Episode] = sliceToString(charIdsSlc[epsds.Episode])
	}
	return charIdsSlc
}
