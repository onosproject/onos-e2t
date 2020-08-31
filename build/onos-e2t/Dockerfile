ARG ONOS_E2T_BASE_VERSION=latest

FROM onosproject/onos-e2t-base:$ONOS_E2T_BASE_VERSION as base

FROM alpine:3.11
RUN apk add libc6-compat

USER nobody

COPY --from=base /go/src/github.com/onosproject/onos-e2t/build/_output/onos-e2t /usr/local/bin/onos-e2t

ENTRYPOINT ["onos-e2t"]
