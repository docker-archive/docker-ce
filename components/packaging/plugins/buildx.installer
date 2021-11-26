#!/usr/bin/env bash

set -e

source "$(dirname "$0")/.common"
PKG=github.com/docker/buildx
GOPATH=$(go env GOPATH)
REPO=https://${PKG}.git
: "${BUILDX_COMMIT=v0.7.1}"
DEST=${GOPATH}/src/${PKG}

build() {
    if [ ! -d "${DEST}" ]; then
        git clone "${REPO}" "${DEST}"
    fi
    (
        cd "${DEST}"
        git fetch --all
        git checkout -q "${BUILDX_COMMIT}"
        local LDFLAGS
        LDFLAGS="-X ${PKG}/version.Version=$(git describe --match 'v[0-9]*' --always --tags)-docker -X ${PKG}/version.Revision=$(git rev-parse HEAD) -X ${PKG}/version.Package=${PKG}"
        set -x
        GO111MODULE=on go build -mod=vendor -o bin/docker-buildx -ldflags "${LDFLAGS}" ./cmd/buildx
    )
}

install_plugin() {
    (
        cd "${DEST}"
        install_binary bin/docker-buildx
    )
}

build_or_install "$@"
