package main

import (
	"fmt"
	"io"
	"net/http"
)

const (
	BaseURL   = "https://rickandmortyapi.com/api/"
	Location  = "location/"
	Episode   = "episode/"
	Character = "character/"
)

func formatURL(endpoint string) string {
	return fmt.Sprintf("%s%s", BaseURL, endpoint)
}

func requestMethod(reqMethod, endpoint string) (*http.Response, error) {
	url := formatURL(endpoint)
	req, err := http.NewRequest(reqMethod, url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func getReq(endpoint string) ([]byte, error) {
	res, err := requestMethod(http.MethodGet, endpoint)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
