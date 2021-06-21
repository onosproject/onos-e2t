module github.com/onosproject/onos-e2t

go 1.15

require (
	github.com/atomix/atomix-go-client v0.5.19
	github.com/atomix/atomix-go-framework v0.7.0
	github.com/cenkalti/backoff v2.2.1+incompatible
	github.com/cenkalti/backoff/v4 v4.0.0
	github.com/envoyproxy/protoc-gen-validate v0.4.1
	github.com/gogo/protobuf v1.3.2
	github.com/golang/mock v1.4.4
	github.com/golang/protobuf v1.5.0
	github.com/google/uuid v1.1.2
	github.com/onosproject/helmit v0.6.13
	github.com/onosproject/onos-api/go v0.7.62
	github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm v0.7.35
	github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2 v0.7.35
	github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre v0.7.35
	github.com/onosproject/onos-e2sub v0.7.3
	github.com/onosproject/onos-lib-go v0.7.13
	github.com/onosproject/onos-ric-sdk-go v0.7.18
	github.com/onosproject/onos-test v0.6.4
	github.com/stretchr/testify v1.7.0
	golang.org/x/net v0.0.0-20210510120150-4163338589ed
	google.golang.org/grpc v1.33.2
	google.golang.org/protobuf v1.26.0
	gotest.tools v2.2.0+incompatible
	k8s.io/api v0.21.0
	k8s.io/apimachinery v0.21.0
	k8s.io/utils v0.0.0-20201110183641-67b214c5f920
)

replace github.com/docker/docker => github.com/docker/engine v1.4.2-0.20200229013735-71373c6105e3
