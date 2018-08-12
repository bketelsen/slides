GIT_SUMMARY := $(shell git describe --tags --dirty --always)
REPO=bketelsen/slides
DOCKER_IMAGE := $(REPO):$(GIT_SUMMARY)
default: publish

repo:
	@echo $(DOCKER_IMAGE)

build:
	@go install github.com/bketelsen/slides
	@GOOS=linux CGO_ENABLE=0 go build -o slides main.go
	@docker build -t $(DOCKER_IMAGE) .
	@docker tag $(DOCKER_IMAGE) $(REPO)

push:
	@docker push $(DOCKER_IMAGE)
	@docker push $(REPO)

install:
	@go install github.com/bketelsen/slides

publish: install
	@slides build

release-snapshot:
	@goreleaser -snapshot

release:
	@goreleaser

release:
	@build
	@push
