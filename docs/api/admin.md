# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/openapi/e2/v1beta1/openapi_e2.proto](#api/openapi/e2/v1beta1/openapi_e2.proto)
    - [AppRequest](#openapi.e2.v1beta1.AppRequest)
    - [AppResponse](#openapi.e2.v1beta1.AppResponse)
    - [Indication](#openapi.e2.v1beta1.Indication)
    - [SubscribeDeleteRequest](#openapi.e2.v1beta1.SubscribeDeleteRequest)
    - [SubscribeDeleteResponse](#openapi.e2.v1beta1.SubscribeDeleteResponse)
    - [SubscribeRequest](#openapi.e2.v1beta1.SubscribeRequest)
    - [SubscribeResponse](#openapi.e2.v1beta1.SubscribeResponse)
  
    - [ResponseStatus](#openapi.e2.v1beta1.ResponseStatus)
  
    - [E2TService](#openapi.e2.v1beta1.E2TService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api/openapi/e2/v1beta1/openapi_e2.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/openapi/e2/v1beta1/openapi_e2.proto



<a name="openapi.e2.v1beta1.AppRequest"></a>

### AppRequest
AppRequest


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [RequestHeader](#openapi.e2.v1beta1.RequestHeader) |  |  |
| sub_req | [SubscribeRequest](#openapi.e2.v1beta1.SubscribeRequest) |  |  |
| sub_del_req | [SubscribeDeleteRequest](#openapi.e2.v1beta1.SubscribeDeleteRequest) |  |  |
| payload | [bytes](#bytes) |  |  |






<a name="openapi.e2.v1beta1.AppResponse"></a>

### AppResponse
AppResponse


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [RequestHeader](#openapi.e2.v1beta1.RequestHeader) |  |  |
| sub_resp | [SubscribeResponse](#openapi.e2.v1beta1.SubscribeResponse) |  |  |
| sub_del_resp | [SubscribeDeleteResponse](#openapi.e2.v1beta1.SubscribeDeleteResponse) |  |  |
| indication | [Indication](#openapi.e2.v1beta1.Indication) |  |  |
| payload | [bytes](#bytes) |  |  |






<a name="openapi.e2.v1beta1.Indication"></a>

### Indication
Indication an indication message






<a name="openapi.e2.v1beta1.SubscribeDeleteRequest"></a>

### SubscribeDeleteRequest
SubscribeDeleteRequest a subscription delete request






<a name="openapi.e2.v1beta1.SubscribeDeleteResponse"></a>

### SubscribeDeleteResponse
SubscribeDeleteResponse a subscription delete response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [ResponseStatus](#openapi.e2.v1beta1.ResponseStatus) |  |  |






<a name="openapi.e2.v1beta1.SubscribeRequest"></a>

### SubscribeRequest
SubscribeRequest a subscription request






<a name="openapi.e2.v1beta1.SubscribeResponse"></a>

### SubscribeResponse
SubscribeResponse a subscription response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [ResponseStatus](#openapi.e2.v1beta1.ResponseStatus) |  |  |





 


<a name="openapi.e2.v1beta1.ResponseStatus"></a>

### ResponseStatus
ResponseStatus

| Name | Number | Description |
| ---- | ------ | ----------- |
| RESPONSE_STATUS_FAILED | 0 |  |
| RESPONSE_STATUS_SUCCESSFUL | 1 |  |


 

 


<a name="openapi.e2.v1beta1.E2TService"></a>

### E2TService
E2TService provides means for enhanced interactions with the ONOS RIC E2 Termination service.

List of registered/available SMs

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| RegisterApp | [AppRequest](#openapi.e2.v1beta1.AppRequest) stream | [AppResponse](#openapi.e2.v1beta1.AppResponse) stream | RegisterApp establishes a bi-directional stream for conducting interactions with the E2 nodes in the RAN environment. |

 



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

