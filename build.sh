#!/bin/bash

set -ex

CUR_DIR=$PWD
export GOBIN=$PWD/bin
go install -race .

cd ./example
go install -race main.go
cmake .

cd $CUR_DIR
