#!/bin/bash

command -v go 2>&1 >/dev/null || {
    go version
    exit 0
}

go get -u github.com/freecracy/todo
