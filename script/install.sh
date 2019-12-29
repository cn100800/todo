#!/bin/bash

if command -v go 2>&1 >/dev/null; then
    go version
    exit 0
fi

go get -u github.com/freecracy/todo
