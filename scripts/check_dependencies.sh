#!/usr/bin/env bash
# https://kvz.io/blog/2013/11/21/bash-best-practices/
set -o xtrace   # set -x, debugging, trace what gets executed

set -o errexit  # set -e, exit when a command fails
set -o nounset  # set -u, exit when tries to use undeclared variables.
set -o pipefail # exit when last command that threw a non-zero exit code is returned

# From https://github.com/Boostport/kubernetes-vault/blob/master/deployments/quick-start/setup.sh#L3
# Check script dependencies
for command in jq kubectl vault
do
    if ! hash $command 2>/dev/null; then
        echo $command is not available, please install $command
        exit 1
    fi
done
