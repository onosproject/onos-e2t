module github.com/onosproject/onos-e2t

go 1.15

require (
	github.com/Microsoft/go-winio v0.4.15 // indirect
	github.com/docker/docker v1.13.1 // indirect
	github.com/envoyproxy/protoc-gen-validate v0.4.1
	github.com/garyburd/redigo v1.1.1-0.20170914051019-70e1b1943d4f // indirect
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.4.3
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/ishidawataru/sctp v0.0.0-20191218070446-00ab2ac2db07
	github.com/mattn/go-colorable v0.0.10-0.20170816031813-ad5389df28cd // indirect
	github.com/onosproject/helmit v0.6.8
	github.com/onosproject/onos-api/go v0.7.3
	github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm v0.7.6
	github.com/onosproject/onos-lib-go v0.7.0
	github.com/onosproject/onos-ric-sdk-go v0.7.7
	github.com/onosproject/onos-test v0.6.4
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/spf13/cobra v1.1.1 // indirect
	github.com/stretchr/testify v1.7.0
	google.golang.org/grpc v1.33.2
	google.golang.org/protobuf v1.25.0
	gotest.tools v2.2.0+incompatible
)

replace github.com/docker/docker => github.com/docker/engine v1.4.2-0.20200229013735-71373c6105e3
