#!/usr/bin/env bash

set -e

docker run \
  --name mooxe-golang-util \
  --rm \
  -ti \
  -v $(pwd):/root/mooxe-golang-util \
  mooxe/gopm \
  /bin/bash
