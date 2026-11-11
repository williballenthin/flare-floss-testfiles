#!/bin/bash
# build.sh GO OUT GOOS GOARCH: obfuscate main.go with gobfuscate (GOPATH mode), then build the result with GO.
# gobfuscate (github.com/unixpickle/gobfuscate@13796335) must be built with Go 1.19 or older,
# whose math/rand is seeded the same each run, so its masks (and this build) are reproducible.
set -e
GO=$1; OUT=$(realpath -m $2); export GOOS=$3 GOARCH=$4
GOBFUSCATE=${GOBFUSCATE:-gobfuscate}
T=$(mktemp -d); trap "rm -rf $T" EXIT
mkdir -p $T/in/src/flosstest && cp main.go $T/in/src/flosstest/
GO111MODULE=off GOPATH=$T/in PATH=$(dirname $GO):$PATH $GOBFUSCATE -outdir -padding=flare-floss flosstest $T/out
cd $T/out/src/*/ && GO111MODULE=off GOPATH=$T/out CGO_ENABLED=0 $GO build -trimpath -ldflags="-s -w -buildid=" -o $OUT .
