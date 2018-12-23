#!/bin/bash

set -ex

SRC_DIR=$PWD
docker run --rm -v $SRC_DIR:$SRC_DIR -w $SRC_DIR znly/protoc --go_out=. -I=. myproto1.proto
docker run --rm -v $SRC_DIR:$SRC_DIR -w $SRC_DIR znly/protoc --gogofaster_out=. -I=. myproto2.proto
