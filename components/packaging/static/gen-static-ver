#!/usr/bin/env bash

REPO_DIR="$1"
VERSION="$2"

if [ -z "$REPO_DIR" ] || [ -z "$VERSION" ]; then
	# shellcheck disable=SC2016
	echo 'usage: ./gen-static-ver ${REPO_DIR} ${VERSION}'
	exit 1
fi

GIT_COMMAND="git -C $REPO_DIR"

staticVersion="$VERSION"
if [[ "$VERSION" == *-dev ]]; then
	export TZ=UTC

	DATE_COMMAND="date"
	if [[ $(uname) == "Darwin" ]]; then
		DATE_COMMAND="docker run --rm alpine date"
	fi

	# based on golang's pseudo-version: https://groups.google.com/forum/#!topic/golang-dev/a5PqQuBljF4
	#
	# using a "pseudo-version" of the form v0.0.0-yyyymmddhhmmss-abcdefabcdef,
	# where the time is the commit time in UTC and the final suffix is the prefix
	# of the commit hash. The time portion ensures that two pseudo-versions can
	# be compared to determine which happened later, the commit hash identifes
	# the underlying commit, and the v0.0.0- prefix identifies the pseudo-version
	# as a pre-release before version v0.0.0, so that the go command prefers any
	# tagged release over any pseudo-version.
	gitUnix="$($GIT_COMMAND log -1 --pretty='%ct')"
	gitDate="$($DATE_COMMAND --utc --date "@$gitUnix" +'%Y%m%d%H%M%S')"
	gitCommit="$($GIT_COMMAND log -1 --pretty='%h')"
	# generated version is now something like '0.0.0-20180719213702-cd5e2db'
	staticVersion="0.0.0-${gitDate}-${gitCommit}"
fi

echo "$staticVersion"
