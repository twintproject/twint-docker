package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
	"time"

	"github.com/hashicorp/go-version"
	"github.com/k0kubun/pp"
	"github.com/wolfeidau/envfile"
	cfg "github.com/x0rzkov/twint-docker/pkg/configor"
	dfg "github.com/x0rzkov/twint-docker/pkg/generator"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

/*
	Refs:
	- https://github.com/dahernan/godockerize/blob/master/godockerize.go
	- https://github.com/ozankasikci/dockerfile-generator
*/

var (
	debugMode   = true
	verboseMode = false
	silentMode  = true
	lastVersion string
	vcsTags     []string
	genCfg      Config
)

// Retrieve remote tags without cloning repository
func main() {

	flag.Parse()
	projectName := flag.Arg(0)
	pp.Println("projectName", projectName)

	err, tags := getRemoteTags()
	if err != nil {
		log.Fatalln(err)
	}
	for _, tag := range tags {
		if strings.HasPrefix(tag, "v") {
			tag = strings.Replace(tag, "v", "", -1)
		}
		vcsTags = append(vcsTags, tag)
	}
	log.Printf("Tags found: %v", vcsTags)
	err, envMap := readEnvFile(".env")
	if err != nil {
		log.Fatalln(err)
	}
	if debugMode {
		pp.Println(envMap)
	}

	lastVersion = getLastVersion(tags)
	log.Printf("Detected version: %v", lastVersion)

	loadConfig("config.yaml")
	createDirectories(vcsTags)
	// genFromYAML("stages.yml", "", "")
	for _, vcsTag := range vcsTags {
		generateDockerfile("twint", "alpineTemplate", alpineTemplate, vcsTag, genCfg)
	}
}

// genFromYAML("stages.yml", "", "")
func genFromYAML(path, prefixPath, output string) error {
	data, err := dfg.NewDockerFileDataFromYamlFile(path)
	tmpl := dfg.NewDockerfileTemplate(data)
	// write to a file
	// file, err := os.Create("Dockerfile")
	// err = tmpl.Render(file)
	// or write to stdout
	err = tmpl.Render(os.Stdout)

	return err
}

func checkIfExists(projectName string) error {
	if _, err := os.Stat(projectName); os.IsNotExist(err) {
		os.Mkdir(filepath.Join(".", projectName), 0770)
	} else {
		return err
	}
	return nil
}

type DockerfileData struct {
	Name        string
	BaseImage   string
	BaseVersion string
}

// generateDockerfile("twint", "alpineTemplate", "alpineTemplate")
func generateDockerfile(projectName, tmplName, tmplID string, cfg Config) error {
	fmt.Printf("Generating a new dockerized application : %s\n", projectName)

	tDockerfile := template.Must(template.New(tmplName).Parse(tmplID))
	dockerfile, err := os.Create(filepath.Join(projectName, "Dockerfile"))
	if err != nil {
		fmt.Println("Error creating the template :", err)
		return err
	}

	err = tDockerfile.Execute(dockerfile, cfg)
	if err != nil {
		fmt.Println("Error creating the template :", err)
		return err
	}
	return nil
}

func getLastVersion(tags []string) string {
	versions := make([]*version.Version, len(tags))
	for i, raw := range tags {
		v, _ := version.NewVersion(raw)
		versions[i] = v
	}
	// After this, the versions are properly sorted
	sort.Sort(version.Collection(versions))
	return versions[len(versions)-1].String()
}

func commitLocal(version string) {
	r, _ := git.PlainOpen("./")
	w, _ := r.Worktree()
	status, _ := w.Status()
	if status.File("Dockerfile").Worktree == git.Modified {
		_, _ = w.Add("Dockerfile")
		_, _ = w.Commit(version, &git.CommitOptions{
			Author: &object.Signature{
				Name:  "x0rzkov",
				Email: "x0rzkov@protonmail.com",
				When:  time.Now(),
			},
		})
		_ = r.Push(&git.PushOptions{})
	}
}

func loadConfig(files ...string) {
	configorSettings := &cfg.Config{
		Debug:   debugMode,
		Verbose: verboseMode,
		Silent:  silentMode,
	}
	cfg.New(configorSettings).Load(&genCfg, files...)
	genCfg.Vcs.Version = lastVersion
	genCfg.Vcs.Tags = vcsTags
	pp.Println("config:", genCfg)
}

