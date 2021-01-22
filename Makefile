export CGO_ENABLED=1
export GO111MODULE=on

.PHONY: build

ONOS_E2T_VERSION := latest
ONOS_BUILD_VERSION := v0.6.7
ONOS_PROTOC_VERSION := v0.6.7
BUF_VERSION := 0.36.0

build: # @HELP build the Go binaries and run all validations (default)
build:
	export GOPRIVATE="github.com/onosproject/onos-e2sub"
	go build -o build/_output/onos-e2t ./cmd/onos-e2t

test: # @HELP run the unit tests and source code validation
test: build deps linters license_check
	go test -race github.com/onosproject/onos-e2t/pkg/...
	go test -race github.com/onosproject/onos-e2t/cmd/...

coverage: # @HELP generate unit test coverage data
coverage: build deps linters license_check
	./build/bin/coveralls-coverage

deps: # @HELP ensure that the required dependencies are in place
	go build -v ./...
	bash -c "diff -u <(echo -n) <(git diff go.mod)"
	bash -c "diff -u <(echo -n) <(git diff go.sum)"

linters: # @HELP examines Go source code and reports coding problems
	golangci-lint run --timeout 30m

license_check: # @HELP examine and ensure license headers exist
	@if [ ! -d "../build-tools" ]; then cd .. && git clone https://github.com/onosproject/build-tools.git; fi
	./../build-tools/licensing/boilerplate.py -v --rootdir=${CURDIR} --boilerplate LicenseRef-ONF-Member-1.0

gofmt: # @HELP run the Go format validation
	bash -c "diff -u <(echo -n) <(gofmt -d pkg/ cmd/ tests/)"

buflint: #@HELP run the "buf check lint" command on the proto files in 'api'
	docker run -it -v `pwd`:/go/src/github.com/onosproject/onos-e2t \
		-w /go/src/github.com/onosproject/onos-e2t/api \
		bufbuild/buf:${BUF_VERSION} lint

protos: # @HELP compile the protobuf files (using protoc-go Docker)
protos: buflint
	docker run -it -v `pwd`:/go/src/github.com/onosproject/onos-e2t \
		-w /go/src/github.com/onosproject/onos-e2t \
		--entrypoint build/bin/compile-protos.sh \
		onosproject/protoc-go:${ONOS_PROTOC_VERSION}

onos-e2t-base-docker: # @HELP build onos-e2t base Docker image
	docker build . -f build/base/Dockerfile \
		--build-arg ONOS_BUILD_VERSION=${ONOS_BUILD_VERSION} \
		--build-arg ONOS_MAKE_TARGET=build \
		-t onosproject/onos-e2t-base:${ONOS_E2T_VERSION}

onos-e2t-docker: # @HELP build onos-e2t Docker image
onos-e2t-docker: onos-e2t-base-docker
	docker build . -f build/onos-e2t/Dockerfile \
		--build-arg ONOS_E2T_BASE_VERSION=${ONOS_E2T_VERSION} \
		-t onosproject/onos-e2t:${ONOS_E2T_VERSION}

images: # @HELP build all Docker images
images: build onos-e2t-docker

kind: # @HELP build Docker images and add them to the currently configured kind cluster
kind: images
	@if [ "`kind get clusters`" = '' ]; then echo "no kind cluster found" && exit 1; fi
	kind load docker-image onosproject/onos-e2t:${ONOS_E2T_VERSION}

all: build images

publish: # @HELP publish version on github and dockerhub
	./../build-tools/publish-version ${VERSION} onosproject/onos-e2t

bumponosdeps: # @HELP update "onosproject" go dependencies and push patch to git. Add a version to dependency to make it different to $VERSION
	./../build-tools/bump-onos-deps ${VERSION}

clean: # @HELP remove all the build artifacts
	rm -rf ./build/_output ./vendor ./cmd/onos-e2t/onos-e2t ./cmd/onos/onos
	go clean -testcache github.com/onosproject/onos-e2t/...

help:
	@grep -E '^.*: *# *@HELP' $(MAKEFILE_LIST) \
    | sort \
    | awk ' \
        BEGIN {FS = ": *# *@HELP"}; \
        {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}; \
    '
