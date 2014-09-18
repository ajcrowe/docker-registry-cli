package main

import (
	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandTags,
	commandSearch,
	commandList,
}

var commandTags = cli.Command{
	Name:        "tags",
	Usage:       "list/create/delete tags for a specific repository",
	Description: `...`,
	Subcommands: []cli.Command{
		{
			Name:   "list",
			Usage:  "list tags for an image",
			Action: doListTags,
			Flags: []cli.Flag{
				QuietFlag(),
			},
		},
		{
			Name:   "delete",
			Usage:  "delete tag for an image",
			Action: doDeleteTag,
		},
		{
			Name:   "create",
			Usage:  "create a tag for an image",
			Action: doCreateTag,
		},
	},
}

var commandSearch = cli.Command{
	Name:        "search",
	Usage:       "search the index",
	Description: `...`,
	Action:      doSearch,
}

var commandList = cli.Command{
	Name:        "list",
	Usage:       "list all repositories in the index",
	Description: `...`,
	Action:      doSearchAll,
}
