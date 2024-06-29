#!/bin/bash

DEPS_ROOT=bin/deps

mkdir -p ./"$DEPS_ROOT"

function get-dedoc {
    VERSION=0.2.5
    URL="https://github.com/toiletbril/dedoc/releases/download/$VERSION"

    for bin in dedoc dedoc.exe; do
        wget -P "$DEPS_ROOT" "$URL"/$bin
    done
}

function get