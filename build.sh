#!/usr/bin/env bash

IMAGE_ID=$(docker build --quiet --tag radkomih/tinygo-dev:1.0 .)
docker run -it $IMAGE_ID /bin/bash
cd ../tinygo

touch dev_runtime.go
echo "package main; func main() {}" > dev_runtime.go

build/tinygo build -target=wasm -o=dev_runtime.wasm dev_runtime.go