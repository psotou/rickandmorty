package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// iCharacter interface defines the methods associated with the CharacterObj struct
type iCharacter interface {
	countChar(string) int
	locationName() map[string]string
}

// CharacterObj is the struct type that implemts the iCharacter interface
// Contains a collection of the object CharacterResults
type CharacterObj struct {
	characters []CharacterResults
}
type CharacterResults struct {
	Id       int             `json:"id"`
	Name     string          `json:"name"`
	Location CharacterOrigin `json:"origin"`
}

type CharacterOrigin struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

// getCharacters embeds CharaterObj struct and indirectly implements
// the iCharacter interface. This approach allows for the use of a syntax
// like getCharacters().countChar() declared in one line
func getCharacters() iCharacter {
	var characterResults []CharacterResults
	characterNumber := getInfo(Character).Count
	characterRange := makeRange(1, characterNumber)
	// characterEndpointMultipleIds returns the ids in range to fetch multiple characters
	// See https://rickandmortyapi.com/documentation/#get-multiple-characters
	characterEndpointMultipleIds := fmt.Sprintf("%s%s", Character, sliceToString(characterRange))

	characterData, _ := getReq(characterEndpointMultipleIds)
	err := json.Unmarshal(characterData, &characterResults)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &CharacterObj{characters: characterResults}
}

// countChar method implemented on the CharacterObj struct
// Counts the ocurrence of a certain character in the CharacterResult.Name field
func (c *CharacterObj) countChar(char string) int {
	var count int
	for _, character := range c.characters {
		count += strings.Count(character.Name, char)
	}
	return count
}

// locationName method implemented on the CharacterObj struct
// returns a map with character id and the location (origin) name for that character
func (c *CharacterObj) locationName() map[string]string {
	charIdOrigin := make(map[string]string)
	for _, character := range c.characters {
		strId := strconv.Itoa(character.Id)
		charIdOrigin[strId] = character.Location.Name
	}
	return charIdOrigin
}
