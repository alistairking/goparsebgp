#!/usr/bin/env bash
#
# Builds embedded libparsebgp as a static library
set -e

if [[ ! -f "./libparsebgp/lib/parsebgp.c" ]]; then
    echo "ERROR: cannot find embedded libparsebgp"
    exit 1
fi

SRCDIR=$(pwd)

mkdir -p build
cd ./libparsebgp

./autogen.sh

./configure \
    --prefix=${SRCDIR}/build \
    --disable-shared \
    --enable-static

make

make install

