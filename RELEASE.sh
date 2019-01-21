#!/usr/bin/env bash

IMAGE=prodalpcrreg.azurecr.io/alpctl

# ensure we're up to date
git pull
version=`git describe --tags`
echo "version: $version"

# Checkout Version
docker build . -t ${IMAGE}:latest
docker tag ${IMAGE}:latest ${IMAGE}:$version

# run container and copy  binaries
id=$(docker create ${IMAGE}:$version)
docker cp $id:/release /tmp/
docker rm -v $id

# check if binary works
docker run ${IMAGE}:$version | grep ALP.Lab

# push it
docker push ${IMAGE}:$version
