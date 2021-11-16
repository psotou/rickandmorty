package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"reflect"
	"testing"
)

func Test_getLocations(t *testing.T) {
	tests := []struct {
		name string
		want iLocation
	}{
		{
			name: "returns the location object",
			want: locationInterfaceGenerator(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLocations(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLocations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocationsObj_countChar(t *testing.T) {
	type fields struct {
		locations []LocationResults
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
			name:   "counts the ocurrence of a certain character in the field name",
			fields: fields{locationObjGenerator().locations},
			args:   args{"e"},
			want:   955,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loc := &LocationsObj{
				locations: tt.fields.locations,
			}
			if got := loc.countChar(tt.args.char); got != tt.want {
				t.Errorf("LocationsObj.countChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func locationInterfaceGenerator() iLocation {
	var locRes []LocationResults
	jsonFile, _ := os.Open("fixtures/location.json")
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	err := json.Unmarshal(byteValue, &locRes)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return &LocationsObj{locations: locRes}
}

func locationObjGenerator() LocationsObj {
	var locRes []LocationResults
	jsonFile, _ := os.Open("fixtures/character.json")
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	err := json.Unmarshal(byteValue, &locRes)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return LocationsObj{locations: locRes}
}
