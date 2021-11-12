package main

import "encoding/json"

type FinalJSON struct {
	*CharCounter
	*EpiLocations
}

func (f *FinalJSON) toJSON() ([]byte, error) {
	return json.Marshal(f)
}
