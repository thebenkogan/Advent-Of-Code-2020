#!/usr/bin/zsh

go build -o ./build/$1-$2 ./cmd/$1/$2.go && ./build/$1-$2