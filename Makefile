APP := $(shell basename $(shell git remote get-url origin))
APP_NAME := kbot
REGISTRY := viktordevopscourse
#VERSION=$(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)
VERSION="v1.0.1"
TARGETOS:=linux #linux darwin windows
TARGETARCH:=$(shell uname -m) #amd64 arm64

build: format lint
	CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -v -o kbot -ldflags "-X="github.com/ViktorDevOpsCourse/kbot/config/config.Version=${VERSION}

macOS: TARGETOS=darwin
macOS: build

linux: TARGETOS=linux
linux: build

windows: TARGETOS=windows
windows: APP_NAME=kbot.exe
windows: TARGETARCH=$(shell echo %PROCESSOR_ARCHITECTURE%)
windows: build

docker:

#image:
#	docker build . -t ${REGISTRY}/${APP}:${VERSION}-${TARGETARCH}  --build-arg TARGETARCH=${TARGETARCH}

#push:
#	docker push ${REGISTRY}/${APP}:${VERSION}-${TARGETARCH}

format:
	gofmt -s -w ./

lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2
	golangci-lint run -v

test:
	go test -v

.PHONY: clean
clean:
	rm -rf kbot
	docker rmi ${REGISTRY}/${APP}:${VERSION}-${TARGETARCH}