package main

import (
	"fmt"
)

func main() {
	locationNames := getLocationNames()
	charatcerNames := getCharacterNames()
	episodeNames := getEpisodeNames()

	fmt.Println(locationNames.countChar("l"))
	fmt.Println(charatcerNames.countChar("c"))
	fmt.Println(episodeNames.countChar("e"))

}
