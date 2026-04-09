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
	Tag    string  `json:"tag_name"`
}

// GetReleases fetches the releases for a given repository
func GetReleases(repo string) ([]Release, error) {
	return RequestMany[Release](fmt.Sprintf("repos/%s/releases", repo))
}

type AggregateDownloadCountRequest struct {
	Repo      string           `json:"repo"`
	Tag       string           `json:"tag"`
	Predicate func(Asset) bool `json:"predicate"`
}

// AggregateDownloadCount aggregates the download count for each asset, applying
// a predicate to filter the assets. If the predicate is nil, all assets will be included
func AggregateDownloadCount(req AggregateDownloadCountRequest) (map[Asset]int, error) {
	releases, err := GetReleases(req.Repo)
	if err != nil {
		return nil, err
	}

	filteredReleases := make([]Release, 0, len(releases))
	for _, release := range releases {
		if req.Tag == "" || release.Tag == req.Tag {
			filteredReleases = append(filteredReleases, release)
		}
	}

	totalCount := make(map[Asset]int)

	for _, release := range filteredReleases {
		for _, asset := range release.Assets {
			if req.Predicate == nil || req.Predicate(asset) {
				totalCount[asset] += asset.DownloadCount
			}
		}
	}

	return totalCount, nil
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

type GetDownloadsForRepositoryRequest struct {
	Repo      string           `json:"repo"`
	Tag       string           `json:"tag"`
	Predicate func(Asset) bool `json:"predicate"`
}

// GetDownloadsForRepository fetches the download count for a given repository
// and applies a predicate to filter the assets. If the predicate is nil, all assets will be included
func GetDownloadsForRepository(req GetDownloadsForRepositoryRequest) (int, error) {
	assetCounts, err := AggregateDownloadCount(AggregateDownloadCountRequest{
		Repo:      req.Repo,
		Tag:       req.Tag,
		Predicate: req.Predicate,
	})
	if err != nil {
		return 0, err
	}

	var total int

	for _, count := range assetCounts {
		total += count
	}

	return total, nil
}
