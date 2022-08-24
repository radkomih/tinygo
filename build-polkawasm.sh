#!/usr/bin/env bash

docker build --tag polkawasm/tinygo:0.25.0 -f Dockerfile.polkawasm .
docker run --rm -v $(pwd)/src:/src polkawasm/tinygo:0.25.0 tinygo build -target=polkawasm -o /src/examples/wasm/polkadot/dev_runtime.wasm examples/wasm/polkadot/ 