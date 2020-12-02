#!/usr/bin/env bash

r=$(dirname $(readlink -f $0))
r=$PWD

arg=$@
if [ -z $1 ]; then
    arg="bash"
fi

docker run -ti --rm -u 1000:1000 -v $r:/app -w /app pypy:3.7 pypy $arg
