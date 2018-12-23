#!/bin/bash

set -ex

CUR_DIR=$PWD
cd ./benchmarks
go test -test.bench=".*" -count=1
cd $CUR_DIR

