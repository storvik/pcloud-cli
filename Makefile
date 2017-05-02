APPNAME = pcloud-cli
PACKAGE = github.com/storvik/${APPNAME}

COMMIT_HASH = `git rev-parse --short HEAD 2>/dev/null`
BUILD_TIME=`date +%FT%T%z`

LDFLAGS=-ldflags "-w \
    -X main.BuildTime=${BUILD_TIME} \
    -X main.CommitHash=${COMMIT_HASH}"

LDFLAGS_WITHOUT_GIT=-ldflags "-w \
    -X main.BuildTime=${BUILD_TIME} \
    -X main.CommitHash=dev"


.PHONY: all clean vendor fmt lint help
.DEFAULT_GOAL := help

all: build

clean:
	go clean
	rm ./${APPNAME} || true
	rm -rf ./release || true

vendor: ## Install govendor and sync dependencies
	go get github.com/kardianos/govendor
	govendor sync ${PACKAGE}

build: vendor ## Build pcloud-cli binary
	go build -o ${APPNAME} ${LDFLAGS}

install: vendor ## Install pcloud-cli binary
	go install ${LDFLAGS} ${PACKAGE}

build-without-git: LDFLAGS = ${NOGI_LDFLAGS}
build-without-git: vendor hugo ## Build pcloud-cli without commit hash from git

fmt: ## Run gofmt linter
	@for d in `govendor list -no-status +local | sed 's/github.com.storvik.pcloud-cli/./'` ; do \
		if [ "`gofmt -l $$d/*.go | tee /dev/stderr`" ]; then \
			echo "^ improperly formatted go files" && echo && exit 1; \
		fi \
	done

lint: ## Run golint linter
	@for d in `govendor list -no-status +local | sed 's/github.com.storvik.pcloud-cli/./'` ; do \
		if [ "`golint $$d | tee /dev/stderr`" ]; then \
			echo "^ golint errors!" && echo && exit 1; \
		fi \
	done

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

check-vendor: ## Verify that vendored packages match git HEAD
	@git diff-index --quiet HEAD vendor/ || (echo "check-vendor target failed: vendored packages out of sync" && echo && git diff vendor/ && exit 1)

help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
