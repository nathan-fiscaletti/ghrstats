package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"log"

	"github.com/nathan-fiscaletti/ghrstats/pkg/ghrstats"
)

func main() {
	args, err := getArgs()
	if err != nil {
		log.Fatalf("%v", err.Error())
	}

	downloadCount, err := ghrstats.GetDownloadsForRepository(args.repo, args.filter)
	if err != nil {
		log.Fatalf("Error fetching downloads: %v", err)
	}

	fmt.Printf("%d\n", downloadCount)
}

type arguments struct {
	repo   string
	filter func(ghrstats.Asset) bool
}

func getArgs() (*arguments, error) {
	var res arguments

	patterns := flag.String("patterns", "", "File path patterns to filter by, separated by commas")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		return nil, fmt.Errorf("usage: %v [-patterns *.ext1,*.ext2,...] <repo>", os.Args[0])
	}

	res.repo = args[0]
	if patterns != nil && *patterns != "" {
		res.filter = ghrstats.ByFileNamePatterns(strings.Split(*patterns, ",")...)
	}

	return &res, nil
}
