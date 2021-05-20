# Docker CE Packaging

This repo contains the open source scripts for packaging the
[Docker Engine](https://docs.docker.com/engine/), the Docker CLI, CLI plugins,
and rootless-extras packages.

The repository contains Dockerfiles to build packages for various distributions,
which can be found in the "rpm" and "deb" subdirectories, as well as scripts to
build static binaries.

Docker uses these recipes to build and release packages that are available on the
https://download.docker.com package repositories. We welcome contributions to
this repository, including the addition of new distros or distro-versions. Note,
however, that Docker makes a subselection of distros and architectures for release,
and not all distros available in this repository may be released to download.docker.com,
but you can use these scripts to build your own packages.
