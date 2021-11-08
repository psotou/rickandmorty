package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type episodeResults struct {
	Name string `json:"name"`
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
