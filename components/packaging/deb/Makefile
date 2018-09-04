include ../containerd.mk

SHELL:=/bin/bash
ARCH:=$(shell uname -m)
CLI_DIR:=$(CURDIR)/../../cli
GITCOMMIT?=$(shell cd $(CLI_DIR) && git rev-parse --short HEAD)
VERSION?=0.0.0-dev
GO_BASE_IMAGE=golang
GO_VERSION:=1.10.4
GO_IMAGE=$(GO_BASE_IMAGE):$(GO_VERSION)
DEB_VERSION=$(shell ./gen-deb-ver $(CLI_DIR) "$(VERSION)")
CHOWN:=docker run --rm -v $(CURDIR):/v -w /v alpine chown
EPOCH?=2

COMMON_FILES=common
BUILD?=docker build \
	--build-arg GO_IMAGE=$(GO_IMAGE) \
	--build-arg COMMON_FILES=$(COMMON_FILES) \
	-t debbuild-$@/$(ARCH) \
	-f $(CURDIR)/$@/Dockerfile .
RUN=docker run --rm -i \
	-e EPOCH='$(EPOCH)' \
	-e DEB_VERSION=$(word 1, $(DEB_VERSION)) \
	-e VERSION=$(word 2, $(DEB_VERSION)) \
	-e DOCKER_GITCOMMIT=$(GITCOMMIT) \
	-v $(CURDIR)/debbuild/$@:/build \
	debbuild-$@/$(ARCH)

SOURCE_FILES=containerd-proxy.tgz cli.tgz containerd-shim-process.tar docker.service dockerd.json engine.tar
SOURCES=$(addprefix sources/, $(SOURCE_FILES))
ENGINE_IMAGE=docker/engine-community
ENGINE_SCOPE=ce

IMAGE_TAG=nightly

.PHONY: help
help: ## show make targets
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf " \033[36m%-20s\033[0m  %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: clean
clean: ## remove build artifacts
	[ ! -d debbuild ] || $(CHOWN) -R $(shell id -u):$(shell id -g) debbuild
	$(RM) -r debbuild
	[ ! -d sources ] || $(CHOWN) -R $(shell id -u):$(shell id -g) sources
	$(RM) -r sources
	[ ! -d artifacts ] || $(CHOWN) -R $(shell id -u):$(shell id -g) artifacts
	$(RM) -r artifacts
	-docker rm docker2oci

engine-$(ARCH).tar:
	$(MAKE) -C ../image image-linux
	docker save -o $@ $$(cat ../image/image-linux)

.PHONY: deb
deb: ubuntu debian raspbian ## build all deb packages

.PHONY: ubuntu
ubuntu: ubuntu-bionic ubuntu-xenial ubuntu-trusty ## build all ubuntu deb packages

.PHONY: debian
debian: debian-stretch debian-jessie ## build all debian deb packages

.PHONY: raspbian
raspbian: raspbian-stretch debian-jessie ## build all raspbian deb packages

.PHONY: ubuntu-bionic
ubuntu-bionic: ## build ubuntu bionic deb packages
ubuntu-bionic: $(SOURCES)
	$(BUILD)
	$(RUN)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

.PHONY: ubuntu-xenial
ubuntu-xenial: ## build ubuntu xenial deb packages
ubuntu-xenial: $(SOURCES)
	$(BUILD)
	$(RUN)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

.PHONY: ubuntu-trusty
ubuntu-trusty: ## build ubuntu trusty deb packages
ubuntu-trusty: $(SOURCES)
	$(BUILD)
	$(RUN)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

.PHONY: debian-buster
debian-buster: ## build debian buster deb packages
debian-buster: $(SOURCES)
	$(BUILD)
	$(RUN)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

.PHONY: debian-jessie
debian-jessie: ## build debian jessie deb packages
debian-jessie: $(SOURCES)
	$(BUILD)
	$(RUN)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

.PHONY: debian-stretch
debian-stretch: ## build debian stretch deb packages
debian-stretch: $(SOURCES)
	$(BUILD)
	$(RUN)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

.PHONY: raspbian-jessie
raspbian-jessie: ## build raspbian jessie deb packages
raspbian-jessie: $(SOURCES)
	$(BUILD)
	$(RUN)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

.PHONY: raspbian-stretch
raspbian-stretch: ## build raspbian stretch deb packages
raspbian-stretch: $(SOURCES)
	$(BUILD)
	$(RUN)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

sources/cli.tgz:
	mkdir -p $(@D)
	docker run --rm -i -w /v \
		-v $(CLI_DIR):/cli \
		-v $(CURDIR)/$(@D):/v \
		alpine \
		tar -C / -c -z -f /v/cli.tgz --exclude .git cli

sources/containerd-proxy.tgz:
	mkdir -p tmp/
	curl -fL -o tmp/containerd-proxy.tgz "https://github.com/crosbymichael/containerd-proxy/archive/$(CONTAINERD_PROXY_COMMIT).tar.gz"
	tar xzf tmp/containerd-proxy.tgz -C tmp/
	mv tmp/containerd-proxy-$(CONTAINERD_PROXY_COMMIT) tmp/containerd-proxy
	mkdir -p $(@D)
	$(CHOWN) -R $(shell id -u):$(shell id -g) $$(dirname $(@D))
	tar -zcf $@ -C tmp/ containerd-proxy
	rm -rf tmp/

sources/containerd-shim-process.tar:
	$(CTR) content fetch $(CONTAINERD_SHIM_PROCESS_IMAGE)
	$(CTR) image export artifacts/containerd-shim-process.tar $(CONTAINERD_SHIM_PROCESS_IMAGE)
	mkdir -p $(@D)
	cp artifacts/containerd-shim-process.tar $@
	$(CHOWN) -R $(shell id -u):$(shell id -g) $$(dirname $(@D))

sources/docker.service: ../systemd/docker.service
	mkdir -p $(@D)
	cp $< $@

sources/dockerd.json: ../common/dockerd.json
	mkdir -p $(@D)
	sed \
	    -e 's!$${ENGINE_IMAGE}!$(ENGINE_IMAGE)!' \
	    -e 's!$${ENGINE_SCOPE}!$(ENGINE_SCOPE)!' \
	    -e 's/$${IMAGE_TAG}/$(IMAGE_TAG)/' \
	    $< > $@

# TODO: Eventually clean this up when we release an image with a manifest
DOCKER2OCI=artifacts/docker2oci
$(DOCKER2OCI):
	-$(CHOWN) -R $(shell id -u):$(shell id -g) $(@D)
	docker run --name docker2oci $(GO_IMAGE) sh -c 'go get github.com/coolljt0725/docker2oci'
	mkdir -p $(@D)
	docker cp docker2oci:/go/bin/docker2oci "$@"
	docker rm -f docker2oci
	$(CHOWN) -R $(shell id -u):$(shell id -g) $(@D)

# offline bundle
sources/engine.tar: $(DOCKER2OCI)
	$(MAKE) -C ../image ENGINE_IMAGE=$(ENGINE_IMAGE) ENGINE_SCOPE=$(ENGINE_SCOPE) image-linux
	mkdir -p artifacts
	docker save -o artifacts/docker-engine.tar $$(cat ../image/image-linux)
	./$(DOCKER2OCI) -i artifacts/docker-engine.tar artifacts/engine-image
	mkdir -p $(@D)
	tar c -C artifacts/engine-image . > $@
