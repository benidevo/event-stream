#!/bin/sh

set -e # Exit early if any commands fail

(
  cd "$(dirname "$0")"
  go build -o /tmp/codecrafters-build-kafka-go app/*.go
)


exec /tmp/codecrafters-build-kafka-go "$@"
