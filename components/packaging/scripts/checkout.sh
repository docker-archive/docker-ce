#!/usr/bin/env sh

#   Copyright 2018-2020 Docker Inc.

#   Licensed under the Apache License, Version 2.0 (the "License");
#   you may not use this file except in compliance with the License.
#   You may obtain a copy of the License at

#       http://www.apache.org/licenses/LICENSE-2.0

#   Unless required by applicable law or agreed to in writing, software
#   distributed under the License is distributed on an "AS IS" BASIS,
#   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#   See the License for the specific language governing permissions and
#   limitations under the License.

checkout() (
	set -ex
	SRC="$1"
	REF="$2"
	REF_FETCH="$REF"
	# if ref is branch or tag, retrieve its canonical form
	REF=$(git -C "$SRC" ls-remote --refs --heads --tags origin "$REF" | awk '{print $2}')
	if [ -n "$REF" ]; then
		# if branch or tag then create it locally too
		REF_FETCH="$REF:$REF"
	else
		REF="FETCH_HEAD"
	fi
	git -C "$SRC" fetch --update-head-ok --depth 1 origin "$REF_FETCH"
	git -C "$SRC" checkout -q "$REF"
)


# Only execute checkout function above if this file is executed, not sourced from another script
prog=checkout.sh # needs to be in sync with this file's name
if [ "$(basename -- $0)" = "$prog" ]; then
	checkout $*
fi
