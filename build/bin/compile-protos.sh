#!/bin/sh

proto_imports=".:${GOPATH}/src/github.com/gogo/protobuf/protobuf:${GOPATH}/src/github.com/gogo/protobuf:${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate:${GOPATH}/src":"${GOPATH}/src/github.com/onosproject/onos-e2t/api"

go_import_paths="Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types"
go_import_paths="${go_import_paths},Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types"
go_import_paths="${go_import_paths},Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types"

protoc -I=$proto_imports --validate_out=lang=go:. --proto_path=api --gogo_out=$go_import_paths,import_path=github.com/onosproject/onos-e2t/api:. api/e2ap/v1beta1/e2ap_commondatatypes.proto
protoc -I=$proto_imports --validate_out=lang=go:. --proto_path=api --gogo_out=$go_import_paths,import_path=github.com/onosproject/onos-e2t/api:. api/e2ap/v1beta1/e2ap_constants.proto
protoc -I=$proto_imports --validate_out=lang=go:. --proto_path=api --gogo_out=$go_import_paths,import_path=github.com/onosproject/onos-e2t/api:. api/e2ap/v1beta1/e2ap_containers.proto
protoc -I=$proto_imports --validate_out=lang=go:. --proto_path=api --gogo_out=$go_import_paths,import_path=github.com/onosproject/onos-e2t/api:. api/e2ap/v1beta1/e2ap_ies.proto
protoc -I=$proto_imports --validate_out=lang=go:. --proto_path=api --gogo_out=$go_import_paths,import_path=github.com/onosproject/onos-e2t/api:. api/e2ap/v1beta1/e2ap_pdu_contents.proto
protoc -I=$proto_imports --validate_out=lang=go:. --proto_path=api --gogo_out=$go_import_paths,import_path=github.com/onosproject/onos-e2t/api:. api/e2ap/v1beta1/e2ap_pdu_descriptions.proto

protoc -I=$proto_imports --validate_out=lang=go:. --proto_path=api --gogo_out=$go_import_paths,import_path=github.com/onosproject/onos-e2t/api:. api/e2ap/v1beta2/e2ap_commondatatypes.proto
protoc -I=$proto_imports --validate_out=lang=go:. --proto_path=api --gogo_out=$go_import_paths,import_path=github.com/onosproject/onos-e2t/api:. api/e2ap/v1beta2/e2ap_constants.proto
protoc -I=$proto_imports --validate_out=lang=go:. --proto_path=api --gogo_out=$go_import_paths,import_path=github.com/onosproject/onos-e2t/api:. api/e2ap/v1beta2/e2ap_containers.proto
protoc -I=$proto_imports --validate_out=lang=go:. --proto_path=api --gogo_out=$go_import_paths,import_path=github.com/onosproject/onos-e2t/api:. api/e2ap/v1beta2/e2ap_ies.proto
protoc -I=$proto_imports --validate_out=lang=go:. --proto_path=api --gogo_out=$go_import_paths,import_path=github.com/onosproject/onos-e2t/api:. api/e2ap/v1beta2/e2ap_pdu_contents.proto
protoc -I=$proto_imports --validate_out=lang=go:. --proto_path=api --gogo_out=$go_import_paths,import_path=github.com/onosproject/onos-e2t/api:. api/e2ap/v1beta2/e2ap_pdu_descriptions.proto

protoc -I=$proto_imports --gogofaster_out=$go_import_paths,import_path=github.com/onosproject/onos-e2t/sim/e2,plugins=grpc:. sim/e2/*.proto

cp -r github.com/onosproject/onos-e2t/* .
rm -rf github.com
