package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"sort"
	"strings"

	"github.com/blang/semver"
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

func getImagesInfo(dockerRepository, vcsRepository string) {
	r, err := request(fmt.Sprintf("https://hub.docker.com/v2/repositories/%s/tags/?page_size=100", dockerRepository))
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
		log.Fatalln(err)
	}

	fmt.Println("| Image   |      Size      |  Os |  Arch |  Link |")
	fmt.Println("|----------|:-------------:|------|------|------|")

	for _, tagInfoResponse := range tagInfos {
		linkVersion := strings.Replace(tagInfoResponse.Name, "-", "/", -1)
		link := fmt.Sprintf("[`./dockerfiles/%s`](https://github.com/%s/tree/%s/dockerfiles/%s/)", linkVersion, vcsRepository, branch, linkVersion)
		fmt.Println("| docker pull", dockerRepository+"/"+tagInfoResponse.Name, "|", tagInfoResponse.HumanSize, "|", tagInfoResponse.Architecture, "|", tagInfoResponse.Os, "|", link, "|")
	}
}

func semverSorting(version string) {
	v, err := semver.Make(version)
	if err != nil {
		log.Fatalln("version:", version, "error:", err)
	}
	fmt.Printf("Major: %d\n", v.Major)
	fmt.Printf("Minor: %d\n", v.Minor)
	fmt.Printf("Patch: %d\n", v.Patch)
	fmt.Printf("Pre: %s\n", v.Pre)
	fmt.Printf("Build: %s\n", v.Build)

	// Prerelease versions array
	if len(v.Pre) > 0 {
		fmt.Println("Prerelease versions:")
		for i, pre := range v.Pre {
			fmt.Printf("%d: %q\n", i, pre)
		}
	}

	// Build meta data array
	if len(v.Build) > 0 {
		fmt.Println("Build meta data:")
		for i, build := range v.Build {
			fmt.Printf("%d: %q\n", i, build)
		}
	}

}

func tagInfoSorting(tagsInfo TagsInfoResponse) {
	versionsRaw := []string{"1.1", "0.7.1", "1.4-beta", "1.4", "2"}
	versions := make([]*version.Version, len(versionsRaw))
	for i, raw := range versionsRaw {
		v, _ := version.NewVersion(raw)
		versions[i] = v
	}
	// After this, the versions are properly sorted
	sort.Sort(version.Collection(versions))
}

func getGitConfig(config string) (string, error) {
	buf := new(bytes.Buffer)

	cmd := exec.Command("git", "config", "--get", config)
	cmd.Stdout = buf
	err := cmd.Run()

	return strings.TrimSpace(buf.String()), err
}

func getGitBranch() (string, error) {
	buf := new(bytes.Buffer)
	cmd := exec.Command("git", "branch")
	cmd.Stdout = buf
	err := cmd.Run()
	cleanBranch := strings.Replace(strings.Replace(strings.TrimSpace(buf.String()), "* ", "", -1), "\n", "", -1)
	return cleanBranch, err
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
