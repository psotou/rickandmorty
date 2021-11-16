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
	tests := []struct {
		name string
		want []CharCounterResults
	}{
		{
			name: "",
			want: charCounterGenerator(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := charCounterResult().Results; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("charCounterResult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func charCounterGenerator() []CharCounterResults {
	var charRes CharCounter
	jsonFile, _ := os.Open("fixtures/char_counter.json")
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	err := json.Unmarshal(byteValue, &charRes)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return charRes.Results
}
