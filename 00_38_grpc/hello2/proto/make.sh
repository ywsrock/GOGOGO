#!/bin/bash

cd `pwd`
printf "%s\n" $1

protoc --go_out=`pwd` --go_opt=paths=source_relative --go-grpc_out=`pwd` --go-grpc_opt=paths=source_relative $1
