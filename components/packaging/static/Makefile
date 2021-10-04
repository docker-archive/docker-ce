include ../common.mk

CLI_DIR=$(realpath $(CURDIR)/../src/github.com/docker/cli)
ENGINE_DIR=$(realpath $(CURDIR)/../src/github.com/docker/docker)
GEN_STATIC_VER=$(shell ./gen-static-ver $(CLI_DIR) $(VERSION))
HASH_CMD=docker run -v $(CURDIR):/sum -w /sum debian:jessie bash hash_files
DIR_TO_HASH:=build/linux
DOCKER_CLI_GOLANG_IMG=golang:$(GO_VERSION)

DOCKER_BUILD_OPTS=

ifneq ($(strip $(CONTAINERD_VERSION)),)
# Set custom build-args to override the containerd version to build for static
# packages. The Dockerfile for 20.10 and earlier used CONTAINERD_COMMIT, later
# versions use CONTAINERD_VERSION. We can remove CONTAINERD_VERSION once 20.10.x
# reaches EOL.
DOCKER_BUILD_OPTS +=--build-arg=CONTAINERD_VERSION=$(CONTAINERD_VERSION)
DOCKER_BUILD_OPTS +=--build-arg=CONTAINERD_COMMIT=$(CONTAINERD_VERSION)
endif

ifneq ($(strip $(RUNC_VERSION)),)
# Set custom build-args to override the runc version to build for static packages.
# The Dockerfile for 20.10 and earlier used RUNC_COMMIT, later versions use
# RUNC_VERSION. We can remove RUNC_COMMIT once 20.10.x reaches EOL.
DOCKER_BUILD_OPTS +=--build-arg=RUNC_VERSION=$(RUNC_VERSION)
DOCKER_BUILD_OPTS +=--build-arg=RUNC_COMMIT=$(RUNC_VERSION)
endif

.PHONY: help
help: ## show make targets
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf " \033[36m%-20s\033[0m  %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: clean
clean: ## remove build artifacts
	[ ! -d build ] || $(CHOWN) -R $(shell id -u):$(shell id -g) build
	$(RM) -r build
	-docker builder prune -f --filter until=24h

.PHONY: static
static: static-linux cross-mac cross-win cross-arm ## create all static packages

.PHONY: static-linux
static-linux: static-cli static-engine ## create tgz with linux x86_64 client and server
	mkdir -p build/linux/docker
	cp $(CLI_DIR)/build/docker build/linux/docker/
	for f in dockerd containerd ctr containerd-shim containerd-shim-runc-v2 docker-init docker-proxy runc; do \
		cp -L $(ENGINE_DIR)/bundles/binary-daemon/$$f build/linux/docker/$$f; \
	done
	tar -C build/linux -c -z -f build/linux/docker-$(GEN_STATIC_VER).tgz docker

	# extra binaries for running rootless
	mkdir -p build/linux/docker-rootless-extras
	for f in rootlesskit rootlesskit-docker-proxy dockerd-rootless.sh dockerd-rootless-setuptool.sh vpnkit; do \
		if [ -f $(ENGINE_DIR)/bundles/binary-daemon/$$f ]; then \
			cp -L $(ENGINE_DIR)/bundles/binary-daemon/$$f build/linux/docker-rootless-extras/$$f; \
		fi \
	done
	tar -C build/linux -c -z -f build/linux/docker-rootless-extras-$(GEN_STATIC_VER).tgz docker-rootless-extras

.PHONY: hash_files
hash_files:
	@echo "Hashing directory $(DIR_TO_HASH)"
	$(HASH_CMD) "$(DIR_TO_HASH)"

.PHONY: buildx
buildx:
	docker buildx inspect | grep -q 'Driver: docker-container' || docker buildx create --use

.PHONY: cross-mac
cross-mac: buildx cross-mac-plugins
	cd $(CLI_DIR) && VERSION=$(GEN_STATIC_VER) docker buildx bake --set binary.platform=darwin/amd64,darwin/arm64 binary
	dest=$$PWD/build/mac; cd $(CLI_DIR)/build && for platform in *; do \
		arch=$$(echo $$platform | cut -d_ -f2); \
		mkdir -p $$dest/$$arch/docker; \
		cp $$platform/docker-darwin-* $$dest/$$arch/docker/docker && \
		tar -C $$dest/$$arch -c -z -f $$dest/$$arch/docker-$(GEN_STATIC_VER).tgz docker; \
	done

