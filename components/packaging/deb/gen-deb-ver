#!/usr/bin/env bash

REPO_DIR="$1"
VERSION="$2"

if [ -z "$REPO_DIR" ] || [ -z "$VERSION" ]; then
	# shellcheck disable=SC2016
	echo 'usage: ./gen-deb-ver ${REPO_DIR} ${VERSION}'
	exit 1
fi

GIT_COMMAND="git -C $REPO_DIR"
origVersion="$VERSION"
debVersion="${VERSION#v}"

gen_deb_version() {
	# Adds an increment to the deb version to get proper order
	# 18.01.0-tp1   -> 18.01.0-0.1-tp1
	# 18.01.0-beta1 -> 18.01.0-1.1-beta1
	# 18.01.0-rc1   -> 18.01.0-2.1-rc1
	# 18.01.0       -> 18.01.0-3
	fullVersion="$1"
	pattern="$2"
	increment="$3"
	testVersion="${fullVersion#*-$pattern}"
	baseVersion="${fullVersion%-"$pattern"*}"
	echo "$baseVersion-$increment.$testVersion.$pattern$testVersion"
}

case "$debVersion" in
	*-dev)
		;;
	*-tp[0-9]*)
		debVersion="$(gen_deb_version "$debVersion" tp 0)"
		;;
	*-beta[0-9]*)
		debVersion="$(gen_deb_version "$debVersion" beta 1)"
		;;
	*-rc[0-9]*)
		debVersion="$(gen_deb_version "$debVersion" rc 2)"
		;;
	*)
		debVersion="$debVersion-3"
		;;
esac

tilde='~'                            # ouch Bash 4.2 vs 4.3, you keel me
debVersion="${debVersion//-/$tilde}" # using \~ or '~' here works in 4.3, but not 4.2; just ~ causes $HOME to be inserted, hence the $tilde
# if we have a "-dev" suffix or have change in Git, let's make this package version more complex so it works better
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
	debVersion="0.0.0-${gitDate}-${gitCommit}"
	origVersion=$debVersion

	# verify that nightly builds are always < actual releases
	#
	# $ dpkg --compare-versions 1.5.0 gt 1.5.0~rc1 && echo true || echo false
	# true
	# $ dpkg --compare-versions 1.5.0~rc1 gt 0.0.0-20180719213347-5daff5a && echo true || echo false
	# true
	# $ dpkg --compare-versions 18.06.0-ce-rc3 gt 18.06.0-ce-rc2  && echo true || echo false
	# true
	# $ dpkg --compare-versions 18.06.0-ce gt 18.06.0-ce-rc2  && echo true || echo false
	# false
	# $ dpkg --compare-versions 18.06.0-ce-rc3 gt 0.0.0-20180719213347-5daff5a  && echo true || echo false
	# true
	# $ dpkg --compare-versions 18.06.0-ce gt 0.0.0-20180719213347-5daff5a  && echo true || echo false
	# true
	# $ dpkg --compare-versions 0.0.0-20180719213702-cd5e2db gt 0.0.0-20180719213347-5daff5a && echo true || echo false
	# true
fi

echo "$debVersion" "$origVersion"
