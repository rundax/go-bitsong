#!/usr/bin/env bash

set -e
set -u
set -o pipefail


if [[ $1 == "bash" ]]; then
  exec /bin/bash
  exit;
fi

###
### Globals
###

# The following global variables are available by our Dockerfile itself:
#   MY_USER
#   MY_GROUP
#   MY_UID
#   MY_GID

# Path to scripts to source
CONFIG_DIR="/docker-entrypoint.d"


###
### Source libs
###
init="$( find "${CONFIG_DIR}" -name '*.sh' -type f | sort -u )"
for f in ${init}; do
	# shellcheck disable=SC1090
	echo "running $f"
	. "${f}"
done


echo
echo 'Init process done. Ready for start up.'
echo

exec "${@}"
