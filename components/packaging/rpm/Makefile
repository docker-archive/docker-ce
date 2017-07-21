ARCH=$(shell uname -m)
ALPINE_IMG:=$(shell $(CURDIR)/../detect_alpine_image)
KERNEL:=$(shell docker version --format '{{.Server.Os}}')
ENGINE_DIR:=$(CURDIR)/../../engine
COMMON_DIR:=$(CURDIR)/../common
PLUGIN_DIR:=$(CURDIR)/plugin
CLI_DIR:=$(CURDIR)/../../cli
GITCOMMIT=$(shell cd $(ENGINE_DIR) && git rev-parse --short HEAD)
VERSION=$(shell cat $(ENGINE_DIR)/VERSION)
PLUGIN_VERSION?=1.0.0.$(KERNEL)-$(ARCH)-test
DOCKER_EXPERIMENTAL=0
GEN_RPM_VER=$(shell ./gen-rpm-ver $(ENGINE_DIR) $(VERSION))
CHOWN=docker run --rm -i -v $(CURDIR):/v -w /v $(ALPINE_IMG) chown
RPMBUILD=docker run --privileged --rm -i\
	-v $(CURDIR)/rpmbuild/SOURCES:/root/rpmbuild/SOURCES \
	-v $(CURDIR)/rpmbuild/BUILD:/root/rpmbuild/BUILD \
	-v $(CURDIR)/rpmbuild/BUILDROOT:/root/rpmbuild/BUILDROOT \
	-v $(CURDIR)/rpmbuild/RPMS:/root/rpmbuild/RPMS \
	-v $(CURDIR)/rpmbuild/SRPMS:/root/rpmbuild/SRPMS \
	-v $(CURDIR)/systemd:/systemd \
	-v $(COMMON_DIR):/common
RPMBUILD_FLAGS=-ba\
	--define '_gitcommit $(word 3,$(GEN_RPM_VER))' \
	--define '_release $(word 2,$(GEN_RPM_VER))' \
	--define '_version $(word 1,$(GEN_RPM_VER))' \
	--define '_plugin_version $(PLUGIN_VERSION)' \
	--define '_origversion $(VERSION)' \
	--define '_experimental $(DOCKER_EXPERIMENTAL)' \
	SPECS/docker-ce.spec

.PHONY: help clean rpm fedora centos fedora-25 fedora-24 centos-7

help: ## show make targets
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf " \033[36m%-20s\033[0m  %s\n", $$1, $$2}' $(MAKEFILE_LIST)

clean: ## remove build artifacts
	[ ! -d rpmbuild ] || $(CHOWN) -R $(shell id -u):$(shell id -g) rpmbuild
	rm -rf rpmbuild

rpm: fedora centos ## build all rpm packages

fedora: fedora-25 fedora-24 ## build all fedora rpm packages

centos: centos-7 ## build all centos rpm packages

fedora-25: rpmbuild/SOURCES/engine.tgz rpmbuild/SOURCES/cli.tgz rpmbuild/SOURCES/telemetry.tgz ## build fedora-25 rpm packages
	docker build -t rpmbuild-$@/$(ARCH) -f $@/Dockerfile.$(ARCH) $@
	$(RPMBUILD) rpmbuild-$@/$(ARCH) $(RPMBUILD_FLAGS)
	$(CHOWN) -R $(shell id -u):$(shell id -g) rpmbuild

fedora-24: rpmbuild/SOURCES/engine.tgz rpmbuild/SOURCES/cli.tgz rpmbuild/SOURCES/telemetry.tgz ## build fedora-24 rpm packages
	docker build -t rpmbuild-$@/$(ARCH) -f $@/Dockerfile.$(ARCH) $@
	$(RPMBUILD) rpmbuild-$@/$(ARCH) $(RPMBUILD_FLAGS)
	$(CHOWN) -R $(shell id -u):$(shell id -g) rpmbuild

centos-7: rpmbuild/SOURCES/engine.tgz rpmbuild/SOURCES/cli.tgz rpmbuild/SOURCES/telemetry.tgz ## build centos-7 rpm packages
	docker build -t rpmbuild-$@/$(ARCH) -f $@/Dockerfile.$(ARCH) $@
	$(RPMBUILD) rpmbuild-$@/$(ARCH) $(RPMBUILD_FLAGS)
	$(CHOWN) -R $(shell id -u):$(shell id -g) rpmbuild

rpmbuild/SOURCES/cli.tgz:
	mkdir -p rpmbuild/SOURCES
	docker run --rm -i -w /v \
		-v $(CLI_DIR):/cli \
		-v $(CURDIR)/rpmbuild/SOURCES:/v \
		$(ALPINE_IMG) \
		tar -C / -c -z -f /v/cli.tgz --exclude .git cli

rpmbuild/SOURCES/engine.tgz:
	mkdir -p rpmbuild/SOURCES
	docker run --rm -i -w /v \
		-v $(ENGINE_DIR):/engine \
		-v $(CURDIR)/rpmbuild/SOURCES:/v \
		$(ALPINE_IMG) \
		tar -C / -c -z -f /v/engine.tgz --exclude .git engine

rpmbuild/SOURCES/telemetry.tgz:
	mkdir -p rpmbuild/SOURCES
	mkdir -p $(PLUGIN_DIR)
	docker create --name tmp docker/telemetry-tgz:$(PLUGIN_VERSION)
	docker export tmp | tar -x -C $(PLUGIN_DIR)
	docker run --rm -i -w /plugin \
		-v $(PLUGIN_DIR):/plugin \
		-v $(CURDIR)/rpmbuild/SOURCES:/v \
		alpine \
		cp telemetry_$(PLUGIN_VERSION).tgz /v
	docker rm -f tmp
