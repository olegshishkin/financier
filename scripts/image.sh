#!/bin/sh
echo "Create container image"

while getopts e:i:p: flag; do
  case "${flag}" in
  e) engine=${OPTARG} ;;
  i) image_name=${OPTARG} ;;
  p) server_ports=${OPTARG} ;;
  *)
    echo "Unknown flag"
    exit 1
    ;;
  esac
done

echo "Container engine: $engine"
echo "Image name: $image_name"
echo "Expose ports: $server_ports"

latest_tag=$(git describe --tags)
echo "Image version: ${latest_tag}"
"$engine" build . -t "$image_name":"$latest_tag" --build-arg PORTS="$server_ports" -f ./build/package/Dockerfile
