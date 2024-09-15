package cli

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/nathan-fiscaletti/ghrstats/pkg/ghrstats"
)

const USAGE = `usage: %v -r <repo> [-a <total|itemized>] [-f *.ext1,*.ext2,...]`

type Arguments struct {
	Repo   string
	Action Action
	Filter func(ghrstats.Asset) bool
}

func GetArguments() (*Arguments, error) {
	var res Arguments

	repo := flag.String("r", "", "The repository to fetch download stats for")
	action := flag.String("a", string(ActionAggregateTotal), "The action to perform")
	filter := flag.String("f", "", "File path patterns to filter by, separated by commas")
	flag.Parse()

	if repo == nil || *repo == "" {
		return nil, fmt.Errorf(USAGE, os.Args[0])
	}
	res.Repo = *repo

	if action != nil {
		var found bool
		for _, a := range Actions {
			if Action(*action) == a {
				res.Action = a
				found = true
				break
			}
		}

		if !found {
			return nil, fmt.Errorf("invalid action: %v", *action)
		}
	}

	if filter != nil && *filter != "" {
		res.Filter = ghrstats.ByFileNamePatterns(strings.Split(*filter, ",")...)
	}

	return &res, nil
}
