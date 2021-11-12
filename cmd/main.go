package main

import "fmt"

func main() {
	result := resultJSON()
	writeJSON(result)

	fmt.Println(episodeLocationsResult())
}
