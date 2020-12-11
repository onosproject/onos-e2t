module github.com/onosproject/onos-e2t

go 1.15

require (
	github.com/envoyproxy/protoc-gen-validate v0.4.1
	github.com/gogo/protobuf v1.3.1
	github.com/golang/mock v1.4.4
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.1.2
	github.com/ishidawataru/sctp v0.0.0-20191218070446-00ab2ac2db07
	github.com/onosproject/helmit v0.6.8
	github.com/onosproject/onos-api/go v0.6.10
	github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm v0.6.8
	github.com/onosproject/onos-lib-go v0.6.25
	github.com/onosproject/onos-ric-sdk-go v0.6.9
	github.com/onosproject/onos-test v0.6.4
	github.com/spf13/cobra v1.1.1
	github.com/stretchr/testify v1.6.1
	google.golang.org/grpc v1.33.2
	google.golang.org/protobuf v1.25.0
	gotest.tools v2.2.0+incompatible
)

replace github.com/docker/docker => github.com/docker/engine v1.4.2-0.20200229013735-71373c6105e3
