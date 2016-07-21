#!/bin/sh

bin_file="./bin/story"
export GOROOT=/usr/local/go
export GOPATH=$GOPATH:/home/huhailang/develop/go/go-learn

go install -v github.com/dbdn/story
