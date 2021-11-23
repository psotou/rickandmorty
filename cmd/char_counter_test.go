package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"reflect"
	"testing"
)

func Test_charCounterResult(t *testing.T) {
	tests := struct {
		name string
		want []CharCounterResults
	}{
		name: "",
		want: charCounterGenerator(),
	}

	resourceRangeOfIds := []ResourceRange{
		{
			Resource: "location",
			Range:    []string{"1", "2", "3", "4", "5"},
		},
		{
			Resource: "episode",
			Range:    []string{"1", "2", "3", "4", "5"},
		},
		{
			Resource: "character",
			Range:    []string{"1", "2", "3", "4", "5"},
		},
	}

	t.Run(tests.name, func(t *testing.T) {
		if got := charCounterResult(resourceRangeOfIds).Results; !reflect.DeepEqual(got, tests.want) {
			t.Errorf("charCounterResult() = %v, want %v", got, tests.want)
		}
	})
}

func charCounterGenerator() []CharCounterResults {
	charRes := CharCounter{}
	jsonFile, _ := os.Open("fixtures/char_counter.json")
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	err := json.Unmarshal(byteValue, &charRes)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return charRes.Results
}
