package ghrstats

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

const API_URL = "https://api.github.com/repos/%s/releases"

type Asset struct {
	Name          string `json:"name"`
	DownloadCount int    `json:"download_count"`
}

type Release struct {
	Assets []Asset `json:"assets"`
}

// GetReleases fetches the releases for a given repository
func GetReleases[R any](repo string) ([]R, error) {
	// Perform the GET request
	resp, err := http.Get(fmt.Sprintf(API_URL, repo))
	if err != nil {
		return nil, fmt.Errorf("error fetching URL: %v", err)
	}
	defer resp.Body.Close()

	// Decode the JSON response
	var releases []R
	if err := json.NewDecoder(resp.Body).Decode(&releases); err != nil {
		return nil, fmt.Errorf("error decoding JSON: %v", err)
	}

	return releases, nil
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
