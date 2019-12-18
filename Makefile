IMAGE := x0rzkov/twint-docker
VERSION:= $(shell grep IBM_CLOUD_CLI Dockerfile | awk '{print $2}' | cut -d '=' -f 2)

## test		:	test.
test:
	true

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
