#!/bin/bash

set -x

CUR_DIR=$(pwd)
SRC_DIR=$CUR_DIR/src
OUTPUT_DIR=$SRC_DIR
mkdir -p $OUTPUT_DIR
protoc -I=$SRC_DIR --go_out=$OUTPUT_DIR --go-grpc_out=$OUTPUT_DIR $SRC_DIR/interfaces/protobuf_v3/ledger.proto
