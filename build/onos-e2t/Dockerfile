FROM onosproject/golang-build:v0.6.7 as build

ENV GO111MODULE=on
ARG ONOS_MAKE_TARGET=build

COPY Makefile go.mod go.sum /go/src/github.com/onosproject/onos-e2t/
COPY cmd/ /go/src/github.com/onosproject/onos-e2t/cmd/
COPY pkg/ /go/src/github.com/onosproject/onos-e2t/pkg/
COPY vendor/ /go/src/github.com/onosproject/onos-e2t/vendor/
COPY api/ /go/src/github.com/onosproject/onos-e2t/api

RUN cd /go/src/github.com/onosproject/onos-e2t && GOFLAGS=-mod=vendor make ${ONOS_MAKE_TARGET}

# So that dynamically linked modules that depend on glibc can be loaded i.e. the service models
FROM frolvlad/alpine-glibc:alpine-3.12

USER nobody

COPY --from=build /go/src/github.com/onosproject/onos-e2t/build/_output/onos-e2t /usr/local/bin/onos-e2t

ENTRYPOINT ["onos-e2t"]
