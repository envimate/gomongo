#!/bin/bash
export GOPATH=$(dirname $(readlink -f $0))

go get -d -v gomongo
go install -v gomongo
