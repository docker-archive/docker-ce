include common.mk

STATIC_VERSION=$(shell static/gen-static-ver $(realpath $(CURDIR)/src/github.com/docker/docker) $(VERSION))

# Taken from: https://www.cmcrossroads.com/article/printing-value-makefile-variable
print-%  : ; @echo $($*)

.PHONY: help
help: ## show make targets
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf " \033[36m%-20s\033[0m  %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: clean-src
clean-src:
	[ ! -d src ] || $(CHOWN) -R $(shell id -u):$(shell id -g) src
	$(RM) -r src

.PHONY: src
src: src/github.com/docker/cli src/github.com/docker/docker src/github.com/docker/scan-cli-plugin ## clone source

ifdef CLI_DIR
src/github.com/docker/cli:
	mkdir -p "$(@D)"
	cp -r "$(CLI_DIR)" $@
else
src/github.com/docker/cli:
	git init $@
	git -C $@ remote add origin "$(DOCKER_CLI_REPO)"
endif

ifdef ENGINE_DIR
src/github.com/docker/docker:
	mkdir -p "$(@D)"
	cp -r "$(ENGINE_DIR)" $@
else
src/github.com/docker/docker:
	git init $@
	git -C $@ remote add origin "$(DOCKER_ENGINE_REPO)"
endif

src/github.com/docker/scan-cli-plugin:
	git init $@
	git -C $@ remote add origin "$(DOCKER_SCAN_REPO)"


.PHONY: checkout-cli
checkout-cli: src/github.com/docker/cli
	./scripts/checkout.sh src/github.com/docker/cli "$(DOCKER_CLI_REF)"

.PHONY: checkout-docker
checkout-docker: src/github.com/docker/docker
	./scripts/checkout.sh src/github.com/docker/docker "$(DOCKER_ENGINE_REF)"

.PHONY: checkout-scan-cli-plugin
checkout-scan-cli-plugin: src/github.com/docker/scan-cli-plugin
	./scripts/checkout.sh src/github.com/docker/scan-cli-plugin "$(DOCKER_SCAN_REF)"

.PHONY: checkout
checkout: checkout-cli checkout-docker checkout-scan-cli-plugin ## checkout source at the given reference(s)

.PHONY: clean
clean: clean-src ## remove build artifacts
	$(MAKE) -C rpm clean
	$(MAKE) -C deb clean
	$(MAKE) -C static clean

.PHONY: deb rpm
deb rpm: checkout ## build rpm/deb packages
	$(MAKE) -C $@ VERSION=$(VERSION) GO_VERSION=$(GO_VERSION) $@

.PHONY: centos-% fedora-% rhel-%
centos-% fedora-% rhel-%: checkout ## build rpm packages for the specified distro
	$(MAKE) -C rpm VERSION=$(VERSION) GO_VERSION=$(GO_VERSION) $@

.PHONY: debian-% raspbian-% ubuntu-%
debian-% raspbian-% ubuntu-%: checkout ## build deb packages for the specified distro
	$(MAKE) -C deb VERSION=$(VERSION) GO_VERSION=$(GO_VERSION) $@

.PHONY: static
static: DOCKER_BUILD_PKGS:=static-linux cross-mac cross-win cross-arm
static: checkout ## build static-compiled packages
	for p in $(DOCKER_BUILD_PKGS); do \
		$(MAKE) -C $@ VERSION=$(VERSION) GO_VERSION=$(GO_VERSION) TARGETPLATFORM=$(TARGETPLATFORM) CONTAINERD_VERSION=$(CONTAINERD_VERSION) RUNC_VERSION=$(RUNC_VERSION) $${p}; \
	done
