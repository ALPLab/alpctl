#!/usr/bin/env bash

set -ex

IMAGE_BASE=prodalpcrreg.azurecr.io/alplab-client-v2

docker build -t ${IMAGE_BASE}:latest .
