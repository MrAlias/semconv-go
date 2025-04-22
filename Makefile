.DEFAULT_GOAL := generate

# The weaver docker image to use for semconv-generate.
WEAVER_IMAGE := $(shell awk '$$4=="weaver" {print $$2}' ./dependencies.Dockerfile)

DOCKER_USER=$(shell id -u):$(shell id -g)

# The output of "git ls-remote" looks something like this:
#
#    e531541025992b68177a68b87628c5dc75c4f7d9        refs/tags/v1.21.0
#    cadfe53949266d33476b15ca52c92f682600a29c        refs/tags/v1.22.0
#    ...
#
# .. which is why some additional processing is required to extract the
# latest version number and strip off the "v" prefix.
ifeq ($(TAG),)
TAG := $( git ls-remote --tags git@github.com:open-telemetry/semantic-conventions.git | cut -f 2 | sort --reverse | head -n 1 | cut -d '/' -f3 )
endif

.PHONY: generate
generate:
	@mkdir -p $(PWD)/$(TAG)
	@mkdir -p ~/.weaver
	@docker run --rm \
		-u $(DOCKER_USER) \
		--mount 'type=bind,source=$(PWD)/templates,target=/home/weaver/templates,readonly' \
		--mount 'type=bind,source=$(PWD)/$(TAG),target=/home/weaver/target' \
		--mount 'type=bind,source=$(HOME)/.weaver,target=/tmp/weaver/.weaver' \
		${WEAVER_IMAGE} registry generate \
		--registry=https://github.com/open-telemetry/semantic-conventions/archive/refs/tags/$(TAG).zip[model] \
		--templates=/home/weaver/templates \
		--param tag=$(TAG) \
		--future \
		go \
		/home/weaver/target
