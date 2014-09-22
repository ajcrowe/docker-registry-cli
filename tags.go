package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"os"
	"time"
)

// A Tag contains all the information obtained from:
// v1/repositories/<namespace>/<repo>/tags/<tag>/json
type Tag struct {
	ImageID         string
	Arch            string `json:"arch"`
	DockerGoVersion string `json:"docker_go_version"`
	DockerVersion   string `json:"docker_version"`
	Kernel          string `json:"kernel"`
	LastUpdate      int64  `json:"last_update"`
	OS              string `json:"os"`
}

// method to return RFC850 formatted time
func (t *Tag) timeRFC3339() string {
	return time.Unix(t.LastUpdate, 0).Format(time.RFC850)
}

// doListTags outputs all the tags associated with a specific repository
func doListTags(c *cli.Context) {
	repo := c.Args().First()
	if repo == "" {
		log.Println("please enter a repository.")
		return
	}

	tags := getRepoTags(repo)

	w := getTabWriter()

	if c.Bool("quiet") == false {
		writeLine(w, fmt.Sprintf("%d tags for repo %s", len(tags), repo))
		writeHeader(w, "Tag", "ImageID")
	}

	for tag, img := range tags {
		writeLine(w, tag, img)
	}
	w.Flush()
}

func getRepoTags(repo string) map[string]string {
	var tags map[string]string
	json.Unmarshal(newRequestGet(fmt.Sprintf("repositories/%s/tags", repo)), &tags)
	return tags
}

// doTagInfo output all the detailed informa
func doTagInfo(c *cli.Context) {
	repo := c.Args().First()
	tag := c.Args().Get(1)

	w := getTabWriter()

	if !checkTagExists(repo, tag) {
		writeLine(w, fmt.Sprintf("Tag: \"%s\" does not exist in repo \"%s\"", tag, repo))
		w.Flush()
		os.Exit(1)
	}

	var t Tag
	json.Unmarshal(newRequestGet(fmt.Sprintf("repositories/%s/tags/%s/json", repo, tag)), &t)
	t.ImageID = getImageIDByTag(repo, tag)
	log.Println(t.ImageID)

	if c.Bool("quiet") == false {
		writeLine(w, fmt.Sprintf("Detail for: %s:%s", repo, tag))
		writeHeader(w, "Parameter", "Value")
	}

	writeLine(w, "Image ID", t.ImageID)
	writeLine(w, "Architecture", t.Arch)
	writeLine(w, "Docker Go Version", t.DockerGoVersion)
	writeLine(w, "Docker Version", t.DockerVersion)
	writeLine(w, "Kernel", t.Kernel)
	writeLine(w, "Last Update", t.timeRFC3339())
	writeLine(w, "OS", t.OS)
	w.Flush()
}

func getImageIDByTag(repo, tag string) (id string) {
	json.Unmarshal(newRequestGet(fmt.Sprintf("repositories/%s/tags/%s", repo, tag)), &id)
	log.Println(id)
	return id
}

func checkTagExists(repo, tag string) bool {
	// get the imageid to check tag exists
	var id string
	json.Unmarshal(newRequestGet(fmt.Sprintf("repositories/%s/tags/%s", repo, tag)), &id)

	// check if imageid is returned
	if id != "" {
		return true
	} else {
		return false
	}
}

func doCreateTag(c *cli.Context) {
	repo := c.Args().Get(0)
	image := c.Args().Get(1)
	tag := c.Args().Get(2)

	w := getTabWriter()

	status := newRequestPut(fmt.Sprintf("repositories/%s/tags/%s", repo, tag), image)
	if statusOK(w, "Create Tag", status) {
		fmt.Sprintln("Successfully created tag: %s:%s for image %s", repo, tag, image)
		w.Flush()
	} else {
		w.Flush()
	}

}

func doDeleteTag(c *cli.Context) {
	//repo := c.Args().First()
	//tag := c.Args().Get(1)
}
