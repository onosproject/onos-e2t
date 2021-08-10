export CGO_ENABLED=1
export GO111MODULE=on

.PHONY: build

ONOS_E2T_VERSION := latest
ONOS_PROTOC_VERSION := v0.6.7
BUF_VERSION := 0.36.0

build: # @HELP build the Go binaries and run all validations (default)
build:
	go build -o build/_output/onos-e2t ./cmd/onos-e2t

sim-app: # @HELP build the Go binaries for the simulator
sim-app:
	CGO_ENABLED=0 go build -o build/_output/onos-e2t-sim-app ./cmd/onos-e2t-sim-app

test: # @HELP run the unit tests and source code validation producing a golang style report
test: build deps linters license_check
	GODEBUG=cgocheck=0 go test -race github.com/onosproject/onos-e2t/...

jenkins-test:  # @HELP run the unit tests and source code validation producing a junit style report for Jenkins
jenkins-test: build-tools deps license_check linters
	GODEBUG=cgocheck=0 TEST_PACKAGES=github.com/onosproject/onos-e2t/... ./../build-tools/build/jenkins/make-unit

coverage: # @HELP generate unit test coverage data
coverage: build deps linters license_check
	./build/bin/coveralls-coverage

deps: # @HELP ensure that the required dependencies are in place
	go build -v ./...
	bash -c "diff -u <(echo -n) <(git diff go.mod)"
	bash -c "diff -u <(echo -n) <(git diff go.sum)"

linters: golang-ci # @HELP examines Go source code and reports coding problems
	golangci-lint run --timeout 10m

build-tools: # @HELP install the ONOS build tools if needed
	@if [ ! -d "../build-tools" ]; then cd .. && git clone https://github.com/onosproject/build-tools.git; fi

jenkins-tools: # @HELP installs tooling needed for Jenkins
	cd .. && go get -u github.com/jstemmer/go-junit-report && go get github.com/t-yuki/gocover-cobertura

golang-ci: # @HELP install golang-ci if not present
	golangci-lint --version || curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b `go env GOPATH`/bin v1.36.0

license_check: build-tools # @HELP examine and ensure license headers exist
	@if [ ! -d "../build-tools" ]; then cd .. && git clone https://github.com/onosproject/build-tools.git; fi
	./../build-tools/licensing/boilerplate.py -v --rootdir=${CURDIR} --boilerplate LicenseRef-ONF-Member-1.0

gofmt: # @HELP run the Go format validation
	bash -c "diff -u <(echo -n) <(gofmt -d pkg/ cmd/ tests/)"

buflint: #@HELP run the "buf check lint" command on the proto files in 'api'
	docker run -it -v `pwd`:/go/src/github.com/onosproject/onos-e2t \
		-w /go/src/github.com/onosproject/onos-e2t/api \
		bufbuild/buf:${BUF_VERSION} lint --path e2ap

protos: # @HELP compile the protobuf files (using protoc-go Docker)
protos: buflint
	docker run -it -v `pwd`:/go/src/github.com/onosproject/onos-e2t \
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
	kind load docker-image onosproject/onos-e2t-sim-app:${ONOS_E2T_VERSION}

all: build images

publish: # @HELP publish version on github and dockerhub
	./../build-tools/publish-version ${VERSION} onosproject/onos-e2t

jenkins-publish: build-tools jenkins-tools # @HELP Jenkins calls this to publish artifacts
	./build/bin/push-images
	../build-tools/release-merge-commit

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
