package ghrstats

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

const API_ROOT = "https://api.github.com"

type Asset struct {
	Name          string `json:"name"`
	DownloadCount int    `json:"download_count"`
}

type Release struct {
	Assets []Asset `json:"assets"`
}

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

// GetReleases fetches the releases for a given repository
func GetReleases[R any](repo string) ([]R, error) {
	return RequestMany[R](fmt.Sprintf("repos/%s/releases", repo))
}

// AggregateDownloadCount aggregates the download count for each asset
func AggregateDownloadCount(releases []Release) map[Asset]int {
	totalCount := make(map[Asset]int)
	for _, release := range releases {
		for _, asset := range release.Assets {
			totalCount[asset] += asset.DownloadCount
		}
	}
	return totalCount
}

// ByFileNamePatterns returns a predicate function that filters assets by file extension
// The predicate will return true if the asset name matches any of the provided patterns
var ByFileNamePatterns = func(patterns ...string) func(Asset) bool {
	return func(asset Asset) bool {
		for _, pattern := range patterns {
			matched, err := filepath.Match(pattern, asset.Name)
			if err != nil {
				log.Fatalf("Error matching pattern: %v", err)
			}

			if matched {
				return true
			}
		}

		return false
	}
}

// GetDownloadsForRepository fetches the download count for a given repository
// and applies a predicate to filter the assets. If the predicate is nil, all assets will be included
func GetDownloadsForRepository(repo string, predicate func(Asset) bool) (int, error) {
	releases, err := GetReleases[Release](repo)
	assetCounts := AggregateDownloadCount(releases)
	var total int
	for asset, count := range assetCounts {
		if predicate == nil || predicate(asset) {
			total += count
		}
	}
	return total, err
}
