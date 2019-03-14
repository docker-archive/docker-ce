#!/usr/bin/env bash

DESTDIR=${DESTDIR:-}
PREFIX=${PREFIX:-/usr/local}

add_github_ssh_host() {
    # You're not able to clone from github unless you add to known_hosts
    if ! grep ~/.ssh/known_hosts "github.com" >/dev/null 2>/dev/null; then
        mkdir -p ~/.ssh
        ssh-keyscan github.com >> ~/.ssh/known_hosts
    fi
}

install_binary() {
    for binary in "$@"; do
        mkdir -p "${DESTDIR}${PREFIX}"
        install -p -m 755 "${binary}" "${DESTDIR}${PREFIX}"
    done
}

build_or_install() {
    case $1 in
        build)
            build
            ;;
        build_mac)
            build_mac
            ;;
        install_plugin)
            install_plugin
            ;;
        *)
            echo "Are you sure that's a command? o.O"
            exit 1
            ;;
    esac
}
