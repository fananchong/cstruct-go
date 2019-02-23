#!/bin/bash

set -ex

CUR_DIR=$PWD
cd ./benchmarks
go test -test.bench=".*" -count=1 cstrucgo_test.go myproto1.pb.go myproto2.pb.go
cd $CUR_DIR

