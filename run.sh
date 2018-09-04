#!/usr/bin/env bash
set -euox pipefail

pushd idle
  go build
popd


pushd server
  go build
popd

pushd killer
  go test
popd
