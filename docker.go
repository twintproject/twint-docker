package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/hashicorp/go-version"
)

// RequestError interface
type RequestError interface {
	Error() string
}

type TagsInfoResponse struct {
	Count    int             `json:"count"`
	Next     interface{}     `json:"next"`
	Previous interface{}     `json:"previous"`
	Results  []TagInfoResult `json:"results"`
}

type TagInfoResult struct {
	Creator             float64              `json:"creator"`
	FullSize            uint64               `json:"full_size"`
	ID                  float64              `json:"id"`
	ImageID             interface{}          `json:"image_id"`
	Images              []TagInfoResultImage `json:"images"`
	LastUpdated         string               `json:"last_updated"`
	LastUpdater         float64              `json:"last_updater"`
	LastUpdaterUsername string               `json:"last_updater_username"`
	Name                string               `json:"name"`
	Repository          float64              `json:"repository"`
	V2                  bool                 `json:"v2"`
}

type TagInfoResultImage struct {
	Architecture string      `json:"architecture"`
	Digest       string      `json:"digest"`
	Features     string      `json:"features"`
	Os           string      `json:"os"`
	OsFeatures   string      `json:"os_features"`
	OsVersion    interface{} `json:"os_version"`
	Size         float64     `json:"size"`
	Variant      interface{} `json:"variant"`
}

func getImagesInfo(repository string) {
	r, err := request(fmt.Sprintf("https://hub.docker.com/v2/repositories/%s/tags/?page_size=100", repository))
	if err != nil || r.StatusCode != 200 {
		fmt.Printf(" [%s]\n", ("Failed"))
		panic("Failed to connect to docker-hub repository.")
	} else {
		defer r.Body.Close()
	}
	byt := readResponseBody(r)
	var tagsInfoResponse TagsInfoResponse
	if err := json.Unmarshal(byt, &tagsInfoResponse); err != nil {
		panic(err)
	}

	fmt.Println("| Image   |      Size      |  Os |  Arch |  Link |")
	fmt.Println("|----------|:-------------:|------|------|------|")
	for _, tagInfoResponse := range tagsInfoResponse.Results {
		linkVersion := strings.Replace(tagInfoResponse.Name, "-", "/", -1)
		link := fmt.Sprintf("[`./dockerfiles/%s`](https://github.com/x0rzkov/twint-docker/tree/alpine/dockerfiles/%s/)", linkVersion, linkVersion)
		fmt.Println("| docker pull", repository+"/"+tagInfoResponse.Name, "|", humanize.Bytes(tagInfoResponse.FullSize), "|", tagInfoResponse.Images[0].Architecture, "|", tagInfoResponse.Images[0].Os, "|", link, "|")
	}
}

func tagInfoSorting() {
	versionsRaw := []string{"1.1", "0.7.1", "1.4-beta", "1.4", "2"}
	versions := make([]*version.Version, len(versionsRaw))
	for i, raw := range versionsRaw {
		v, _ := version.NewVersion(raw)
		versions[i] = v
	}
	// After this, the versions are properly sorted
	sort.Sort(version.Collection(versions))
}

// Request makes an HTTP request
func request(target string) (*http.Response, RequestError) {
	request, err := http.NewRequest("GET", target, nil)
	if err != nil {
		return nil, err
	}
	// request.Header.Set("User-Agent", userAgent)
	client := &http.Client{}
	return client.Do(request)
}

// ReadResponseBody reads response body and return string
func readResponseBody(response *http.Response) []byte {
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	return bodyBytes
}

// hasElement reports whether elements is within array.
func hasElement(array []string, targets ...string) (bool, int) {
	for index, item := range array {
		for _, target := range targets {
			if item == target {
				return true, index
			}
		}
	}
	return false, -1
}
