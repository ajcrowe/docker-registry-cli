package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"strconv"
	//"strings"
)

type tag struct {
	//Name    string
	ImageID string
}

func doListTags(c *cli.Context) {
	repo := c.Args().First()
	if repo == "" {
		log.Println("please enter a repository.")
		return
	}

	JSONResults := NewRequestGet(fmt.Sprintf("repositories/%s/tags", repo))

	var tags map[string]string
	json.Unmarshal(JSONResults, &tags)

	w := getTabWriter()

	if c.Bool("quiet") == false {
		writeLine(w, "found: "+strconv.Itoa(len(tags))+" tags for repo "+repo)
		writeHeader(w, "Tag", "ImageID")
	}

	for tag, img := range tags {
		writeLine(w, tag, img[0:11])
	}
	w.Flush()
}

func doTagInfo(c *cli.Context) {
	repo := c.Args().First()
	tag := c.Args().Get(1)

	var image string
	json.Unmarshal(NewRequestGet(fmt.Sprintf("repositories/%s/%s", repo, tag)), &image)

	//JSONResults := NewRequestGet(fmt.Sprintf("images/%s/json", image))

}

func doCreateTag(c *cli.Context) {
}

func doDeleteTag(c *cli.Context) {
	//repo := c.Args().First()
	//tag := c.Args().Get(1)
}
