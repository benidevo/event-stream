#!/bin/sh

set -e # Exit on failure

go build -o /tmp/codecrafters-build-kafka-go app/*.go
