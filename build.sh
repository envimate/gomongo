#!/bin/bash
export GOPATH=`pwd`

echo $GOPATH
go get -d -v gomongo
go install -v gomongo
