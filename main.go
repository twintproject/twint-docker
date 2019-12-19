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
	"github.com/jinzhu/configor"
	"github.com/k0kubun/pp"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

/*
	Refs:
	- https://github.com/ozankasikci/dockerfile-generator
	- https://github.com/jinzhu/configor
	- https://github.com/hawx/ggg/blob/master/repos/repo.go (markdown)
	- https://github.com/zet4/go-travis-docker-test/blob/master/.travis.yml
*/

var (
	debugMode      = false
	verboseMode    = false
	silentMode     = true
	dockerImages   = []string{"alpine", "ubuntu", "slim"}
	excludeVersion = []string{"v1.0", "1.1", "1.1.2.1", "1.1.3.1", "1.1.3", "1.1.2"}
	vcsTags        []*vcsTag
	lastVersion    string
	cfg            *Config
)

type Config struct {
	APPName string           `json:"app-name" yaml:"app-name"`
	Images  map[string]Image `json:"images" yaml:"images"`
}

type Image struct {
	Name string `json:"name" yaml:"name"`
}

type vcsTag struct {
	Name string
	Dir  string
}

func isValidVersion(input string) bool {
	for _, version := range excludeVersion {
		if version == input {
			return false
		}
	}
	return true
}

func main() {
	cfg := &Config{}
	config := flag.String("config", "x0rzkov.yml", "configuration file")
	// debugMode = flag.Bool("debug", false, "debug mode")
	flag.StringVar(&cfg.APPName, "name", "", "app name")
	flag.Parse()

	// os.Setenv("CONFIGOR_ENV_PREFIX", "-")
	loadConfig(*config)

	err, tags := getRemoteTags()
	if err != nil {
		log.Fatalln(err)
	}

	if debugMode {
		pp.Println("tags: ", tags)
	}
	var vcsTags []*vcsTag
	for _, tag := range tags {
		dir := tag
		if strings.HasPrefix(tag, "v") {
			dir = strings.Replace(tag, "v", "", -1)
		}
		if isValidVersion(tag) {
			vcsTags = append(vcsTags, &vcsTag{Name: tag, Dir: dir})
		}
	}

	lastVersion = getLastVersion(tags)
	log.Printf("Latest version: %v", lastVersion)
	vcsTags = append(vcsTags, &vcsTag{Name: "v" + lastVersion, Dir: "latest"})

	if debugMode {
		pp.Println("vcsTags: ", vcsTags)
	}
	removeContents("./dockerfiles")
	createDirectories(vcsTags)
	for _, dockerImage := range dockerImages {
		for _, vcsTag := range vcsTags {
			switch dockerImage {
			case "slim":
				generateDockerfile("slim", dockerImage+"Template", debianSlimTemplate, vcsTag)
				generateEntrypoint("slim", "entrypointTemplate", entrypointTemplate, vcsTag)
				generateMakefile("slim", "makefileTemplate", makefileTemplate, vcsTag)
				generateDockerignore("slim", "dockerignoreTemplate", dockerignoreTemplate, vcsTag)
			case "alpine":
				generateDockerfile("alpine", dockerImage+"Template", alpineTemplate, vcsTag)
				generateEntrypoint("alpine", "entrypointTemplate", entrypointTemplate, vcsTag)
				generateMakefile("alpine", "makefileTemplate", makefileTemplate, vcsTag)
				generateDockerignore("alpine", "dockerignoreTemplate", dockerignoreTemplate, vcsTag)
			case "ubuntu":
				generateDockerfile("", dockerImage+"Template", ubuntuTemplate, vcsTag)
				generateEntrypoint("", "entrypointTemplate", entrypointTemplate, vcsTag)
				generateMakefile("", "makefileTemplate", makefileTemplate, vcsTag)
				generateDockerignore("", "dockerignoreTemplate", dockerignoreTemplate, vcsTag)
			}
		}
	}
	generateTravis(vcsTags)
}

func loadConfig(path ...string) (error, *Config) {
	err := configor.New(&configor.Config{
		Debug:                true,
		Verbose:              true,
		AutoReload:           true,
		ErrorOnUnmatchedKeys: false,
		AutoReloadInterval:   time.Minute,
		AutoReloadCallback: func(config interface{}) {
			fmt.Printf("%v changed", config)
		},
	}).Load(cfg, path...)
	return err, cfg
}

type dockerfileData struct {
	Version string
	Dir     string
}

func generateDockerfile(prefixPath, tmplName, tmplID string, vcsTag *vcsTag) error {
	outputPath := filepath.Join("dockerfiles", vcsTag.Dir, prefixPath, "Dockerfile")
	if debugMode {
		pp.Println("outputPath: ", outputPath)
	}
	tDockerfile := template.Must(template.New(tmplName).Parse(tmplID))
	dockerfile, err := os.Create(outputPath)
	if err != nil {
		fmt.Println("Error creating the template :", err)
		return err
	}
	cfg := &dockerfileData{
		Version: vcsTag.Name,
		Dir:     vcsTag.Dir,
	}
	err = tDockerfile.Execute(dockerfile, cfg)
	if err != nil {
		fmt.Println("Error creating the template :", err)
		return err
	}
	return nil
}

