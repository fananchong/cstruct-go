#!/bin/bash

set -ex

go test -v tests/x_test.go tests/nil_test.go tests/array_test.go tests/slice_ignore_nil_test.go tests/slice_struct_test.go
