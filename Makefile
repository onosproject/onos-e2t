# SPDX-FileCopyrightText: 2020 Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0

export CGO_ENABLED=1
export GO111MODULE=on

.PHONY: build

ONOS_E2T_VERSION := latest
ONOS_PROTOC_VERSION := v1.0.2
BUF_VERSION := 1.0.0

build: # @HELP build the Go binaries and run all validations (default)
build:
	go build -o build/_output/onos-e2t ./cmd/onos-e2t

build-tools:=$(shell if [ ! -d "./build/build-tools" ]; then cd build && git clone https://github.com/onosproject/build-tools.git; fi)
include ./build/build-tools/make/onf-common.mk

sim-app: # @HELP build the Go binaries for the simulator
sim-app:
	CGO_ENABLED=0 go build -o build/_output/onos-e2t-sim-app ./cmd/onos-e2t-sim-app

test: # @HELP run the unit tests and source code validation producing a golang style report
test: build linters license
	go test -race github.com/onosproject/onos-e2t/...

jenkins-test:  # @HELP run the unit tests and source code validation producing a junit style report for Jenkins
jenkins-test: deps license linters
	GODEBUG=cgocheck=0 TEST_PACKAGES=github.com/onosproject/onos-e2t/... ./build/build-tools/build/jenkins/make-unit

coverage: # @HELP generate unit test coverage data
coverage: build deps linters
	./build/bin/coveralls-coverage

buflint: #@HELP run the "buf check lint" command on the proto files in 'api'
	docker run -it -v `pwd`:/go/src/github.com/onosproject/onos-e2t \
		-v `pwd`/../onos-lib-go/api/asn1:/go/src/github.com/onosproject/onos-e2t/api/asn1 \
		-w /go/src/github.com/onosproject/onos-e2t/api \
		bufbuild/buf:${BUF_VERSION} lint --path e2ap

protos: # @HELP compile the protobuf files (using protoc-go Docker)
protos: buflint
	docker run -it -v `pwd`:/go/src/github.com/onosproject/onos-e2t \
		-v `pwd`/../onos-lib-go:/go/src/github.com/onosproject/onos-lib-go \
		-w /go/src/github.com/onosproject/onos-e2t \
		--entrypoint build/bin/compile-protos.sh \
		onosproject/protoc-go:${ONOS_PROTOC_VERSION}

onos-e2t-docker: # @HELP build onos-e2t Docker image
onos-e2t-docker:
	@go mod vendor
	docker build . -f build/onos-e2t/Dockerfile \
		-t onosproject/onos-e2t:${ONOS_E2T_VERSION}
	@rm -r vendor

onos-e2t-sim-app-docker: # @HELP build onos-e2t-sim-app Docker image
onos-e2t-sim-app-docker:
	@go mod vendor
	docker build . -f build/onos-e2t-sim-app/Dockerfile \
		-t onosproject/onos-e2t-sim-app:${ONOS_E2T_VERSION}
	@rm -r vendor

images: # @HELP build all Docker images
images: build onos-e2t-docker

kind: # @HELP build Docker images and add them to the currently configured kind cluster
kind: images
	@if [ "`kind get clusters`" = '' ]; then echo "no kind cluster found" && exit 1; fi
	kind load docker-image onosproject/onos-e2t:${ONOS_E2T_VERSION}

all: build images

integration-tests: # @HELP run helmit integration tests
integration-tests: integration-test-namespace
	helmit test -n test ./cmd/onos-e2t-tests --timeout 30m --no-teardown \
		--set sd-ran.onos-topo.logging.loggers.root.level=debug \
		--set sd-ran.onos-e2t.logging.loggers.root.level=debug \
		--set sd-ran.onos-e2t.logging.loggers.onos/grpc/retry.level=error \
		--set sd-ran.onos-e2t.replicaCount=2 \
		--suite e2

publish: # @HELP publish version on github and dockerhub
	./build/build-tools/publish-version ${VERSION} onosproject/onos-e2t

jenkins-publish: jenkins-tools # @HELP Jenkins calls this to publish artifacts
	./build/bin/push-images
	./build/build-tools/release-merge-commit

clean:: # @HELP remove all the build artifacts
	rm -rf ./build/_output ./vendor ./cmd/onos-e2t/onos-e2t ./cmd/onos/onos
	go clean -testcache github.com/onosproject/onos-e2t/...

