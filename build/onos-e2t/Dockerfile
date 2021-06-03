FROM onosproject/golang-build:v0.6.7 as build

ENV GO111MODULE=on
ARG ONOS_MAKE_TARGET=build

COPY Makefile go.mod go.sum /build/
COPY cmd/ /build/cmd/
COPY pkg/ /build/pkg/
COPY vendor/ /build/vendor/
COPY api/ /build/api

WORKDIR /build

RUN GOFLAGS=-mod=vendor make ${ONOS_MAKE_TARGET}

# So that dynamically linked modules that depend on glibc can be loaded i.e. the service models
FROM frolvlad/alpine-glibc:alpine-3.12

USER nobody

COPY --from=build /build/build/_output/onos-e2t /usr/local/bin/onos-e2t

ENTRYPOINT ["onos-e2t"]
