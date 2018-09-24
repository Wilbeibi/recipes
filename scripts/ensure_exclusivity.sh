#!/bin/bash

# https://github.com/spotify/docker-gc/commit/f5e8d3ab55fce92d0ee04b4bcb234660cac1515c

# for example process called docker-gc, it's pid file /var/run/docker-gc
PIDFILE="/var/run/dockergc"
# create/append file /var/run/dockergc, written to it via file descriptor 200 inside
# sub-shell, 200 is an arbitrary one
# Refer: https://stackoverflow.com/questions/21689891/what-does-200somefile-accomplish
exec 200 >>$PIDFILE
# -x obtain an exclusive lock, this is the default; -n fail (with an exit code of 1)
# rather than wait if the lock cannot be immediately acquired
if ! flock -x -n 200; then
    echo "[$(date)] : docker-gc : Process is already running"
    exit 1
fi
# http://redsymbol.net/articles/bash-exit-traps/
# https://unix.stackexchange.com/questions/11376/what-does-double-dash-mean-also-known-as-bare-double-dash
# a double dash (--) is used in bash built-in commands and many other commands to signify the end of command options
# for example, "grep -- -v file" search for literal "-v" string in file
# use it here to prevent weird file name like "-r what"
trap "rm -f -- '$PIDFILE'" EXIT
# $$ is the process ID
echo $$ > $PIDFILE