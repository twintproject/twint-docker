## build			:	build docker container for twitter-graph.
.PHONY: build
build:
	@docker build -t twint-graph:alpine .

## run                        :       run docker container twitter-graph.
.PHONY: run
run:
	@docker run -ti twint-graph:alpine

## help			:	Print commands help.
.PHONY: help
help : Makefile
	@sed -n 's/^##//p' $<

# https://stackoverflow.com/a/6273809/1826109
%:
	@:
