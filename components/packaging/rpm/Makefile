include ../containerd.mk

ARCH=$(shell uname -m)
ENGINE_DIR:=$(CURDIR)/../../engine
CLI_DIR:=$(CURDIR)/../../cli
GITCOMMIT=$(shell cd $(ENGINE_DIR) && git rev-parse --short HEAD)
VERSION?=0.0.0-dev
GO_BASE_IMAGE=golang
GO_VERSION:=1.10.6
GO_IMAGE?=$(GO_BASE_IMAGE):$(GO_VERSION)
GEN_RPM_VER=$(shell ./gen-rpm-ver $(CLI_DIR) $(VERSION))
CHOWN=docker run --rm -i -v $(CURDIR):/v -w /v alpine chown

DOCKERFILE=Dockerfile
ifdef NEEDS_ARCH_SPECIFIC
	DOCKERFILE=Dockerfile.$(ARCH)
endif
ifdef BUILD_IMAGE
	BUILD_IMAGE_FLAG=--build-arg $(BUILD_IMAGE)
endif
BUILD?=docker build \
	$(BUILD_IMAGE_FLAG) \
	--build-arg GO_IMAGE=$(GO_IMAGE) \
	--build-arg ENGINE_IMAGE=$(shell cat rpmbuild/SOURCES/engine-image) \
	-t rpmbuild-$@/$(ARCH) \
	-f $@/$(DOCKERFILE) \
	.

SPEC_FILES?=docker-ce.spec docker-ce-cli.spec
SPECS?=$(addprefix SPECS/, $(SPEC_FILES))
RPMBUILD=docker run --privileged --rm -i\
	-v $(CURDIR)/rpmbuild/SOURCES:/root/rpmbuild/SOURCES \
	-v $(CURDIR)/rpmbuild/RPMS:/root/rpmbuild/RPMS \
	-v $(CURDIR)/rpmbuild/SRPMS:/root/rpmbuild/SRPMS
RPMBUILD_FLAGS?=-ba\
	--define '_gitcommit $(word 3,$(GEN_RPM_VER))' \
	--define '_release $(word 2,$(GEN_RPM_VER))' \
	--define '_version $(word 1,$(GEN_RPM_VER))' \
	--define '_origversion $(word 4, $(GEN_RPM_VER))' \
	$(SPECS)
RUN?=$(RPMBUILD) rpmbuild-$@/$(ARCH) $(RPMBUILD_FLAGS)

SOURCE_FILES=engine-image cli.tgz docker.service docker.socket distribution_based_engine.json
SOURCES=$(addprefix rpmbuild/SOURCES/, $(SOURCE_FILES))


.PHONY: help
help: ## show make targets
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf " \033[36m%-20s\033[0m  %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: clean
clean: ## remove build artifacts
	[ ! -d rpmbuild ] || $(CHOWN) -R $(shell id -u):$(shell id -g) rpmbuild
	$(RM) -r rpmbuild/
	[ ! -d artifacts ] || $(CHOWN) -R $(shell id -u):$(shell id -g) artifacts
	$(RM) -r artifacts/
	[ ! -d tmp ] || $(CHOWN) -R $(shell id -u):$(shell id -g) tmp
	$(RM) -r tmp/
	-docker rm docker2oci
	$(MAKE) -C ../image clean

.PHONY: rpm
rpm: fedora centos ## build all rpm packages

.PHONY: fedora
fedora: fedora-28 fedora-27 fedora-26 ## build all fedora rpm packages

.PHONY: centos
centos: centos-7 ## build all centos rpm packages

.PHONY: fedora-29
fedora-29: ## build fedora-29 rpm packages
fedora-29: $(SOURCES)
	$(CHOWN) -R root:root rpmbuild
	$(BUILD)
	$(RUN)
	$(CHOWN) -R $(shell id -u):$(shell id -g) rpmbuild

.PHONY: fedora-28
fedora-28: ## build fedora-28 rpm packages
fedora-28: $(SOURCES)
	$(CHOWN) -R root:root rpmbuild
	$(BUILD)
	$(RUN)
	$(CHOWN) -R $(shell id -u):$(shell id -g) rpmbuild

.PHONY: fedora-27
fedora-27: ## build fedora-27 rpm packages
fedora-27: $(SOURCES)
	$(CHOWN) -R root:root rpmbuild
	$(BUILD)
	$(RUN)
	$(CHOWN) -R $(shell id -u):$(shell id -g) rpmbuild

.PHONY: centos-7
centos-7: ## build centos-7 rpm packages
centos-7: $(SOURCES)
	$(CHOWN) -R root:root rpmbuild
	$(BUILD)
	$(RUN)
	$(CHOWN) -R $(shell id -u):$(shell id -g) rpmbuild

rpmbuild/SOURCES/cli.tgz:
	mkdir -p rpmbuild/SOURCES
	docker run --rm -i -w /v \
		-v $(CLI_DIR):/cli \
		-v $(CURDIR)/rpmbuild/SOURCES:/v \
		alpine \
		tar -C / -c -z -f /v/cli.tgz --exclude .git cli

rpmbuild/SOURCES/docker.service: ../systemd/docker.service
	mkdir -p $(@D)
	cp $< $@

rpmbuild/SOURCES/docker.socket: ../systemd/docker.socket
	mkdir -p $(@D)
	cp $< $@

rpmbuild/SOURCES/engine-image:
	$(MAKE) -C ../image image-linux-dm
	mkdir -p $(@D)
	cp ../image/image-linux-dm $@

rpmbuild/SOURCES/distribution_based_engine.json: rpmbuild/SOURCES/engine-image
	mkdir -p $(@D)
	docker inspect "$(shell cat $<)" \
		--format '{{index .Config.Labels "com.docker.distribution_based_engine" }}' > $@
