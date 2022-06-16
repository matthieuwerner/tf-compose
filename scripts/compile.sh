#!/bin/bash

for GOOS in darwin linux; do
   for GOARCH in 386 amd64; do
    export GOOS GOARCH
    go build -v -o build/tf-compose-$GOOS-$GOARCH src/*.go
  done
done