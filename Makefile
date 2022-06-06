.PHONY: build serve
VERSION := $(shell git describe --always |sed -e "s/^v//")
FILE_NAME=build

COLOR_YELLOW:=\033[0;33m
COLOR_GREEN:=\033[0;32m
COLOR_END:=\033[0m

build:
	@echo -e "$(COLOR_YELLOW)Compiling source"
	GOOS=linux $(GO_EXTRA_BUILD_ARGS) go build -a -tags sqlite_omit_load_extension -ldflags "-extldflags '-static' -s -w -X main.version=$(VERSION)" -o ${FILE_NAME}
	@echo -e "$(COLOR_END)"

serve: build
	$(PWD)/${FILE_NAME}
