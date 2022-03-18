#!/bin/bash
set -ex
module="checkin-everything"
remote="tanglicai.xyz:5555"

# GO Build
echo "build app: $module"
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags "-s" -o bin/${module} .

echo "build docker image: $module"

tag="$(date '+%Y-%m-%d-%H-%M-%S')"
image="$module":"$tag"
remoteimage="$remote"/"$module"
docker build . -t "$image"
docker tag "$image" "$remoteimage"
docker push "$remoteimage"

rm -rf bin