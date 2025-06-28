#!/bin/bash
set -e

CGO_ENABLED=0 go build  -ldflags="-w -s" -o robovac-rooms