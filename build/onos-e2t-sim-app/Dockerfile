# SPDX-FileCopyrightText: 2022 2019-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0

FROM onosproject/golang-build:v1.3.1 as build

ENV GO111MODULE=on
ARG ONOS_MAKE_TARGET=build

COPY Makefile go.mod go.sum /build/
COPY cmd/ /build/cmd/
COPY pkg/ /build/pkg/
COPY vendor/ /build/vendor/
COPY sim/ /build/sim
COPY test/ /build/test

WORKDIR /build

RUN GOFLAGS=-mod=vendor make sim-app

# So that dynamically linked modules that depend on glibc can be loaded i.e. the service models
FROM frolvlad/alpine-glibc:alpine-3.19

USER nobody

COPY --from=build /build/build/_output/onos-e2t-sim-app /usr/local/bin/onos-e2t-sim-app
ADD build/onos-e2t-sim-app/logging.yaml logging.yaml

ENTRYPOINT ["onos-e2t-sim-app"]
