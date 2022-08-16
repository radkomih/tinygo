#!/usr/bin/env bash

docker build --tag polkawasm/tinygo:0.25.0 -f Dockerfile.polkawasm .

docker run --rm -v $(pwd):/src polkawasm/tinygo:0.25.0 tinygo build -target=wasm -o=examples/wasm/polkadot.wasm examples/wasm/polkadot.go