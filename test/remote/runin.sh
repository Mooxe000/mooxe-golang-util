#!/usr/bin/env bash

set -e

docker run \
  --name mooxe-golang-util \
  --rm \
  -ti \
  -v $(pwd):/root/mooxe-golang-util \
  -e GOPATH=/go:/root/mooxe-golang-util \
  mooxe/golang \
  /bin/bash