.PHONY: cross-win
cross-win: cross-win-engine cross-win-plugins
	cd $(CLI_DIR) && VERSION=$(GEN_STATIC_VER) docker buildx bake --set binary.platform=windows/amd64 binary
	mkdir -p build/win/amd64/docker
	cp $(CLI_DIR)/build/docker-windows-amd64.exe build/win/amd64/docker/docker.exe
	cp $(ENGINE_DIR)/bundles/cross/windows/amd64-daemon/dockerd-$(GEN_STATIC_VER).exe build/win/amd64/docker/dockerd.exe
	cp $(ENGINE_DIR)/bundles/cross/windows/amd64-daemon/docker-proxy-$(GEN_STATIC_VER).exe build/win/amd64/docker/docker-proxy.exe
	docker run --rm -v $(CURDIR)/build/win/amd64:/v -w /v alpine sh -c 'apk update&&apk add zip&&zip -r docker-$(GEN_STATIC_VER).zip docker'
	$(CHOWN) -R $(shell id -u):$(shell id -g) build

.PHONY: cross-arm
cross-arm: cross-all-cli ## create tgz with linux armhf client only
	mkdir -p build/arm/docker
	cp $(CLI_DIR)/build/docker-linux-arm build/arm/docker/docker
	tar -C build/arm -c -z -f build/arm/docker-$(GEN_STATIC_VER).tgz docker

.PHONY: static-cli
static-cli:
	cd $(CLI_DIR) && VERSION=$(GEN_STATIC_VER) docker buildx bake --set binary.platform=$(TARGETPLATFORM) --set binary.args.CGO_ENABLED=$(CGO_ENABLED) binary

.PHONY: static-engine
static-engine:
	$(MAKE) -C $(ENGINE_DIR) VERSION=$(GEN_STATIC_VER) DOCKER_BUILD_OPTS="$(DOCKER_BUILD_OPTS)" binary

.PHONY: cross-all-cli
cross-all-cli:
	$(MAKE) -C $(CLI_DIR) -f docker.Makefile VERSION=$(GEN_STATIC_VER) cross

.PHONY: cross-win-engine
cross-win-engine:
	$(MAKE) -C $(ENGINE_DIR) VERSION=$(GEN_STATIC_VER) DOCKER_CROSSPLATFORMS=windows/amd64 DOCKER_BUILD_OPTS="$(DOCKER_BUILD_OPTS)" cross

BUILD_PLUGIN_RUN_VARS = --rm \
	-e GOOS \
	-e GOARCH \
	-v "$(CURDIR)/build/$(CLI_BUILD_DIR)/$*/docker/cli-plugins":/out \
	-v "$(CURDIR)/../plugins":/plugins:ro \
	-v "$(CURDIR)/scripts/build-cli-plugins":/build:ro

.PHONY: cross-mac-plugins
cross-mac-plugins: cross-mac-plugins-amd64 cross-mac-plugins-arm64

.PHONY: cross-mac-plugins-%
cross-mac-plugins-%: CLI_BUILD_DIR := mac
cross-mac-plugins-%:
	mkdir -p build/$(CLI_BUILD_DIR)/$*/docker
	GOOS=darwin GOARCH=$* docker run $(BUILD_PLUGIN_RUN_VARS) $(DOCKER_CLI_GOLANG_IMG) /build
	$(CHOWN) -R $(shell id -u):$(shell id -g) build/$(CLI_BUILD_DIR)/$*

.PHONY: cross-win-plugins
cross-win-plugins: cross-win-plugins-amd64

.PHONY: cross-win-plugins-%
cross-win-plugins-%: CLI_BUILD_DIR := win
cross-win-plugins-%:
	mkdir -p build/$(CLI_BUILD_DIR)/$*/docker/cli-plugins
	GOOS=windows GOARCH=$* docker run $(BUILD_PLUGIN_RUN_VARS) $(DOCKER_CLI_GOLANG_IMG) /build
	$(CHOWN) -R $(shell id -u):$(shell id -g) build/$(CLI_BUILD_DIR)/$*
	find build/$(CLI_BUILD_DIR)/$*/docker -type f -not -name "*.exe" -exec mv {} {}.exe \;
