package main

import (
	"fmt"
)

func main() {
	// locationNames := getLocationNames()
	// characterNames := getCharacterNames()
	episodeNames := getEpisodeNames()

	// fmt.Println(locationNames.countChar("l"))
	// fmt.Println(charatcerNames.countChar("c"))
	// fmt.Println(episodeNames.countChar("e"))

	// mapEpiLoc := make(map[string][]string)
	// for _, v := range episodeNames.episodes {
	// fmt.Println(v.Episode)
	// fmt.Println(v.Characters)

	// for _, val := range v.Characters {
	// 	idIndex := strings.LastIndex(val, "/")
	// 	mapEpiLoc[v.Episode] = append(mapEpiLoc[v.Episode], val[idIndex+1:])
	// }
	// fmt.Println(mapEpiLoc)

	// stringyfy := strings.Join(mapEpiLoc[v.Episode], ",")
	// fmt.Println(getCharacterLocation(stringyfy))

	charIdsPerEpi := episodeNames.characterIdsPerEpisode()
	// fmt.Println(charIdsPerEpi)

	charIdOrigin := charIdWithOriginName()
	// fmt.Println(charIdOrigin)

	// wrap this into a function for this returns the episodes with the locations (origin) of characters
	// episodeLocationMap := make(map[string][]string)
	// for k, v := range charIdsPerEpi {
	// 	for _, vv := range v {
	// 		episodeLocationMap[k] = append(episodeLocationMap[k], charIdOrigin[vv])
	// 	}
	// }
	// fmt.Println(episodeLocationMap)

	loc := locationPerEpisode(charIdsPerEpi, charIdOrigin)
	// fmt.Println(loc)

	for k, v := range loc {
		fmt.Println(k)
		fmt.Println(v)
	}
}
