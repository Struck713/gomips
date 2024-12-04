#!/bin/bash

GOOS=js GOARCH=wasm go build -o main.wasm
python3 -m http.server
