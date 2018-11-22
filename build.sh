#!/usr/bin/env bash

set -ex

IMAGE_BASE=matthias/alp-client

docker build -t ${IMAGE_BASE}:latest .
