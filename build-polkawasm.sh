#!/usr/bin/env bash

docker build --tag polkawasm/tinygo:0.25.0 -f Dockerfile.polkawasm .
docker run --rm -v $(pwd):/src polkawasm/tinygo:0.25.0 tinygo build -target=polkawasm -o=/src/examples/polkadot/polkadot.wasm examples/wasm/polkadot/polkadot.go


docker run -it -v $(pwd):/src polkawasm/tinygo:0.25.0 /bin/bash
cd ../tinygo
tinygo build -target=polkawasm -o=examples/wasm/polkadot.wasm examples/wasm/polkadot.go