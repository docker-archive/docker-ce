ARCH=$(shell uname -m)
ALPINE_IMG:=$(shell $(CURDIR)/../detect_alpine_image)
ENGINE_DIR:=$(CURDIR)/../../engine
CLI_DIR:=$(CURDIR)/../../cli
GITCOMMIT=$(shell cd $(ENGINE_DIR) && git rev-parse --short HEAD)
VERSION=$(shell cat $(ENGINE_DIR)/VERSION)
DOCKER_EXPERIMENTAL=0
GEN_RPM_VER=$(shell ./gen-rpm-ver $(ENGINE_DIR) $(VERSION))
CHOWN=docker run --rm -i -v $(CURDIR):/v -w /v $(ALPINE_IMG) chown
RPMBUILD=docker run --privileged --rm -i\
	-v $(CURDIR)/rpmbuild/SOURCES:/root/rpmbuild/SOURCES \
	-v $(CURDIR)/rpmbuild/BUILD:/root/rpmbuild/BUILD \
	-v $(CURDIR)/rpmbuild/BUILDROOT:/root/rpmbuild/BUILDROOT \
	-v $(CURDIR)/rpmbuild/RPMS:/root/rpmbuild/RPMS \
	-v $(CURDIR)/rpmbuild/SRPMS:/root/rpmbuild/SRPMS \
	-v $(CURDIR)/systemd:/systemd
RPMBUILD_FLAGS=-ba\
	--define '_gitcommit $(word 3,$(GEN_RPM_VER))' \
	--define '_release $(word 2,$(GEN_RPM_VER))' \
	--define '_version $(word 1,$(GEN_RPM_VER))' \
	--define '_origversion $(VERSION)' \
	--define '_experimental $(DOCKER_EXPERIMENTAL)' \
	SPECS/docker-ce.spec

.PHONY: help clean rpm fedora centos fedora-26 fedora-25 centos-7

help: ## show make targets
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf " \033[36m%-20s\033[0m  %s\n", $$1, $$2}' $(MAKEFILE_LIST)

clean: ## remove build artifacts
	[ ! -d rpmbuild ] || $(CHOWN) -R $(shell id -u):$(shell id -g) rpmbuild
	rm -rf rpmbuild

rpm: fedora centos ## build all rpm packages

fedora: fedora-26 fedora-25 ## build all fedora rpm packages

centos: centos-7 ## build all centos rpm packages

fedora-26: rpmbuild/SOURCES/engine.tgz rpmbuild/SOURCES/cli.tgz ## build fedora-26 rpm packages
	docker build -t rpmbuild-$@/$(ARCH) -f $@/Dockerfile.$(ARCH) $@
	$(RPMBUILD) rpmbuild-$@/$(ARCH) $(RPMBUILD_FLAGS)
	$(CHOWN) -R $(shell id -u):$(shell id -g) rpmbuild

fedora-25: rpmbuild/SOURCES/engine.tgz rpmbuild/SOURCES/cli.tgz ## build fedora-25 rpm packages
	docker build -t rpmbuild-$@/$(ARCH) -f $@/Dockerfile.$(ARCH) $@
	$(RPMBUILD) rpmbuild-$@/$(ARCH) $(RPMBUILD_FLAGS)
	$(CHOWN) -R $(shell id -u):$(shell id -g) rpmbuild

centos-7: rpmbuild/SOURCES/engine.tgz rpmbuild/SOURCES/cli.tgz ## build centos-7 rpm packages
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
