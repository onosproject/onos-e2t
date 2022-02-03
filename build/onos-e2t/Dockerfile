# SPDX-FileCopyrightText: 2022 2019-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0

FROM onosproject/golang-build:v1.0 as build


ENV GO111MODULE=on
ARG ONOS_MAKE_TARGET=build

COPY Makefile go.mod go.sum /build/
COPY cmd/ /build/cmd/
COPY pkg/ /build/pkg/
COPY vendor/ /build/vendor/
COPY api/ /build/api
COPY build/build-tools /build/build/build-tools/

WORKDIR /build

RUN GOFLAGS=-mod=vendor make ${ONOS_MAKE_TARGET}

# So that dynamically linked modules that depend on glibc can be loaded i.e. the service models
FROM frolvlad/alpine-glibc:alpine-3.12

USER nobody

COPY --from=build /build/build/_output/onos-e2t /usr/local/bin/onos-e2t

ENTRYPOINT ["onos-e2t"]
