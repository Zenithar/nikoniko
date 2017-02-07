package dto

//go:generate protoc --proto_path=$GOPATH/src:. --gofast_out=plugins=grpc:. dto.proto
