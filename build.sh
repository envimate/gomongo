#!/bin/bash
export GOPATH=`pwd`

go get gopkg.in/mgo.v2
go install -v gomongo 
