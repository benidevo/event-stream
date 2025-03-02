#!/bin/sh

set -e

PROJECT_ROOT="$(dirname "$(dirname "$0")")"

(
  cd "$PROJECT_ROOT"
  go build -o /tmp/codecrafters-build-kafka-go ./cmd/eventstream
)

exec /tmp/codecrafters-build-kafka-go "$@"
