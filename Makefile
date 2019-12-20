IMAGE := x0rzkov/twint-docker-generator
VERSION:= $(shell grep TWINT_GENERATOR Dockerfile.generator | awk '{print $2}' | cut -d '=' -f 2)

## test		:	test.
test:
	true

## build		:	build generator.
.PHONY: build
build:
	@go-bindata .docker/templates/...

## deps		:	install dependencies.
.PHONY: deps
deps:
	@go get -u github.com/go-bindata/go-bindata/...

## image		:	build image and tag them.
.PHONY: image
image:
	@docker build -t "$(IMAGE):$(VERSION)" -f Dockerfile.generator .
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