type travisData struct {
	Versions []*vcsTag
}

func generateTravis(vcsTag []*vcsTag) error {
	tTravisfile := template.Must(template.New("tmplTravis").Parse(travisTemplate))
	travisfile, err := os.Create(".travis.yml")
	if err != nil {
		if debugMode {
			fmt.Println("Error creating the template :", err)
		}
		return err
	}
	cfg := &travisData{
		Versions: vcsTag,
	}
	err = tTravisfile.Execute(travisfile, cfg)
	if err != nil {
		if debugMode {
			fmt.Println("Error creating the template :", err)
		}
		return err
	}
	return nil
}

type entrypointData struct {
}

func generateEntrypoint(prefixPath, tmplName, tmplID string, vcsTag *vcsTag) error {
	tEntrypoint := template.Must(template.New("tmplEntrypoint").Parse(entrypointTemplate))
	outputPathEntrypoint := filepath.Join("dockerfiles", vcsTag.Dir, prefixPath, "docker-entrypoint.sh")
	if debugMode {
		pp.Println("outputPathEntrypoint: ", outputPathEntrypoint)
	}
	entrypoint, err := os.Create(outputPathEntrypoint)
	if err != nil {
		fmt.Println("Error creating the template :", err)
		return err
	}
	cfg := &entrypointData{}
	err = tEntrypoint.Execute(entrypoint, cfg)
	if err != nil {
		fmt.Println("Error creating the template :", err)
		return err
	}
	err = os.Chmod(outputPathEntrypoint, 0755)
	if err != nil {
		return err
	}
	return nil
}

type makefileData struct {
}

func generateMakefile(prefixPath, tmplName, tmplID string, vcsTag *vcsTag) error {
	tMakefile := template.Must(template.New("tmplMakefile").Parse(makefileTemplate))
	outputPathMakefile := filepath.Join("dockerfiles", vcsTag.Dir, prefixPath, "Makefile")
	if debugMode {
		pp.Println("outputPathMakefile: ", outputPathMakefile)
	}
	makefile, err := os.Create(outputPathMakefile)
	if err != nil {
		fmt.Println("Error creating the template :", err)
		return err
	}
	cfg := &makefileData{}
	err = tMakefile.Execute(makefile, cfg)
	if err != nil {
		fmt.Println("Error creating the template :", err)
		return err
	}
	return nil
}

type dockerignoreData struct {
}

func generateDockerignore(prefixPath, tmplName, tmplID string, vcsTag *vcsTag) error {
	tDockerIgnore := template.Must(template.New("tmplDockerIgnore").Parse(dockerignoreTemplate))
	outputPath := filepath.Join("dockerfiles", vcsTag.Dir, prefixPath, ".dockerignore")
	if debugMode {
		pp.Println("outputPath: ", outputPath)
	}
	makefile, err := os.Create(outputPath)
	if err != nil {
		fmt.Println("Error creating the template :", err)
		return err
	}
	cfg := &dockerignoreData{}
	err = tDockerIgnore.Execute(makefile, cfg)
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

// https://github.com/chilic/docker-hugo/blob/master/cmd/build.go
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
		for _, image := range dockerImages {
			if image != "ubuntu" {
				os.MkdirAll(path.Join("dockerfiles", tag.Dir, image), 0755)
			}
		}
	}
}

func removeContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
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
	alpineTemplate = `FROM alpine:3.10 AS build

WORKDIR /opt/app

# Install Python and external dependencies, including headers and GCC
RUN apk add --no-cache python3 python3-dev py3-pip libffi libffi-dev musl-dev gcc git ca-certificates openblas-dev musl-dev g++

# Install Pipenv
RUN pip3 install pipenv

# Create a virtual environment and activate it
RUN python3 -m venv /opt/venv
ENV PATH="/opt/venv/bin:$PATH" \
	VIRTUAL_ENV="/opt/venv"

# Install dependencies into the virtual environment with Pipenv
RUN git clone --depth=1 -b {{.Version}} https://github.com/twintproject/twint /opt/app \
	&& cd /opt/app \
	&& pip3 install --upgrade pip \
	&& pip3 install cython \
	&& pip3 install numpy \
	&& pip3 install .

FROM alpine:3.10
MAINTAINER x0rxkov <x0rxkov@protonmail.com>

WORKDIR /opt/app

# Install Python and external runtime dependencies only
RUN apk add --no-cache python3 libffi openblas libstdc++

# Copy the virtual environment from the previous image
COPY --from=build /opt/venv /opt/venv

# Activate the virtual environment
ENV PATH="/opt/venv/bin:$PATH" \
	VIRTUAL_ENV="/opt/venv"

# Copy your application
WORKDIR /opt/app

ENTRYPOINT ["twint"]`
)

