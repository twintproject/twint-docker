build:
	@docker build -t x0rzkov/twint:v2.1.10-alpine3.10 -f Dockerfile.alpine --no-cache .

run:
	@docker run -ti --rm x0rzkov/twint:v2.1.10-alpine3.10 -h

## help			:	Print commands help.
.PHONY: help
help : Makefile
	@sed -n 's/^##//p' $<

# https://stackoverflow.com/a/6273809/1826109
%:
	@:
