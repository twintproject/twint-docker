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
	"github.com/k0kubun/pp"
)

// RequestError interface
type RequestError interface {
	Error() string
}

type tagInfo struct {
	Name          string           `json:"name"`
	FullSize      uint64           `json:"full_size"`
	HumanSize     string           `json:"human_size"`
	Version       *version.Version `json:"version"`
	VersionNumber string           `json:"version-number"`
	Architecture  string           `json:"arch"`
	Os            string           `json:"os"`
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

func getImagesInfo(dockerRepository, vcsRepository string) (string, error) {
	r, err := request(fmt.Sprintf("https://hub.docker.com/v2/repositories/%s/tags/?page_size=100", dockerRepository))
	if err != nil || r.StatusCode != 200 {
		fmt.Printf(" [%s]\n", ("Failed"))
		fmt.Println("Failed to connect to docker-hub repository.")
		return "", err
	} else {
		defer r.Body.Close()
	}
	byt := readResponseBody(r)
	var tagsInfoResponse TagsInfoResponse
	if err := json.Unmarshal(byt, &tagsInfoResponse); err != nil {
		return "", err
	}
	var tagInfos []*tagInfo
	for _, t := range tagsInfoResponse.Results {
		ti := &tagInfo{
			Name:         t.Name,
			HumanSize:    humanize.Bytes(t.FullSize),
			FullSize:     t.FullSize,
			Architecture: t.Images[0].Architecture,
			Os:           t.Images[0].Os,
		}
		v, _ := version.NewVersion(t.Name)
		if v != nil {
			ti.Version = v
			ti.VersionNumber = v.String()
		}
		tagInfos = append(tagInfos, ti)
	}

	sort.Slice(tagInfos, func(i, j int) bool { return tagInfos[i].VersionNumber < tagInfos[j].VersionNumber })
	if cfg.DebugMode {
		pp.Println(tagInfos)
	}

	branch, err := getCurrentBranch(".")
	if err != nil {
		return "", err
	}

	var dockerImageTable string
	dockerImageTable = fmt.Sprint("| Image   |      Size      |  Os |  Arch |  Link |\n")
	dockerImageTable += fmt.Sprint("|----------|:-------------:|------|------|------|\n")

	for _, tagInfoResponse := range tagInfos {
		linkVersion := strings.Replace(tagInfoResponse.Name, "-", "/", -1)
		link := fmt.Sprintf("[`./dockerfiles/%s`](https://github.com/%s/tree/%s/dockerfiles/%s/)", linkVersion, vcsRepository, branch, linkVersion)
		dockerImageTable += fmt.Sprint("| docker pull ", dockerRepository+"/"+tagInfoResponse.Name, "|", tagInfoResponse.HumanSize, "|", tagInfoResponse.Architecture, "|", tagInfoResponse.Os, "|", link, "|\n")
	}

	return dockerImageTable, nil
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
