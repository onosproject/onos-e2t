ARG ONOS_E2T_BASE_VERSION=latest

FROM onosproject/onos-e2t-base:$ONOS_E2T_BASE_VERSION as base

# So that dynamically linked modules that depend on glibc can be loaded i.e. the service models
FROM frolvlad/alpine-glibc:alpine-3.12

USER nobody

COPY --from=base /go/src/github.com/onosproject/onos-e2t/build/_output/onos-e2t /usr/local/bin/onos-e2t

ENTRYPOINT ["onos-e2t"]
