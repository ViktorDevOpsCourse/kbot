include ./mk/linux.mk
include ./mk/macOS.mk
include ./mk/windows.mk

APP = $(shell basename $(shell git remote get-url origin))
REGISTRY = us-east1-docker.pkg.dev/viktordevopscourse/k8s-k3s
VERSION=$(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)
TARGETOS = linux#linux darwin windows
TARGETARCH = arm64#amd64
TAG = ${REGISTRY}/${APP}:${VERSION}-$(TARGETOS)-${TARGETARCH}

build: format lint ## build application with defined OS and ARCH including [format] and [lint]
	CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
	go build -v -o kbot -ldflags "-X="github.com/ViktorDevOpsCourse/kbot/config/config.Version=${VERSION}

image: ## building docker image with tag and defined OS and ARCH
	docker build --build-arg TARGETOS=$(TARGETOS) \
		--build-arg TARGETARCH=$(TARGETARCH) \
		--build-arg TARGET=$(TARGETOS) \
		--tag $(TAG) .

push: ## push docker image with last tag to docker registry
	docker push $(TAG)

format: ## formatting golang code to be beautiful
	gofmt -s -w ./

lint: ## analyze code to identify potential errors, stylistic inconsistencies, and other aspects that can be
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2
	golangci-lint run -v

test: ## launch tests for code
	go test -v

.PHONY: clean
clean: ## remove old artifacts
	rm -rf kbot
	rm -rf kbot.exe
	docker rmi -f ${REGISTRY}/${APP}:${VERSION}-${TARGETARCH}

.PHONY: help
help:
	@echo "Usage: make [target]"
	@echo "Available targets:"
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-30s\033[0m %s\n", $$1, $$2}'

