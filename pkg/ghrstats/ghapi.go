package ghrstats

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// API_ROOT is the root URL for the GitHub API
const API_ROOT = "https://api.github.com"

// Request fetches a JSON response from the GitHub API and decodes it into the provided type
func Request[R any](path string) (*R, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s", API_ROOT, path))
	if err != nil {
		return nil, fmt.Errorf("error fetching URL: %v", err)
	}
	defer resp.Body.Close()

	var res R
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, fmt.Errorf("error decoding JSON: %v", err)
	}

	return &res, nil
}

// RequestMany fetches a JSON response from the GitHub API and decodes it into a slice of the provided type
func RequestMany[R any](path string) ([]R, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s", API_ROOT, path))
	if err != nil {
		return nil, fmt.Errorf("error fetching URL: %v", err)
	}
	defer resp.Body.Close()

	var res []R
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, fmt.Errorf("error decoding JSON: %v", err)
	}

	return res, nil
}
