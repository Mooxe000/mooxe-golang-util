#!/usr/bin/env bash

# TODO use go

gopm clean
rm -rf .vendor/
rm -rf test/local/.vendor/
rm -rf test/remote/.vendor/
