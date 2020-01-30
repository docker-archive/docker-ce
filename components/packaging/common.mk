ARCH=$(shell uname -m)
BUILDTIME=$(shell date -u -d "@$${SOURCE_DATE_EPOCH:-$$(date +%s)}" --rfc-3339 ns 2> /dev/null | sed -e 's/ /T/')
DEFAULT_PRODUCT_LICENSE:=Community Engine
DOCKER_GITCOMMIT:=abcdefg
GO_VERSION:=1.12.16
PLATFORM=Docker Engine - Community
SHELL:=/bin/bash
VERSION?=0.0.0-dev

export BUILDTIME
export DEFAULT_PRODUCT_LICENSE
export PLATFORM
