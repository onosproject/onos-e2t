# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/ricapi/e2/subscription/v1beta1/subscription.proto](#api/ricapi/e2/subscription/v1beta1/subscription.proto)
    - [AddSubscriptionRequest](#ricapi.e2.subscription.v1beta1.AddSubscriptionRequest)
    - [AddSubscriptionResponse](#ricapi.e2.subscription.v1beta1.AddSubscriptionResponse)
    - [GetSubscriptionRequest](#ricapi.e2.subscription.v1beta1.GetSubscriptionRequest)
    - [GetSubscriptionResponse](#ricapi.e2.subscription.v1beta1.GetSubscriptionResponse)
    - [ListSubscriptionsRequest](#ricapi.e2.subscription.v1beta1.ListSubscriptionsRequest)
    - [ListSubscriptionsResponse](#ricapi.e2.subscription.v1beta1.ListSubscriptionsResponse)
    - [RemoveSubscriptionRequest](#ricapi.e2.subscription.v1beta1.RemoveSubscriptionRequest)
    - [RemoveSubscriptionResponse](#ricapi.e2.subscription.v1beta1.RemoveSubscriptionResponse)
    - [Subscription](#ricapi.e2.subscription.v1beta1.Subscription)
    - [UpdateSubscriptionRequest](#ricapi.e2.subscription.v1beta1.UpdateSubscriptionRequest)
    - [UpdateSubscriptionResponse](#ricapi.e2.subscription.v1beta1.UpdateSubscriptionResponse)
    - [WatchSubscriptionsRequest](#ricapi.e2.subscription.v1beta1.WatchSubscriptionsRequest)
    - [WatchSubscriptionsResponse](#ricapi.e2.subscription.v1beta1.WatchSubscriptionsResponse)
  
    - [EventType](#ricapi.e2.subscription.v1beta1.EventType)
    - [State](#ricapi.e2.subscription.v1beta1.State)
  
    - [SubscriptionService](#ricapi.e2.subscription.v1beta1.SubscriptionService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api/ricapi/e2/subscription/v1beta1/subscription.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/ricapi/e2/subscription/v1beta1/subscription.proto



<a name="ricapi.e2.subscription.v1beta1.AddSubscriptionRequest"></a>

### AddSubscriptionRequest
AddSubscriptionRequest a subscription request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subscription | [Subscription](#ricapi.e2.subscription.v1beta1.Subscription) |  |  |






<a name="ricapi.e2.subscription.v1beta1.AddSubscriptionResponse"></a>

### AddSubscriptionResponse
AddSubscriptionResponse a subscription response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subscription | [Subscription](#ricapi.e2.subscription.v1beta1.Subscription) |  |  |






<a name="ricapi.e2.subscription.v1beta1.GetSubscriptionRequest"></a>

### GetSubscriptionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |






<a name="ricapi.e2.subscription.v1beta1.GetSubscriptionResponse"></a>

### GetSubscriptionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subscription | [Subscription](#ricapi.e2.subscription.v1beta1.Subscription) |  |  |






<a name="ricapi.e2.subscription.v1beta1.ListSubscriptionsRequest"></a>

### ListSubscriptionsRequest







<a name="ricapi.e2.subscription.v1beta1.ListSubscriptionsResponse"></a>

### ListSubscriptionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subscriptions | [Subscription](#ricapi.e2.subscription.v1beta1.Subscription) | repeated |  |






<a name="ricapi.e2.subscription.v1beta1.RemoveSubscriptionRequest"></a>

### RemoveSubscriptionRequest
RemoveSubscriptionRequest a subscription delete request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subscription | [Subscription](#ricapi.e2.subscription.v1beta1.Subscription) |  |  |






<a name="ricapi.e2.subscription.v1beta1.RemoveSubscriptionResponse"></a>

### RemoveSubscriptionResponse
RemoveSubscriptionResponse a subscription delete response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subscription | [Subscription](#ricapi.e2.subscription.v1beta1.Subscription) |  |  |






<a name="ricapi.e2.subscription.v1beta1.Subscription"></a>

### Subscription
Subscription is a subscription state


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| revision | [uint64](#uint64) |  |  |
| app | [string](#string) |  |  |






<a name="ricapi.e2.subscription.v1beta1.UpdateSubscriptionRequest"></a>

### UpdateSubscriptionRequest
UpdateSubscriptionRequest is a subscription update request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subscription | [Subscription](#ricapi.e2.subscription.v1beta1.Subscription) |  |  |






<a name="ricapi.e2.subscription.v1beta1.UpdateSubscriptionResponse"></a>

### UpdateSubscriptionResponse
UpdateSubscriptionResponse is a subscription update response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subscription | [Subscription](#ricapi.e2.subscription.v1beta1.Subscription) |  |  |






<a name="ricapi.e2.subscription.v1beta1.WatchSubscriptionsRequest"></a>

### WatchSubscriptionsRequest







<a name="ricapi.e2.subscription.v1beta1.WatchSubscriptionsResponse"></a>

### WatchSubscriptionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [EventType](#ricapi.e2.subscription.v1beta1.EventType) |  |  |
| subscription | [Subscription](#ricapi.e2.subscription.v1beta1.Subscription) |  |  |





 


<a name="ricapi.e2.subscription.v1beta1.EventType"></a>

### EventType
EventType is a subscription event type

| Name | Number | Description |
| ---- | ------ | ----------- |
| NONE | 0 |  |
| ADDED | 1 |  |
| UPDATED | 2 |  |
| REMOVED | 3 |  |



<a name="ricapi.e2.subscription.v1beta1.State"></a>

### State
State is a subscription state

| Name | Number | Description |
| ---- | ------ | ----------- |
| INACTIVE | 0 | INACTIVE indicates the subscription is not currently active |
| ACTIVE | 1 | ACTIVE indicates the subscription is active |


 

 


<a name="ricapi.e2.subscription.v1beta1.SubscriptionService"></a>

### SubscriptionService
SubscriptionService manages subscription and subscription delete requests

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| AddSubscription | [AddSubscriptionRequest](#ricapi.e2.subscription.v1beta1.AddSubscriptionRequest) | [AddSubscriptionResponse](#ricapi.e2.subscription.v1beta1.AddSubscriptionResponse) | AddSubscription establishes E2 subscriptions on E2 Node. |
| UpdateSubscription | [UpdateSubscriptionRequest](#ricapi.e2.subscription.v1beta1.UpdateSubscriptionRequest) | [UpdateSubscriptionResponse](#ricapi.e2.subscription.v1beta1.UpdateSubscriptionResponse) | UpdateSubscription updates E2 subscriptions on E2 Node. |
| RemoveSubscription | [RemoveSubscriptionRequest](#ricapi.e2.subscription.v1beta1.RemoveSubscriptionRequest) | [RemoveSubscriptionResponse](#ricapi.e2.subscription.v1beta1.RemoveSubscriptionResponse) | RemoveSubscription removes E2 subscriptions on E2 Node. |
| GetSubscription | [GetSubscriptionRequest](#ricapi.e2.subscription.v1beta1.GetSubscriptionRequest) | [GetSubscriptionResponse](#ricapi.e2.subscription.v1beta1.GetSubscriptionResponse) | GetSubscription retrieves information about a specific subscription in the list of existing subscriptions |
| ListSubscriptions | [ListSubscriptionsRequest](#ricapi.e2.subscription.v1beta1.ListSubscriptionsRequest) | [ListSubscriptionsResponse](#ricapi.e2.subscription.v1beta1.ListSubscriptionsResponse) | ListSubscriptions returns the list of current existing subscriptions |
| WatchSubscriptions | [WatchSubscriptionsRequest](#ricapi.e2.subscription.v1beta1.WatchSubscriptionsRequest) | [WatchSubscriptionsResponse](#ricapi.e2.subscription.v1beta1.WatchSubscriptionsResponse) stream | WatchSubscriptions returns a stream of subscription changes |

 



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

