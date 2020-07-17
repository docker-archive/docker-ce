ARCH=$(shell uname -m)
BUILDTIME=$(shell date -u -d "@$${SOURCE_DATE_EPOCH:-$$(date +%s)}" --rfc-3339 ns 2> /dev/null | sed -e 's/ /T/')
CHOWN:=docker run --rm -v $(CURDIR):/v -w /v alpine chown
DEFAULT_PRODUCT_LICENSE:=Community Engine
DOCKER_GITCOMMIT:=abcdefg
GO_VERSION:=1.13.14
PLATFORM=Docker Engine - Community
SHELL:=/bin/bash
VERSION?=0.0.1-dev

# DOCKER_CLI_REPO and DOCKER_ENGINE_REPO define the source repositories to clone
# the source from. These can be overridden to build from a fork.
DOCKER_CLI_REPO    ?= https://github.com/docker/cli.git
DOCKER_ENGINE_REPO ?= https://github.com/docker/docker.git

# REF can be used to specify the same branch or tag to use for *both* the CLI
# and Engine source code. This can be useful if both the CLI and Engine have a
# release branch with the same name (e.g. "19.03"), or of both repositories have
# tagged a release with the same version.
#
# For other situations, specify DOCKER_CLI_REF and/or DOCKER_ENGINE_REF separately.
REF                ?= HEAD
DOCKER_CLI_REF     ?= $(REF)
DOCKER_ENGINE_REF  ?= $(REF)

export BUILDTIME
export DEFAULT_PRODUCT_LICENSE
export PLATFORM
