#!/bin/bash

set -e

MAIN_FILE=${1:-main.go}

OUTPUT="cmd/${2:-gateway}"

if [ ! -f "$MAIN_FILE" ]; then
  echo "Error: $MAIN_FILE not found."
  exit 1
fi

rm -f "$OUTPUT"

echo "Building $OUTPUT from $MAIN_FILE for Linux (GOOS=linux GOARCH=amd64)..."
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o  "$OUTPUT" "$MAIN_FILE"

chmod +x "$OUTPUT"

echo "Build complete: $OUTPUT"
