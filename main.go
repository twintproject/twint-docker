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
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

var (
	debugMode    = true
	verboseMode  = false
	silentMode   = true
	lastVersion  string
	vcsTags      []*vcsTag
	dockerImages = []string{"alpine", "ubuntu"}
)

type vcsTag struct {
	name string
	dir  string
}

// Retrieve remote tags without cloning repository
func main() {

	flag.Parse()
	projectName := flag.Arg(0)
	pp.Println("projectName", projectName)

	err, tags := getRemoteTags()
	if err != nil {
		log.Fatalln(err)
	}

	pp.Println("tags: ", tags)

	var vcsTags []*vcsTag
	for _, tag := range tags {
		dir := tag
		if strings.HasPrefix(tag, "v") {
			dir = strings.Replace(tag, "v", "", -1)
		}
		vcsTags = append(vcsTags, &vcsTag{name: tag, dir: dir})
	}

	lastVersion = getLastVersion(tags)
	log.Printf("Detected version: %v", lastVersion)
	vcsTags = append(vcsTags, &vcsTag{name: lastVersion, dir: "latest"})

	pp.Println("vcsTags: ", vcsTags)

	createDirectories(vcsTags)
	for _, dockerImage := range dockerImages {
		for _, vcsTag := range vcsTags {
			switch dockerImage {
			case "alpine":
				generateDockerfile("alpine", dockerImage+"Template", alpineTemplate, vcsTag)
			case "ubuntu":
				generateDockerfile("", dockerImage+"Template", ubuntuTemplate, vcsTag)
			}
		}
	}
}

type dockerfileData struct {
	Version string
}

// generateDockerfile("twint", "alpineTemplate", "alpineTemplate")
func generateDockerfile(prefixPath, tmplName, tmplID string, vcsTag *vcsTag) error {
	outputPath := filepath.Join("dockerfiles", vcsTag.dir, prefixPath, "Dockerfile")
	pp.Println("outputPath: ", outputPath)
	tDockerfile := template.Must(template.New(tmplName).Parse(tmplID))
	dockerfile, err := os.Create(outputPath)
	if err != nil {
		fmt.Println("Error creating the template :", err)
		return err
	}
	cfg := &dockerfileData{
		Version: vcsTag.name,
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

func createDirectories(tags []*vcsTag) {
	for _, tag := range tags {
		os.MkdirAll(path.Join("dockerfiles", tag.dir, "alpine"), 0755)
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

const (
	entrypointTemplate = `#!/bin/bash
$@`
)

const (
	alpineTemplate = `FROM alpine:3.10

MAINTAINER x0rxkov@protonmail.com

ARG TWINT_GID=997
ARG TWINT_UID=997

RUN addgroup -g 997 twint && \
    adduser -u 997 -D -h /opt/twint -s /bin/sh -G twint twint

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

RUN git clone --depth=1 -b {{.Version}} https://github.com/twintproject/twint /opt/twint \
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
pip3 install --upgrade -e git+https://github.com/twintproject/twint.git@{{.Version}}#egg=twint

RUN \
apt-get clean autoclean && \
rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

ENTRYPOINT ["/entrypoint.sh"]
VOLUME /twint
WORKDIR /srv/twint
`
)
