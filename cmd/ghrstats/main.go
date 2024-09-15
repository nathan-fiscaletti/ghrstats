package main

import (
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

	downloadCount, err := ghrstats.GetDownloadsForRepository(args.Repo, args.Filter)
	if err != nil {
		log.Fatalf("Error fetching downloads: %v", err)
	}

	fmt.Printf("%d\n", downloadCount)
}
