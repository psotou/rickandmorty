package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type characterResults struct {
	Id       int             `json:"id"`
	Name     string          `json:"name"`
	Location characterOrigin `json:"origin"`
	// Episode  []string        `json:"episode"`
}

type characterOrigin struct {
	Name string `json:"name"`
	Url  string `json:"url"`
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

// returns a map with character id and location (origin) name
func charIdWithOriginName() map[string]string {
	charIdOrigin := make(map[string]string)
	for _, v := range getCharacterNames().characters {
		strId := strconv.Itoa(v.Id)
		charIdOrigin[strId] = v.Location.Name
	}
	return charIdOrigin
}
