package cli

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/nathan-fiscaletti/ghrstats/pkg/ghrstats"
)

type Arguments struct {
	Repo   string
	Filter func(ghrstats.Asset) bool
}

func GetArguments() (*Arguments, error) {
	var res Arguments

	patterns := flag.String("patterns", "", "File path patterns to filter by, separated by commas")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		return nil, fmt.Errorf("usage: %v [-patterns *.ext1,*.ext2,...] <repo>", os.Args[0])
	}

	res.Repo = args[0]
	if patterns != nil && *patterns != "" {
		res.Filter = ghrstats.ByFileNamePatterns(strings.Split(*patterns, ",")...)
	}

	return &res, nil
}
