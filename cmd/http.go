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

// HTTP client interface which allows to set instances of
// either http.Client or our mock http client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var Client HTTPClient

func init() {
	Client = &http.Client{}
}

// MockClient sets the function that our mock Do method will run
// instead of the http.Client.Do method
type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

// Do method that overrides the http.Client.Do method
func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

func formatURL(endpoint string) string {
	return fmt.Sprintf("%s%s", BaseURL, endpoint)
}

func requestMethod(reqMethod, endpoint string) (*http.Response, error) {
	url := formatURL(endpoint)
	req, err := http.NewRequest(reqMethod, url, nil)
	if err != nil {
		return nil, err
	}

	// client := &http.Client{}
	// res, err := client.Do(req)
	res, err := Client.Do(req)
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
