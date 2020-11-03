# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/ricapi/e2/subscription/v1beta1/subscription.proto](#api/ricapi/e2/subscription/v1beta1/subscription.proto)
    - [SubscribeDeleteRequest](#ricapi.e2.subscription.v1beta1.SubscribeDeleteRequest)
    - [SubscribeDeleteResponse](#ricapi.e2.subscription.v1beta1.SubscribeDeleteResponse)
    - [SubscribeRequest](#ricapi.e2.subscription.v1beta1.SubscribeRequest)
    - [SubscribeResponse](#ricapi.e2.subscription.v1beta1.SubscribeResponse)
  
    - [SubscriptionService](#ricapi.e2.subscription.v1beta1.SubscriptionService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api/ricapi/e2/subscription/v1beta1/subscription.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/ricapi/e2/subscription/v1beta1/subscription.proto



<a name="ricapi.e2.subscription.v1beta1.SubscribeDeleteRequest"></a>

### SubscribeDeleteRequest
SubscribeDeleteRequest a subscription delete request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [ricapi.e2.headers.v1beta1.RequestHeader](#ricapi.e2.headers.v1beta1.RequestHeader) |  |  |
| payload | [bytes](#bytes) |  |  |






<a name="ricapi.e2.subscription.v1beta1.SubscribeDeleteResponse"></a>

### SubscribeDeleteResponse
SubscribeDeleteResponse a subscription delete response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [ricapi.e2.headers.v1beta1.ResponseHeader](#ricapi.e2.headers.v1beta1.ResponseHeader) |  |  |
| payload | [bytes](#bytes) |  |  |






<a name="ricapi.e2.subscription.v1beta1.SubscribeRequest"></a>

### SubscribeRequest
SubscribeRequest a subscription request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [ricapi.e2.headers.v1beta1.RequestHeader](#ricapi.e2.headers.v1beta1.RequestHeader) |  |  |
| payload | [bytes](#bytes) |  |  |






<a name="ricapi.e2.subscription.v1beta1.SubscribeResponse"></a>

### SubscribeResponse
SubscribeResponse a subscription response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [ricapi.e2.headers.v1beta1.ResponseHeader](#ricapi.e2.headers.v1beta1.ResponseHeader) |  |  |
| payload | [bytes](#bytes) |  |  |





 

 

 


<a name="ricapi.e2.subscription.v1beta1.SubscriptionService"></a>

### SubscriptionService
SubscriptionService manage subscription and subscription delete requests

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| AddSubscription | [SubscribeRequest](#ricapi.e2.subscription.v1beta1.SubscribeRequest) | [SubscribeResponse](#ricapi.e2.subscription.v1beta1.SubscribeResponse) | AddSubscription establish E2 subscriptions on E2 Node. |
| DeleteSubscription | [SubscribeDeleteRequest](#ricapi.e2.subscription.v1beta1.SubscribeDeleteRequest) | [SubscribeDeleteResponse](#ricapi.e2.subscription.v1beta1.SubscribeDeleteResponse) | DeleteSubscription delete E2 subscriptions on E2 Node. |

 



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

