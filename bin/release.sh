#!/usr/bin/env bash

set -euo pipefail

if [ "$#" -lt 1 ]; then
  echo
  echo "usage: release.sh <ImageTag>"
  echo "  ie. release.sh v0.1.30"
  exit 255
fi

APP_VERSION=${1}
APP_VERSION=${APP_VERSION} docker-compose build release
APP_VERSION=${APP_VERSION} docker-compose push release