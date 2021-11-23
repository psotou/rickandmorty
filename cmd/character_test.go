package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"reflect"
	"testing"
)

func Test_getCharacters(t *testing.T) {
	tests := struct {
		name string
		want iCharacter
	}{
		name: "returns the characters object",
		want: characterInterfaceGenerator(),
	}

	rng := []string{"1", "2", "3", "4", "5"}
	t.Run(tests.name, func(t *testing.T) {
		if got := getCharacters(rng); !reflect.DeepEqual(got, tests.want) {
			t.Errorf("getCharacters() = %v, want %v", got, tests.want)
		}
	})
}

func TestCharacterObj_countChar(t *testing.T) {
	type fields struct {
		characters []CharacterResults
	}
	type args struct {
		char string
	}
	tests := struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		name:   "the ocurrence of a certain character in the field name",
		fields: fields{characterObjGenerator().characters},
		args:   args{"c"},
		want:   2,
	}

	t.Run(tests.name, func(t *testing.T) {
		c := &CharacterObj{
			characters: tests.fields.characters,
		}
		if got := c.countChar(tests.args.char); got != tests.want {
			t.Errorf("CharacterObj.countChar() = %v, want %v", got, tests.want)
		}
	})
}

func TestCharacterObj_locationName(t *testing.T) {
	type fields struct {
		characters []CharacterResults
	}
	tests := struct {
		name   string
		fields fields
		want   map[string]string
	}{
		name:   "",
		fields: fields{characterObjGenerator().characters},
		want:   locationNameGenerator(),
	}

	t.Run(tests.name, func(t *testing.T) {
		c := &CharacterObj{
			characters: tests.fields.characters,
		}
		if got := c.locationName(); !reflect.DeepEqual(got, tests.want) {
			t.Errorf("CharacterObj.locationName() = %v, want %v", got, tests.want)
		}
	})
}

//
// Helper methods for the character testing
//
func characterInterfaceGenerator() iCharacter {
	charRes := []CharacterResults{}
	jsonFile, _ := os.Open("fixtures/character.json")
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	err := json.Unmarshal(byteValue, &charRes)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return &CharacterObj{characters: charRes}
}

func characterObjGenerator() CharacterObj {
	charRes := []CharacterResults{}
	jsonFile, _ := os.Open("fixtures/character.json")
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	err := json.Unmarshal(byteValue, &charRes)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return CharacterObj{characters: charRes}
}

func locationNameGenerator() map[string]string {
	charRes := make(map[string]string)
	jsonFile, _ := os.Open("fixtures/location_names.json")
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	err := json.Unmarshal(byteValue, &charRes)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return charRes
}
