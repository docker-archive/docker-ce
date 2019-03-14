SHELL:=/bin/bash
ENGINE_DIR:=$(CURDIR)/../../engine
CLI_DIR:=$(CURDIR)/../../cli
VERSION?=0.0.0-dev
STATIC_VERSION=$(shell ./gen-static-ver $(ENGINE_DIR) $(VERSION))
CHOWN=docker run --rm -v $(CURDIR):/v -w /v alpine chown
HASH_CMD=docker run -v $(CURDIR):/sum -w /sum debian:jessie bash hash_files
DIR_TO_HASH:=build/linux
DEFAULT_PRODUCT_LICENSE?=Community Engine

DOCKER_CLI_GOLANG_IMG=$(shell awk '$$1=="FROM"{split($$2,a,"-");print a[1];exit}' $(CLI_DIR)/dockerfiles/Dockerfile.dev)

.PHONY: help
help: ## show make targets
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf " \033[36m%-20s\033[0m  %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: clean
clean: ## remove build artifacts
	[ ! -d build ] || $(CHOWN) -R $(shell id -u):$(shell id -g) build
	$(RM) -r build

.PHONY: static
static: static-linux cross-mac cross-win cross-arm ## create all static packages

.PHONY: static-linux
static-linux: static-cli static-engine ## create tgz with linux x86_64 client and server
	mkdir -p build/linux/docker
	cp $(CLI_DIR)/build/docker build/linux/docker/
	for f in dockerd containerd ctr containerd-shim docker-init docker-proxy runc; do \
		cp -L $(ENGINE_DIR)/bundles/binary-daemon/$$f build/linux/docker/$$f; \
	done
	tar -C build/linux -c -z -f build/linux/docker-$(STATIC_VERSION).tgz docker
	
	# extra binaries for running rootless
	mkdir -p build/linux/docker-rootless-extras
	for f in rootlesskit dockerd-rootless.sh vpnkit; do \
		if [ -f $(ENGINE_DIR)/bundles/binary-daemon/$$f ]; then \
			cp -L $(ENGINE_DIR)/bundles/binary-daemon/$$f build/linux/docker-rootless-extras/$$f; \
		fi \
	done
	tar -C build/linux -c -z -f build/linux/docker-rootless-extras-$(STATIC_VERSION).tgz docker-rootless-extras

.PHONY: hash_files
hash_files:
	@echo "Hashing directory $(DIR_TO_HASH)"
	$(HASH_CMD) "$(DIR_TO_HASH)"

.PHONY: cross-mac
cross-mac: cross-all-cli cross-mac-plugins ## create tgz with darwin x86_64 client only
	mkdir -p build/mac/docker
	cp $(CLI_DIR)/build/docker-darwin-amd64 build/mac/docker/docker
	tar -C build/mac -c -z -f build/mac/docker-$(STATIC_VERSION).tgz docker

.PHONY: cross-win
cross-win: cross-all-cli cross-win-engine ## create zip file with windows x86_64 client and server
	mkdir -p build/win/docker
	cp $(CLI_DIR)/build/docker-windows-amd64 build/win/docker/docker.exe
	cp $(ENGINE_DIR)/bundles/cross/windows/amd64/dockerd-$(STATIC_VERSION).exe build/win/docker/dockerd.exe
	docker run --rm -v $(CURDIR)/build/win:/v -w /v alpine sh -c 'apk update&&apk add zip&&zip -r docker-$(STATIC_VERSION).zip docker'
	$(CHOWN) -R $(shell id -u):$(shell id -g) build

.PHONY: cross-arm
cross-arm: cross-all-cli ## create tgz with linux armhf client only
	mkdir -p build/arm/docker
	cp $(CLI_DIR)/build/docker-linux-arm build/arm/docker/docker
	tar -C build/arm -c -z -f build/arm/docker-$(STATIC_VERSION).tgz docker

.PHONY: static-cli
static-cli:
	$(MAKE) -C $(CLI_DIR) -f docker.Makefile VERSION=$(STATIC_VERSION) build

.PHONY: static-engine
static-engine:
	$(MAKE) -C $(ENGINE_DIR) VERSION=$(STATIC_VERSION) binary

.PHONY: cross-all-cli
cross-all-cli:
	$(MAKE) -C $(CLI_DIR) -f docker.Makefile VERSION=$(STATIC_VERSION) cross

.PHONY: cross-win-engine
cross-win-engine:
	$(MAKE) -C $(ENGINE_DIR) VERSION=$(STATIC_VERSION) DOCKER_CROSSPLATFORMS=windows/amd64 cross

BUILD_MAC_RUN_VARS:=--rm -i \
	-e GOOS=darwin \
	-v "$(PWD)/build/mac/docker":/out \
	-v "$(PWD)/../plugins":/plugins:ro \
	-v "$(PWD)/scripts/build-mac-cli-plugins":/build:ro

.PHONY: cross-mac-plugins
cross-mac-plugins:
	mkdir -p build/mac/docker
	docker run $(BUILD_MAC_RUN_VARS) $(DOCKER_CLI_GOLANG_IMG) /build
	$(CHOWN) -R $(shell id -u):$(shell id -g) build
