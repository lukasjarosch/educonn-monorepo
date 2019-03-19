#!/usr/bin/env bash

# This script creates go stubs of all proto files. It will also generate the HTTP gateway as well as the swagger spec

PROTO_ROOT=$(pwd)/api/proto
SWAGGER_DIR=$(pwd)/api/swagger

for dir in $PROTO_ROOT/*/ ; do
    echo -n "=> '$(basename ${dir})' ... "

    protoc -I$GOPATH/src \
        -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
        -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
        --proto_path=${PROTO_ROOT} --go_out=plugins=grpc:$PROTO_ROOT $dir*.proto

    protoc -I$GOPATH/src \
        -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
        -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
        --proto_path=${PROTO_ROOT} --grpc-gateway_out=logtostderr=true:$PROTO_ROOT $dir*.proto

    protoc -I$GOPATH/src \
        -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
        -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
        --proto_path=${PROTO_ROOT} --swagger_out=logtostderr=true:$SWAGGER_DIR $dir*.proto
    echo "DONE!"
done
