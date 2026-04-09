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
		downloadCount, err := ghrstats.GetDownloadsForRepository(ghrstats.GetDownloadsForRepositoryRequest{
			Repo:      args.Repo,
			Tag:       args.Tag,
			Predicate: args.Filter,
		})
		if err != nil {
			log.Fatalf("Error fetching downloads: %v", err)
		}

		data["downloads"] = downloadCount
	case cli.ActionAggregateItemized:
		aggregate, err := ghrstats.AggregateDownloadCount(ghrstats.AggregateDownloadCountRequest{
			Repo:      args.Repo,
			Tag:       args.Tag,
			Predicate: args.Filter,
		})
		if err != nil {
			log.Fatalf("Error aggregating downloads: %v", err)
		}

		for asset, count := range aggregate {
			if v, ok := data[asset.Name]; ok {
				data[asset.Name] = v.(int) + count
			} else {
				data[asset.Name] = count
			}
		}
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}

	fmt.Println(string(jsonData))
}
