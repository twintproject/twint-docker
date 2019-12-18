package main

import (
	"log"
	"os"
	"path"
	"strings"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

// Retrieve remote tags without cloning repository
func main() {
	err, tags := getRemoteTags()
	if err != nil {
		log.Fatalln(err)
	}
	var cleanTags []string
	for _, tag := range tags {
		if strings.HasPrefix(tag, "v") {
			tag = strings.Replace(tag, "v", "", -1)
		}
		cleanTags = append(cleanTags, tag)
	}
	log.Printf("Tags found: %v", cleanTags)
	createDirectories(cleanTags)
}

func createDirectories(tags []string) {
	for _, tag := range tags {
		os.MkdirAll(tag, 0755)
		os.MkdirAll(path.Join(tag, "alpine"), 0755)
	}
}

func getRemoteTags() (error, []string) {
	// Create the remote with repository URL
	rem := git.NewRemote(memory.NewStorage(), &config.RemoteConfig{
		Name: "origin",
		URLs: []string{"https://github.com/twintproject/twint"},
	})
	log.Print("Fetching tags...")
	// We can then use every Remote functions to retrieve wanted information
	refs, err := rem.List(&git.ListOptions{})
	if err != nil {
		return err, []string{}
	}
	// Filters the references list and only keeps tags
	var tags []string
	for _, ref := range refs {
		if ref.Name().IsTag() {
			tags = append(tags, ref.Name().Short())
		}
	}
	return nil, tags
}
