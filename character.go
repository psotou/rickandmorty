package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type characterResults struct {
	Name string `json:"name"`
}

type characterObj struct {
	characters []characterResults
}

func getCharacterNames() characterObj {
	var charResults []characterResults
	characterNumber := getInfo(Character).Count
	characterRange := makerange(1, characterNumber)
	characterWithIdsURL := fmt.Sprintf("%s%s", Character, sliceToString(characterRange))

	characterData, _ := getReq(characterWithIdsURL)
	err := json.Unmarshal(characterData, &charResults)
	if err != nil {
		log.Fatal(err.Error())
	}
	return characterObj{characters: charResults}
}

// this approach has a better performance
func (c *characterObj) countChar(char string) int {
	var count int
	for _, v := range c.characters {
		count += strings.Count(v.Name, char)
	}
	return count
}
