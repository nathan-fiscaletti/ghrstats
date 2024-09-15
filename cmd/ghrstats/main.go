package main

import (
	"encoding/json"
	"fmt"

	"log"

	"github.com/nathan-fiscaletti/ghrstats/internal/cli"
	"github.com/nathan-fiscaletti/ghrstats/pkg/ghrstats"
)

func main() {
	args, err := cli.GetArguments()
	if err != nil {
		log.Fatalf("%v", err.Error())
	}

	data := map[string]any{}

	switch args.Action {
	case cli.ActionAggregateTotal:
		downloadCount, err := ghrstats.GetDownloadsForRepository(args.Repo, args.Filter)
		if err != nil {
			log.Fatalf("Error fetching downloads: %v", err)
		}

		data["aggregate_downloads"] = downloadCount
	case cli.ActionAggregateItemized:
		releases, err := ghrstats.GetReleases[ghrstats.Release](args.Repo)
		if err != nil {
			log.Fatalf("Error fetching releases: %v", err)
		}

		aggregate := ghrstats.AggregateDownloadCount(releases, args.Filter)
		for asset, count := range aggregate {
			data[asset.Name] = count
		}
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}

	fmt.Println(string(jsonData))
}
