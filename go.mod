module github.com/onosproject/onos-e2t

go 1.15

require (
	github.com/atomix/atomix-go-client v0.5.16
	github.com/atomix/atomix-go-framework v0.6.16
	github.com/cenkalti/backoff/v4 v4.0.0
	github.com/envoyproxy/protoc-gen-validate v0.4.1
	github.com/gogo/protobuf v1.3.2
	github.com/golang/mock v1.4.4
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.1.2
	github.com/onosproject/helmit v0.6.12
	github.com/onosproject/onos-api/go v0.7.45
	github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm v0.7.35
	github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2 v0.7.35
	github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre v0.7.35
	github.com/onosproject/onos-e2sub v0.7.3
	github.com/onosproject/onos-lib-go v0.7.12
	github.com/onosproject/onos-ric-sdk-go v0.7.11
	github.com/onosproject/onos-test v0.6.4
	github.com/stretchr/testify v1.7.0
	golang.org/x/net v0.0.0-20210510120150-4163338589ed
	google.golang.org/grpc v1.33.2
	google.golang.org/protobuf v1.25.0
	gotest.tools v2.2.0+incompatible
)

replace github.com/docker/docker => github.com/docker/engine v1.4.2-0.20200229013735-71373c6105e3

replace github.com/onosproject/onos-api/go => github.com/kuujo/onos-api/go v0.0.0-20210531165102-2d153b71fbb9

replace github.com/onosproject/onos-ric-sdk-go => github.com/kuujo/onos-ric-sdk-go v0.7.6-0.20210531231823-59b9d267a8cb
