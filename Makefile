GIT_SUMMARY := $(shell git describe --tags --dirty --always)
REPO=bketelsen/slides
DOCKER_IMAGE := $(REPO):$(GIT_SUMMARY)
directory = ./statik
dir_target = $(directory)-$(wildcard $(directory))
dir_present = $(directory)-$(directory)
dir_absent = $(directory)-

default: install

.PHONY: repo
repo:
	@echo $(DOCKER_IMAGE)

.PHONY: build
build: deps install
	@GOOS=linux CGO_ENABLE=0 go build -o slides main.go
	@docker build -t $(DOCKER_IMAGE) .
	@docker tag $(DOCKER_IMAGE) $(REPO)

.PHONY: push
push:
	@docker push $(DOCKER_IMAGE)
	@docker push $(REPO)

.PHONY: clean
clean:
	@rm -rf dist/

.PHONY: install
install: deps clean statik
	@go install

.PHONY: release-snapshot
release-snapshot: clean
	@goreleaser --snapshot

.PHONY: release
release: clean
	@github-release-notes -org bketelsen -repo slides -include-commits > .releasenotes
	@goreleaser --release-notes=.releasenotes

.PHONY: release-notes
release-notes:
	@github-release-notes -org bketelsen -repo slides -include-commits


.PHONY: statik
statik: | $(dir_target)
	@echo "✔ code generated"

$(dir_present):
	@echo "✔ statik package exists"

$(dir_absent): deps
	@echo "❓ generating statik package"
	@statik -src=slides-template


.PHONY: clean-statik
clean-statik:
	@rm -rf statik/

.PHONY: deps
deps:
	@echo "❓ Checking for statik"
	@if ! [ -x "$$(command -v statik)" ]; then\
		echo "🤵 Getting statik";\
		go get github.com/rakyll/statik;\
		statik -help;\
	fi
	@echo "✔ statik binary installed"