const (
	debianSlimTemplate = `FROM python:3.7-slim-stretch

MAINTAINER x0rxkov <x0rxkov@protonmail.com>

ARG TWINT_VERSION={{.Version}}

COPY docker-entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

RUN \
apt-get update && \
apt-get install -y \
git

RUN \
pip3 install --upgrade -e git+https://github.com/twintproject/twint.git@{{.Version}}#egg=twint

RUN \
apt-get clean autoclean && \
rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

ENTRYPOINT ["/entrypoint.sh"]
VOLUME /twint
WORKDIR /srv/twint`
)

const (
	ubuntuTemplate = `FROM ubuntu:19.10

MAINTAINER Sébastien Houzet (yoozio.com) <sebastien@yoozio.com>

ARG TWINT_VERSION={{.Version}}

COPY docker-entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

RUN \
apt-get update && \
apt-get install -y \
git \
python3-pip

RUN \
pip3 install --upgrade -e git+https://github.com/twintproject/twint.git@{{.Version}}#egg=twint

RUN \
apt-get clean autoclean && \
rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

ENTRYPOINT ["/entrypoint.sh"]
VOLUME /twint
WORKDIR /srv/twint`
)

const (
	travisTemplate = `after_script:
  - docker images

sudo: required

before_install:
  - sudo rm -f /usr/local/bin/docker-slim
  - sudo rm -f /usr/local/bin/docker-slim-sensor
  - curl -L https://github.com/docker-slim/docker-slim/releases/download/1.26.1/dist_linux.tar.gz --output docker-slim.tar.gz
  - tar xvf docker-slim.tar.gz
  - chmod +x dist_linux/docker-slim
  - chmod +x dist_linux/docker-slim-sensor
  - sudo mv dist_linux/docker-slim /usr/local/bin
  - sudo mv dist_linux/docker-slim-sensor /usr/local/bin

before_script:
  - cd dockerfiles/"$VERSION"
  - IMAGE="x0rzkov/twint:${VERSION/\//-}"

env:{{range $val := .Versions}}
  - VERSION={{ $val.Dir }}
  - VERSION={{ $val.Dir }}/alpine
  - VERSION={{ $val.Dir }}/slim{{end}}

language: bash

script:
  - docker-slim version
  - docker build -t "$IMAGE" .
  - sudo docker-slim build "$IMAGE"
  - docker images
  - docker login -u $DOCKER_USERNAME -p $DOCKER_PASSWORD
  - docker push "$IMAGE"

services: docker`
)

const (
	dockerComposeTemplate = `---
version: '3.7'
services:

  twint:
    build:
      context: dockerfiles/latest/alpine
      dockerfile: Dockerfile
    container_name: twint
  
  elasticsearch:
    image: docker.elastic.co/elasticsearch/elasticsearch:${KIBANA_VERSION}
    container_name: twint_elastic
    environment:
    - node.name=elasticsearch
    - cluster.initial_master_nodes=elasticsearch
    - cluster.name=docker-cluster
    - bootstrap.memory_lock=true
    - "ES_JAVA_OPTS=${ELASTIC_JAVA_OPTS}"
    ulimits:
      memlock:
        soft: -1
        hard: -1
    volumes:
    - esdata01:/usr/share/elasticsearch/data
    ports:
    - 9200:9200

  kibana:
    image: docker.elastic.co/kibana/kibana:${KIBANA_VERSION}
    container_name: twint_kibana
    ports:
    - 5601:5601

volumes:
  esdata01:
    driver: local

networks:
  default:
    external:
      name: nw_twint`
)

const (
	dockerignoreTemplate = `slim
alpine
Makefile
docker-compose.yml
docker-compose.*.yml
.git
.git/
.git/*
.git/**
`
)

const (
	makefileTemplate = `IMAGE := x0rzkov/twint-docker
VERSION:= $(shell grep TWINT_VERSION Dockerfile | awk '{print $2}' | cut -d '=' -f 2)

## test		:	test.
test:
	true

## version	:	display version.
version:
	@echo $(VERSION)

## image		:	build image and tag them.
.PHONY: image
image:
	@docker build -t ${IMAGE}:${VERSION} .
	@docker tag ${IMAGE}:${VERSION} ${IMAGE}:latest

## push-image	:	push docker image.
.PHONY: push-image
push-image:
	@docker push ${IMAGE}:${VERSION}
	@docker push ${IMAGE}:latest

## help		:	Print commands help.
.PHONY: help
help : Makefile
	@sed -n 's/^##//p' $<

# https://stackoverflow.com/a/6273809/1826109
%:
	@:
`
)
