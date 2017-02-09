GO   := go
pkgs  = $(shell $(GO) list ./... | grep -v /vendor/)

all: format build

boostrap:
	@echo ">> installing dependencies"
	@$(GO) get -u -v github.com/golang/protobuf/protoc-gen-go
	@$(GO) get -u -v github.com/gogo/protobuf/protoc-gen-gofast
	@$(GO) get -u -v github.com/vektra/mockery/.../
	@$(GO) get -u -v google.golang.org/grpc
	@$(GO) get -u -v -insecure zenithar.org/go/common/...

generate:
	@echo ">> generate code"
	@$(GO) generate $(pkgs)

mocks: FORCE
	@echo ">> generating mocks"
	@rm -rf repositories/mock
	@mockery -output="./repositories/mock" -dir=$(GOPATH)/src/zenithar.org/go/nikoniko/repositories -all

test: mocks
	@echo ">> running tests"
	@$(GO) test -short $(pkgs)

format:
	@echo ">> formatting code"
	@$(GO) fmt $(pkgs)

vet:
	@echo ">> vetting code"
	@$(GO) vet $(pkgs)

build: generate
	@echo ">> building binaries"
	@./ci/build.sh

pack: build
	@echo ">> packing all binaries"
	@upx -9 bin/*

docker: pack
	@docker build -t nikoniko:$(shell cat version/VERSION)-$(shell git rev-parse --short HEAD) .

docker-release: docker
	@docker tag nikoniko:$(shell cat version/VERSION)-$(shell git rev-parse --short HEAD) nikoniko:$(shell cat version/VERSION)

FORCE:

.PHONY: all format build test vet docker assets generate mocks
