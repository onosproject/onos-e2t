# SPDX-License-Identifier: Apache-2.0
# Copyright 2019 Open Networking Foundation
# Copyright 2024 Intel Corporation

export CGO_ENABLED=1
export GO111MODULE=on

.PHONY: build

ONOS_E2T_VERSION ?= latest
ONOS_PROTOC_VERSION := v1.0.2
BUF_VERSION := 1.0.0

GOLANG_CI_VERSION := v1.52.2

all: build docker-build

build: # @HELP build the Go binaries and run all validations (default)
	go build -o build/_output/onos-e2t ./cmd/onos-e2t

test: # @HELP run the unit tests and source code validation producing a golang style report
test: build lint license
	go test -race github.com/onosproject/onos-e2t/...

proto-preliminaries: #@HELP preliminary command set for proto build
	rm -rf build/_input/onos-lib-go
	git clone https://github.com/onosproject/onos-lib-go.git build/_input/onos-lib-go

buflint: proto-preliminaries #@HELP run the "buf check lint" command on the proto files in 'api'
	docker run -it -v `pwd`:/go/src/github.com/onosproject/onos-e2t \
		-v `pwd`/build/_input/onos-lib-go/api/asn1:/go/src/github.com/onosproject/onos-e2t/api/asn1 \
		-w /go/src/github.com/onosproject/onos-e2t/api \
		bufbuild/buf:${BUF_VERSION} lint --path e2ap

protos: # @HELP compile the protobuf files (using protoc-go Docker)
protos: buflint
	docker run -it -v `pwd`:/go/src/github.com/onosproject/onos-e2t \
		-v `pwd`/build/_input/onos-lib-go:/go/src/github.com/onosproject/onos-lib-go \
		-w /go/src/github.com/onosproject/onos-e2t \
		--entrypoint build/bin/compile-protos.sh \
		onosproject/protoc-go:${ONOS_PROTOC_VERSION}

docker-build-onos-e2t: # @HELP build onos-e2t Docker image
	@go mod vendor
	docker build . -f build/onos-e2t/Dockerfile \
		-t onosproject/onos-e2t:${ONOS_E2T_VERSION}
	@rm -r vendor

docker-build: # @HELP build all Docker images
docker-build: build docker-build-onos-e2t

docker-push-onos-e2t: # @HELP push onos-e2t Docker image
	docker push onosproject/onos-e2t:${ONOS_E2T_VERSION}

docker-push: # @HELP push docker images
docker-push: docker-push-onos-e2t

lint: # @HELP examines Go source code and reports coding problems
	golangci-lint --version | grep $(GOLANG_CI_VERSION) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b `go env GOPATH`/bin $(GOLANG_CI_VERSION)
	golangci-lint run --timeout 15m

license: # @HELP run license checks
	rm -rf venv
	python3 -m venv venv
	. ./venv/bin/activate;\
	python3 -m pip install --upgrade pip;\
	python3 -m pip install reuse;\
	reuse lint

check-version: # @HELP check version is duplicated
	./build/bin/version_check.sh all

clean:: # @HELP remove all the build artifacts
	rm -rf ./build/_output ./vendor ./cmd/onos-e2t/onos-e2t ./cmd/onos/onos ./build/_input ./api/asn1/*
	go clean github.com/onosproject/onos-e2t/...

help:
	@grep -E '^.*: *# *@HELP' $(MAKEFILE_LIST) \
    | sort \
    | awk ' \
        BEGIN {FS = ": *# *@HELP"}; \
        {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}; \
    '
