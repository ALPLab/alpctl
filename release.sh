#!/usr/bin/env bash

set -ex

IMAGE_BASE=prodalpcrreg.azurecr.io/alplab-client-v2

# ensure we're up to date
# git pull

# bump version
docker run --rm -v "$PWD":/app treeder/bump patch
version=`cat VERSION`
echo "version: $version"

# tag it
git add -A
git commit -m "version $version"
git tag -a "$version" -m "version $version"
docker build . -t ${IMAGE_BASE}:latest
docker tag ${IMAGE_BASE}:latest ${IMAGE_BASE}:$version

# run container to generate binaries
# linux binary
docker run --rm -v "$PWD":/usr/src/alp -w /usr/src/alp ${IMAGE_BASE}:latest go build -v
# windows binary
docker run --rm -v "$PWD":/usr/src/alp -w /usr/src/alp -e GOOS=windows -e GOARCH=386 ${IMAGE_BASE}:latest go build -v

# check if binary works
./alp --help | grep ALP.Lab

# push it
git push
git push --tags
docker push ${IMAGE_BASE}:latest
docker push ${IMAGE_BASE}:$version


# docker push  prodalpcrreg.azurecr.io/alplab-client-v2
