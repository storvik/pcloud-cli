APPNAME = pcloud-cli
PACKAGE = github.com/storvik/${APPNAME}

VERSION_TAG=0.9.5
COMMIT_HASH = `git rev-parse --short HEAD 2>/dev/null`
BUILD_TIME=`date +%FT%T%z`

LDFLAGS=-ldflags "-w \
    -X main.BUILD_TIME=${BUILD_TIME} \
    -X main.COMMIT_HASH=${COMMIT_HASH} \
    -X main.VERSION=${VERSION_TAG}"

.PHONY: vendor docker check fmt lint test test-race vet test-cover-html help
.DEFAULT_GOAL := help

all: build

clean:
	go clean
	rm ./${APPNAME} || true
	rm -rf ./release || true

build:
	go build -o ${APPNAME} ${LDFLAGS}

vendor: ## Install govendor and sync dependencies
	go get github.com/kardianos/govendor
	govendor sync ${PACKAGE}

release: clean darwin linux ## Build pcloud-cli for Os X and Linux

darwin: ## Build pcloud-cli for Mac Os X
	GOOS=darwin GOARCH=386 go build ${LDFLAGS} -o ./release/${APPNAME}_darwin_386
	GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o ./release/${APPNAME}_darwin_amd64

linux: ## Build pcloud-cli for Linux
	GOOS=linux GOARCH=386 go build ${LDFLAGS} -o ./release/${APPNAME}_linux_386
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ./release/${APPNAME}_linux_amd64

windows: ## Build pcloud-cli for Windows
	GOOS=windows GOARCH=386 go build ${LDFLAGS} -o ./release/${APPNAME}_windows_386.exe
	GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o ./release/${APPNAME}_windows_amd64.exe

help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
