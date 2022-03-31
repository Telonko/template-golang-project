.PHONY: build server
VERSION := $(shell git describe --always |sed -e "s/^v//")
FILE_NAME=broker

COLOR_YELLOW:=\033[0;33m
COLOR_GREEN:=\033[0;32m
COLOR_END:=\033[0m

build:
	@echo -e "$(COLOR_YELLOW)Compiling source"
	CGO_ENABLED=0 $(GO_EXTRA_BUILD_ARGS) go build -ldflags "-s -w -X main.version=$(VERSION)" -o ${FILE_NAME}
	@echo -e "$(COLOR_END)"

serve: build
	$(PWD)/${FILE_NAME}
