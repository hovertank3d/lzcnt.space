#!/bin/bash

pushd deps/genhl
make
popd
go generate ./...
go build -o lzcnt.space