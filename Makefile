include ./makefiles/linux.mk
include ./makefiles/macOS.mk
include ./makefiles/windows.mk

APP = $(shell basename $(shell git remote get-url origin))
REGISTRY = us-east1-docker.pkg.dev/viktordevopscourse/k8s-k3s
VERSION=$(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)
TARGETOS = linux#linux darwin windows
TARGETARCH = arm64#amd64
TAG = ${REGISTRY}/${APP}:${VERSION}-$(TARGETOS)-${TARGETARCH}

build: format lint
	CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
	go build -v -o kbot -ldflags "-X="github.com/ViktorDevOpsCourse/kbot/config/config.Version=${VERSION}

image:
	docker buildx build --build-arg TARGETOS=$(TARGETOS) \
    	--build-arg TARGETARCH=$(TARGETARCH) \
    	--build-arg MAKEFILE_RULE=$(MAKEFILE_RULE) \
    	--tag $(TAG) .

push:
	docker push $(TAG)

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
	rm -rf kbot.exe
	docker rmi -f ${REGISTRY}/${APP}:${VERSION}-${TARGETARCH}