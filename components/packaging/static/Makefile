SHELL:=/bin/bash
ENGINE_DIR:=$(CURDIR)/../../engine
CLI_DIR:=$(CURDIR)/../../cli
ENGINE_VER=$(shell cat $(ENGINE_DIR)/VERSION)
VERSION=$(shell cat $(ENGINE_DIR)/VERSION)
CHOWN=docker run --rm -v $(CURDIR):/v -w /v alpine chown

.PHONY: help clean static static-linux cross-mac cross-win cross-arm

help: ## show make targets
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf " \033[36m%-20s\033[0m  %s\n", $$1, $$2}' $(MAKEFILE_LIST)

clean: ## remove build artifacts
	$(RM) -r build

$(CLI_DIR)/build/docker:
	$(MAKE) -C $(CLI_DIR) -f docker.Makefile build

$(ENGINE_DIR)/bundles/$(ENGINE_VER)/binary-daemon/dockerd:
	$(MAKE) -C $(ENGINE_DIR) binary

#$(CLI_DIR)/build/docker-linux-amd64:
#$(CLI_DIR)/build/docker-linux-arm:
#$(CLI_DIR)/build/docker-darwin-amd64:
#$(CLI_DIR)/build/docker-windows-amd64:

compile-cross-cli:
	$(MAKE) -C $(CLI_DIR) -f docker.Makefile VERSION=$(VERSION) cross

static: static-linux cross-mac cross-win cross-arm

static-linux: $(CLI_DIR)/build/docker $(ENGINE_DIR)/bundles/$(ENGINE_VER)/binary-daemon/dockerd
	mkdir -p build/linux/docker
	cp $? build/linux/docker
	tar -C build/linux -c -z -f build/linux/docker-$(VERSION).tgz docker

cross-mac: compile-cross-cli ## create tgz with darwin x86_64 client only
	mkdir -p build/mac/docker
	cp $(CLI_DIR)/build/docker-darwin-amd64 build/mac/docker/docker
	tar -C build/mac -c -z -f build/mac/docker-$(VERSION).tgz docker

cross-win: | $(CLI_DIR)/build/docker-windows-amd64 $(ENGINE_DIR)/bundles/$(ENGINE_VER)/cross/windows/amd64/dockerd-$(ENGINE_VER).exe ## create zip file with windows x86_64 client and server
	mkdir -p build/win/docker
	cp $(CLI_DIR)/build/docker-windows-amd64 build/win/docker/docker.exe
	cp $(ENGINE_DIR)/bundles/$(ENGINE_VER)/cross/windows/amd64/dockerd-$(ENGINE_VER).exe build/win/docker/dockerd.exe
	docker run --rm -v $(CURDIR)/build/win:/v -w /v alpine sh -c 'apk update&&apk add zip&&zip -r docker-$(VERSION).zip docker'
	$(CHOWN) -R $(shell id -u):$(shell id -g) build

cross-arm: $(CLI_DIR)/build/docker-linux-arm ## create tgz with linux armhf client only
	mkdir -p build/arm/docker
	cp $< build/arm/docker/docker
	tar -C build/arm -c -z -f build/arm/docker-$(VERSION).tgz docker
