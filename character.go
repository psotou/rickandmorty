package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type CharacterResults struct {
	Id       int             `json:"id"`
	Name     string          `json:"name"`
	Location CharacterOrigin `json:"origin"`
	// Episode  []string        `json:"episode"`
}

type CharacterOrigin struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type CharacterObj struct {
	characters []CharacterResults
}

func getCharacterNames() CharacterObj {
	var charResults []CharacterResults
	characterNumber := getInfo(Character).Count
	characterRange := makeRange(1, characterNumber)
	characterWithIdsURL := fmt.Sprintf("%s%s", Character, sliceToString(characterRange))

	characterData, _ := getReq(characterWithIdsURL)
	err := json.Unmarshal(characterData, &charResults)
	if err != nil {
		log.Fatal(err.Error())
	}
	return CharacterObj{characters: charResults}
}

// this approach has a better performance
func (c *CharacterObj) countChar(char string) int {
	var count int
	for _, v := range c.characters {
		count += strings.Count(v.Name, char)
	}
	return count
}

// returns a map with character id location (origin) name for that character
func (c *CharacterObj) charIdWithLocationName() map[string]string {
	charIdOrigin := make(map[string]string)
	for _, v := range c.characters {
		strId := strconv.Itoa(v.Id)
		charIdOrigin[strId] = v.Location.Name
	}
	return charIdOrigin
}
