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
	tests := struct {
		name string
		want iLocation
	}{
		name: "returns the location object",
		want: locationInterfaceGenerator(),
	}

	rng := makeRange(1, 5)
	t.Run(tests.name, func(t *testing.T) {
		if got := getLocations(rng); !reflect.DeepEqual(got, tests.want) {
			t.Errorf("getLocations() = %v, want %v", got, tests.want)
		}
	})
}

func TestLocationsObj_countChar(t *testing.T) {
	type fields struct {
		locations []LocationResults
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
		name:   "counts the ocurrence of a certain character in the field name",
		fields: fields{locationObjGenerator().locations},
		args:   args{"l"},
		want:   3,
	}

	t.Run(tests.name, func(t *testing.T) {
		loc := &LocationsObj{
			locations: tests.fields.locations,
		}
		if got := loc.countChar(tests.args.char); got != tests.want {
			t.Errorf("LocationsObj.countChar() = %v, want %v", got, tests.want)
		}
	})
}

//
// Helper methods for the locations testing
//
func locationInterfaceGenerator() iLocation {
	locRes := []LocationResults{}
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
	locRes := []LocationResults{}
	jsonFile, _ := os.Open("fixtures/location.json")
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	err := json.Unmarshal(byteValue, &locRes)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return LocationsObj{locations: locRes}
}
