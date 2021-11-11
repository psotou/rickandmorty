package main

import "fmt"

func main() {
	// locationNames := getLocationNames()
	// characterNames := getCharacterNames()
	// episodeNames := getEpisodeNames()
	// fmt.Println(episodeNames.characterIdsPerEpisode())

	// charIdsPerEpi := episodeNames.characterIdsPerEpisode()
	// fmt.Println(charIdsPerEpi)

	// charIdLocNames := characterNames.charIdWithLocationName()
	// locEpi := charIdsPerEpi.locationPerEpisode(charIdLocNames)
	// fmt.Println(locEpi)

	// charIdOrigin := charIdWithLocationName()
	// // fmt.Println(charIdOrigin)

	// loc := locationPerEpisode(charIdsPerEpi, charIdOrigin)

	// for k, v := range loc {
	// 	fmt.Println(k)
	// 	fmt.Println(v)
	// }

	// location := GetResource(Location)
	// // fmt.Println(location)
	// for _, v := range location.Locations {
	// 	fmt.Println(v.Name)
	// }

	// fmt.Println(location.CountChar("ζ"))

	// character := GetResource(Character)
	// for _, v := range character.Characters {
	// 	fmt.Println(v.Id, v.Name)
	// }

	// fmt.Println(character.CountChar("7+7"))

	// a := characterNames.charIdWithLocationName()
	// e := episodeNames.locationPerEpisode(a)
	// fmt.Println(e)

	cc, _ := charCounterResult()
	fmt.Println(string(cc))

	ee, _ := episodeLocations()
	fmt.Println(string(ee))

	// 	aa := locationsEpisodeStruct()
	// 	fmt.Println(aa)
}
