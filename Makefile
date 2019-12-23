IMAGE := x0rzkov/twint-docker-generator
# VERSION:= $(shell grep TWINT_GENERATOR_VERSION Dockerfile.generator | awk '{print $2}' | cut -d '=' -f 2)

VERSION := $(shell git describe HEAD)
BRANCH := $(shell git rev-parse --abbrev-ref HEAD | tr / -)
NOW=$(shell TZ=UTC date +%Y-%m-%dT%H:%M:%SZ)

## test		:	test.
test:
	true

## run		:	run generator (requires golang to be already installed).
.PHONY: run
run: deps
	@go-bindata .docker/templates/...
	@go run *.go

## build		:	build generator (requires golang to be already installed).
.PHONY: build
build: deps
	@go-bindata .docker/templates/...
	@go build -v

## deps		:	install dependencies.
.PHONY: deps
deps:
	@go get -u github.com/go-bindata/go-bindata/...

## image		:	build image and tag them.
.PHONY: image
image:
	@docker build --build-arg NOW=$(NOW) --build-arg VERSION=$(VERSION) -t "$(IMAGE):$(VERSION)" -f Dockerfile.generator .
	@docker tag $(IMAGE):$(VERSION) $(IMAGE):latest

## generate	:	generate dockerfiles and all other templates (travis-ci, makefile,...).
.PHONY: generate
generate:
	@rm -fR ./.travis.yml
	@rm -fR ./dockerfiles
	@docker run -ti -v $(PWD):/opt/twint-docker/data "$(IMAGE):$(VERSION)"

## push-image	:	push docker image.
.PHONY: push-image
push-image:
	@docker push $(IMAGE):$(VERSION)
	@docker push $(IMAGE):latest

## help		:	Print commands help.
.PHONY: help
help : Makefile
	@sed -n 's/^##//p' $<

# https://stackoverflow.com/a/6273809/1826109
%:
	@:
