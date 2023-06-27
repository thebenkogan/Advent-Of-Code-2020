#!/usr/bin/zsh

go build -o ./build/$1 cmd/$1/*.go && ./build/$1 $1 $2