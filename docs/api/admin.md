<!--
SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>

SPDX-License-Identifier: Apache-2.0
-->

# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/admin/v1/admin.proto](#api/admin/v1/admin.proto)
    - [DropE2NodeConnectionsRequest](#admin.v1.DropE2NodeConnectionsRequest)
    - [DropE2NodeConnectionsResponse](#admin.v1.DropE2NodeConnectionsResponse)
    - [ListE2NodeConnectionsRequest](#admin.v1.ListE2NodeConnectionsRequest)
    - [ListE2NodeConnectionsResponse](#admin.v1.ListE2NodeConnectionsResponse)
    - [ListRegisteredServiceModelsRequest](#admin.v1.ListRegisteredServiceModelsRequest)
    - [ListRegisteredServiceModelsResponse](#admin.v1.ListRegisteredServiceModelsResponse)
    - [UploadRegisterServiceModelRequest](#admin.v1.UploadRegisterServiceModelRequest)
    - [UploadRegisterServiceModelResponse](#admin.v1.UploadRegisterServiceModelResponse)
  
    - [E2NodeConnectionType](#admin.v1.E2NodeConnectionType)
  
    - [E2TAdminService](#admin.v1.E2TAdminService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api/admin/v1/admin.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/admin/v1/admin.proto



<a name="admin.v1.DropE2NodeConnectionsRequest"></a>

### DropE2NodeConnectionsRequest
DropE2NodeConnectionsRequest carries drop connection request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| connections | [ListE2NodeConnectionsResponse](#admin.v1.ListE2NodeConnectionsResponse) | repeated |  |






<a name="admin.v1.DropE2NodeConnectionsResponse"></a>

### DropE2NodeConnectionsResponse
DropE2NodeConnectionsResponse carries drop connection response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) | repeated |  |






<a name="admin.v1.ListE2NodeConnectionsRequest"></a>

### ListE2NodeConnectionsRequest
ListE2NodeConnectionsRequest carries request for a list of E2 node SCTP connections.






<a name="admin.v1.ListE2NodeConnectionsResponse"></a>

### ListE2NodeConnectionsResponse
ListE2NodeConnectionsResponse carries information about the SCTP connection to the remote E2 node.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| remote_ip | [string](#string) | repeated |  |
| remote_port | [uint32](#uint32) |  |  |
| id | [string](#string) |  |  |
| plmn_id | [string](#string) |  |  |
| connection_type | [E2NodeConnectionType](#admin.v1.E2NodeConnectionType) |  |  |






<a name="admin.v1.ListRegisteredServiceModelsRequest"></a>

### ListRegisteredServiceModelsRequest
ListRegisteredServiceModelsRequest carries data for querying registered service model plugins.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| model_name | [string](#string) |  | An optional filter on the name of the model plugins to list. |
| model_version | [string](#string) |  | An optional filter on the version of the model plugins to list |






<a name="admin.v1.ListRegisteredServiceModelsResponse"></a>

### ListRegisteredServiceModelsResponse
ListRegisteredServiceModelsResponse is general information about a service model plugin.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | name is the name given to the service model plugin - no spaces and title case. |
| version | [string](#string) |  | version is the semantic version of the Plugin e.g. 1.0.0. |






<a name="admin.v1.UploadRegisterServiceModelRequest"></a>

### UploadRegisterServiceModelRequest
UploadRegisterServiceModelRequest is for streaming a model plugin file to the server.
There is a built in limit in gRPC of 4MB - plugin is usually around 20MB
so break in to chunks of approx 1-2MB.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| so_file | [string](#string) |  | so_file is the name being streamed. |
| content | [bytes](#bytes) |  | content is the bytes content. |






<a name="admin.v1.UploadRegisterServiceModelResponse"></a>

### UploadRegisterServiceModelResponse
UploadRegisterServiceModelResponse carries status of model plugin registration.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | name is name of the model plugin. |
| version | [string](#string) |  | version is the semantic version of the model plugin. |





 


<a name="admin.v1.E2NodeConnectionType"></a>

### E2NodeConnectionType
E2NodeConnectionType specifies the type of an E2 connection

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN | 0 |  |
| G_NB | 1 |  |
| E_NB | 2 |  |
| ENG_MB | 3 |  |
| NGE_NB | 4 |  |


 

 


<a name="admin.v1.E2TAdminService"></a>

### E2TAdminService
E2TAdminService provides means for enhanced interactions with the ONOS RIC E2 Termination service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| UploadRegisterServiceModel | [UploadRegisterServiceModelRequest](#admin.v1.UploadRegisterServiceModelRequest) stream | [UploadRegisterServiceModelResponse](#admin.v1.UploadRegisterServiceModelResponse) | UploadRegisterServiceModel uploads and adds the model plugin to the list of supported models. The file is serialized in to Chunks of less than 4MB so as not to break the gRPC byte array limit |
| ListRegisteredServiceModels | [ListRegisteredServiceModelsRequest](#admin.v1.ListRegisteredServiceModelsRequest) | [ListRegisteredServiceModelsResponse](#admin.v1.ListRegisteredServiceModelsResponse) stream | ListRegisteredServiceModels returns a stream of registered service models. |
| ListE2NodeConnections | [ListE2NodeConnectionsRequest](#admin.v1.ListE2NodeConnectionsRequest) | [ListE2NodeConnectionsResponse](#admin.v1.ListE2NodeConnectionsResponse) stream | ListE2NodeConnections returns a stream of existing SCTP connections. |
| DropE2NodeConnections | [DropE2NodeConnectionsRequest](#admin.v1.DropE2NodeConnectionsRequest) | [DropE2NodeConnectionsResponse](#admin.v1.DropE2NodeConnectionsResponse) | DropE2NodeConnections drops the specified E2 node SCTP connections |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

