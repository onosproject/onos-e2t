# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/xapp/e2t.proto](#api/xapp/e2t.proto)
    - [E2Control](#onos.e2t.xapp.E2Control)
    - [E2Message](#onos.e2t.xapp.E2Message)
  
    - [E2TService](#onos.e2t.xapp.E2TService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api/xapp/e2t.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/xapp/e2t.proto



<a name="onos.e2t.xapp.E2Control"></a>

### E2Control
Request encoding format (ASN.1 or Protobuf)
Add subscriptions
Remove subscriptions
Send control/insert/policy/query messages to specific device






<a name="onos.e2t.xapp.E2Message"></a>

### E2Message



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| e2node | [string](#string) |  | ID of E2 node that sent the message |
| service_model | [string](#string) |  | Service model ID |
| payload | [bytes](#bytes) |  | Message data (encoded as ASN.1 or Protobuf) |





 

 

 


<a name="onos.e2t.xapp.E2TService"></a>

### E2TService
E2TService provides means for enhanced interactions with the ONOS RIC E2 Termination service.

List of registered/available SMs

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| RegisterApp | [E2Control](#onos.e2t.xapp.E2Control) stream | [E2Message](#onos.e2t.xapp.E2Message) stream | RegisterApp establishes a bi-directional stream for conducting interactions with the E2 nodes in the RAN environment. |

 



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

