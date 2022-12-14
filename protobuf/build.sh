#!/bin/bash

#!dir=$1

#!cd $dir
protoc \
   -I=. -I=/Users/hank/go/src -I=/Users/hank/go/src/protobuf/protobuf \
   --gogofaster_out=Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,plugins=grpc:. \
   *.proto