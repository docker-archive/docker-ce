#!/usr/bin/env bash

set -e

source "$(dirname "$0")/.common"
PKG=github.com/docker/scan-cli-plugin
GOPATH=$(go env GOPATH)
REPO=https://${PKG}.git
COMMIT=v0.10.0
DEST=${GOPATH}/src/${PKG}

build() {
    if [ ! -d "${DEST}" ]; then
        git clone "${REPO}" "${DEST}"
    fi
    (
        cd "${DEST}"
        if [ -d ".git" ]; then
            git fetch --all
            git checkout -q "${COMMIT}"
        fi
        # Using goproxy instead of "direct" to work around an issue in go mod
        # not working with older git versions (default version on CentOS 7 is
        # git 1.8), see https://github.com/golang/go/issues/38373
        GOPROXY="https://proxy.golang.org" PLATFORM_BINARY=docker-scan TAG_NAME="${COMMIT}" make native-build
    )
}

install_plugin() {
    (
        cd "${DEST}"
        install_binary bin/docker-scan
    )
}

case "$(uname -i)" in
  aarch64)
    echo "Skipping scan plugin on ARM arch";;
  arm*)
    echo "Skipping scan plugin on ARM arch";;
  *)
    build_or_install "$@";;
esac
