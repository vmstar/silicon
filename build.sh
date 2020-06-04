#!/usr/bin/env bash
RUN_NAME=silicon

mkdir -p output/bin output/conf
export GO111MODULE="on"
go build -o output/bin/$RUN_NAME
chmod +x output/bin/$RUN_NAME
