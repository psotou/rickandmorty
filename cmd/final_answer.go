package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

func resultJSON() string {
	charCount := charCounterResult()
	locNames := episodeLocationsResult()
	res := []interface{}{charCount, locNames}
	dataBytes, err := json.Marshal(res)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return string(dataBytes)
}

func writeJSON(data string) {
	// return os.WriteFile("result", data, 0644)
	file, _ := os.Create("result.json")
	defer file.Close()

	// _, err := file.Write(data)
	_, err := file.WriteString(string(data))
	if err != nil {
		errors.New("Couldn't write to file reslut.json")
		os.Exit(1)
	}
}
