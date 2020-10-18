#!/bin/bash

OUT_DIR=src/proto/

protoc \
  --js_out=import_style=commonjs,binary:$OUT_DIR \
  --grpc-web_out=import_style=typescript,mode=grpcwebtext:$OUT_DIR \
  -I ../server  ../server/*/*.proto
