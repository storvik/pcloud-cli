APPNAME=pcloud-cli

VERSION_TAG=0.9.5
BUILD=`git rev-parse HEAD`
BUILD_TIME=`date +%FT%T%z`

LDFLAGS=-ldflags "-w \
    -X main.BUILD_TIME=${BUILD_TIME} \
    -X main.BUILD=${BUILD} \
    -X main.VERSION=${VERSION_TAG}"

all: build

clean:
	go clean
	rm ./${APPNAME} || true
	rm -rf ./release || true

build:
	go build -o ${APPNAME} ${LDFLAGS}

release: clean darwin linux

android:
	GOOS=android GOARCH=arm64 go build ${LDFLAGS} -o ./release/${APPNAME}_android

darwin:
	GOOS=darwin GOARCH=386 go build ${LDFLAGS} -o ./release/${APPNAME}_darwin_386
	GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o ./release/${APPNAME}_darwin_amd64

linux:
	GOOS=linux GOARCH=386 go build ${LDFLAGS} -o ./release/${APPNAME}_linux_386
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ./release/${APPNAME}_linux_amd64

windows:
	GOOS=windows GOARCH=386 go build ${LDFLAGS} -o ./release/${APPNAME}_windows_386.exe
	GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o ./release/${APPNAME}_windows_amd64.exe
