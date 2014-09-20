package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"time"
)

type Tag struct {
	Arch            string `json:"arch"`
	DockerGoVersion string `json:"docker_go_version"`
	DockerVersion   string `json:"docker_version"`
	Kernel          string `json:"kernel"`
	LastUpdate      int64  `json:"last_update"`
	OS              string `json:"os"`
}

func (t *Tag) timeRFC3339() string {
	return time.Unix(t.LastUpdate, 0).Format(time.RFC850)
}

func doListTags(c *cli.Context) {
	repo := c.Args().First()
	if repo == "" {
		log.Println("please enter a repository.")
		return
	}

	var tags map[string]string
	json.Unmarshal(NewRequestGet(fmt.Sprintf("repositories/%s/tags", repo)), &tags)

	w := getTabWriter()

	if c.Bool("quiet") == false {
		writeLine(w, fmt.Sprintf("%d tags for repo %s", len(tags), repo))
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

	var t Tag
	json.Unmarshal(NewRequestGet(fmt.Sprintf("repositories/%s/tags/%s/json", repo, tag)), &t)

	w := getTabWriter()

	if c.Bool("quiet") == false {
		writeLine(w, fmt.Sprintf("Detail for: %s/%s", repo, tag))
		writeHeader(w, "Parameter", "Value")
	}

	writeLine(w, "Architecture", t.Arch)
	writeLine(w, "Docker Go Version", t.DockerGoVersion)
	writeLine(w, "Docker Version", t.DockerVersion)
	writeLine(w, "Kernel", t.Kernel)
	writeLine(w, "Last Update", t.timeRFC3339())
	writeLine(w, "OS", t.OS)
	w.Flush()
}

func doCreateTag(c *cli.Context) {
}

func doDeleteTag(c *cli.Context) {
	//repo := c.Args().First()
	//tag := c.Args().Get(1)
}