func readEnvFile(path string) (error, map[string]string) {
	envMap := make(map[string]string)
	err := envfile.ReadEnvFile(path, envMap)
	return err, envMap
}

func createDirectories(tags []string) {
	for _, tag := range tags {
		os.MkdirAll(path.Join("dockerfiles", tag, "alpine"), 0755)
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

type Docker struct {
	Images []Image `required:"true" json:"images" yaml:"images"`
}

type Image struct {
	Disable    bool        `default:"false" json:"disable" yaml:"disable"`
	Owner      string      `required:"true" default:"x0rzkov" json:"owner" yaml:"owner" env:"DOCKER_OWNER"`
	Image      string      `required:"true" default:"alpine" json:"image" yaml:"image" env:"DOCKER_BASE_IMAGE"`
	BuildArgs  BuildArgs   `required:"true" json:"arguments" yaml:"arguments"`
	Dockerfile []dfg.Stage `required:"true" json:"dockerfile" yaml:"dockerfile"`
}

type BuildArgs struct {
	BaseImage  string `required:"true" default:"alpine" json:"base" yaml:"base" env:"DOCKER_BASE_IMAGE"`
	BaseTag    string `required:"true" default:"latest" json:"tag" yaml:"tag" env:"DOCKER_BASE_VERSION"`
	Maintainer string `default:"x0rzkov@protonmail.com" json:"maintainer" yaml:"maintainer" env:"DOCKER_MAINTAINER"`
	UserGID    string `default:"1000" json:"user-gid" yaml:"user-gid" env:"DOCKER_USER_GID"`
	UserUID    string `default:"1000" json:"user-uid" yaml:"user-uid" env:"DOCKER_USER_UID"`
}

type VCS struct {
	RemoteURL string   `required:"true" json:"url" yaml:"url" env:"VCS_REMOTE_URL"`
	Version   string   `default:"master" json:"version" yaml:"version" env:"VCS_VERSION"`
	Tags      []string `json:"tags" yaml:"tags" env:"VCS_TAGS"`
}

type Config struct {
	Docker Docker `json:"docker" yaml:"docker"`
	Vcs    VCS    `json:"vcs" yaml:"vcs"`
}

const (
	entrypointTemplate = `#!/bin/bash
$@`
)

const (
	alpineTemplate = `FROM alpine:{{.BaseVersion}}

MAINTAINER {{.Maintainer}}

ARG TWINT_GID={{.TwintGID}}
ARG TWINT_UID={{.TwintUID}}

RUN addgroup -g ${TWINT_GID} twint && \
    adduser -u ${TWINT_UID} -D -h /opt/twint -s /bin/sh -G twint twint

# This hack is widely applied to avoid python printing issues in docker containers.
# See: https://github.com/Docker-Hub-frolvlad/docker-alpine-python3/pull/13
ENV PYTHONUNBUFFERED=1

RUN echo "**** install Python ****" && \
    apk add --no-cache python3 sqlite sqlite-dev git ca-certificates cython openblas-dev musl-dev python3-dev libffi-dev gcc g++ && \
    if [ ! -e /usr/bin/python ]; then ln -sf python3 /usr/bin/python ; fi && \
    \
    echo "**** install pip ****" && \
    python3 -m ensurepip && \
    rm -r /usr/lib/python*/ensurepip && \
    pip3 install --no-cache --upgrade pip setuptools wheel && \
    if [ ! -e /usr/bin/pip ]; then ln -s pip3 /usr/bin/pip ; fi

WORKDIR /opt/twint

RUN git clone --depth=1 -b {{.TwintVersion}} https://github.com/twintproject/twint /opt/twint \
	&& cd /opt/twint \
	&& pip install -e .

WORKDIR /opt/twint

ENTRYPOINT ["twint"]
`
)

const (
	ubuntuTemplate = `FROM ubuntu:18.04

MAINTAINER Sébastien Houzet (yoozio.com) <sebastien@yoozio.com>

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

RUN \
apt-get update && \
apt-get install -y --no-install-recommends \
git \
python3-pip

RUN \
pip3 install --upgrade -e git+https://github.com/twintproject/twint.git@v2.1.10#egg=twint

RUN \
apt-get clean autoclean && \
rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

ENTRYPOINT ["/entrypoint.sh"]
VOLUME /twint
WORKDIR /srv/twint
`
)