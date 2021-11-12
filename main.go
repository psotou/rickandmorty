package main

import "fmt"

func main() {

	ee := episodeLocationsResult()
	cc := charCounterResult()

	fmt.Println(ee)
	fmt.Println(cc)

	// e, _ := toJSON(ee)
	// fmt.Println(string(e))

	// c, _ := toJSON(cc)
	// fmt.Println(string(c))
}
