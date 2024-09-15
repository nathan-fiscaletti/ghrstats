package ghrstats

import (
	"fmt"
	"log"
	"path/filepath"
)

// Asset represents a GitHub release asset
type Asset struct {
	Name          string `json:"name"`
	DownloadCount int    `json:"download_count"`
}

// Release represents a GitHub release
type Release struct {
	Assets []Asset `json:"assets"`
}

// GetReleases fetches the releases for a given repository
func GetReleases[R any](repo string) ([]R, error) {
	return RequestMany[R](fmt.Sprintf("repos/%s/releases", repo))
}

// AggregateDownloadCount aggregates the download count for each asset, applying
// a predicate to filter the assets. If the predicate is nil, all assets will be included
func AggregateDownloadCount(releases []Release, predicate func(Asset) bool) map[Asset]int {
	totalCount := make(map[Asset]int)
	for _, release := range releases {
		for _, asset := range release.Assets {
			if predicate == nil || predicate(asset) {
				totalCount[asset] += asset.DownloadCount
			}
		}
	}
	return totalCount
}

// ByFileNamePatterns returns a predicate function that filters assets by file extension
// The predicate will return true if the asset name matches any of the provided patterns
var ByFileNamePatterns = func(patterns ...string) func(Asset) bool {
	return func(asset Asset) bool {
		var matched bool

		for _, pattern := range patterns {
			var err error
			matched, err = filepath.Match(pattern, asset.Name)
			if err != nil {
				log.Fatalf("Error matching pattern: %v", err)
			}

			if matched {
				break
			}
		}

		return matched
	}
}

// GetDownloadsForRepository fetches the download count for a given repository
// and applies a predicate to filter the assets. If the predicate is nil, all assets will be included
func GetDownloadsForRepository(repo string, predicate func(Asset) bool) (int, error) {
	releases, err := GetReleases[Release](repo)
	if err != nil {
		return 0, err
	}

	assetCounts := AggregateDownloadCount(releases, predicate)
	var total int
	for _, count := range assetCounts {
		total += count
	}
	return total, nil
}
