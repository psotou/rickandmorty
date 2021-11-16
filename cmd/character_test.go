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
	tests := []struct {
		name string
		want iCharacter
	}{
		{
			name: "returns the characters object",
			want: characterInterfaceGenerator(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCharacters(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCharacterObj_countChar(t *testing.T) {
	type fields struct {
		characters []CharacterResults
	}
	type args struct {
		char string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "the ocurrence of a certain character in the field name",
			fields: fields{characterObjGenerator().characters},
			args:   args{"a"},
			want:   710,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CharacterObj{
				characters: tt.fields.characters,
			}
			if got := c.countChar(tt.args.char); got != tt.want {
				t.Errorf("CharacterObj.countChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCharacterObj_locationName(t *testing.T) {
	type fields struct {
		characters []CharacterResults
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]string
	}{
		{
			name:   "",
			fields: fields{characterObjGenerator().characters},
			want:   locationNameGenerator(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CharacterObj{
				characters: tt.fields.characters,
			}
			if got := c.locationName(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CharacterObj.locationName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func characterInterfaceGenerator() iCharacter {
	var charRes []CharacterResults
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
	var charRes []CharacterResults
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
	var charRes = make(map[string]string)
	jsonFile, _ := os.Open("fixtures/location_names.json")
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	err := json.Unmarshal(byteValue, &charRes)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return charRes
}
