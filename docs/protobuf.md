# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [task.proto](#task-proto)
    - [Task](#api-Task)
    - [Task.Info](#api-Task-Info)
    - [TaskCreateRequest](#api-TaskCreateRequest)
    - [TaskDeleteRequest](#api-TaskDeleteRequest)
    - [TaskGetRequest](#api-TaskGetRequest)
    - [TaskListRequest](#api-TaskListRequest)
    - [TaskListResponse](#api-TaskListResponse)
  
    - [Task.Info.Priority](#api-Task-Info-Priority)
    - [Task.Info.Status](#api-Task-Info-Status)
  
    - [TaskService](#api-TaskService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="task-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## task.proto



<a name="api-Task"></a>

### Task



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| info | [Task.Info](#api-Task-Info) |  |  |






<a name="api-Task-Info"></a>

### Task.Info



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| title | [string](#string) |  |  |
| description | [string](#string) |  |  |
| limit | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| priority | [Task.Info.Priority](#api-Task-Info-Priority) |  |  |
| status | [Task.Info.Status](#api-Task-Info-Status) |  |  |
| labels | [string](#string) | repeated |  |






<a name="api-TaskCreateRequest"></a>

### TaskCreateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| info | [Task.Info](#api-Task-Info) |  |  |






<a name="api-TaskDeleteRequest"></a>

### TaskDeleteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="api-TaskGetRequest"></a>

### TaskGetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="api-TaskListRequest"></a>

### TaskListRequest







<a name="api-TaskListResponse"></a>

### TaskListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tasks | [Task](#api-Task) | repeated |  |





 


<a name="api-Task-Info-Priority"></a>

### Task.Info.Priority


| Name | Number | Description |
| ---- | ------ | ----------- |
| PRIORITY_UNSPECIFIED | 0 |  |
| PRIORITY_LOW | 1 |  |
| PRIORITY_MIDDLE | 2 |  |
| PRIORITY_HIGH | 3 |  |



<a name="api-Task-Info-Status"></a>

### Task.Info.Status


| Name | Number | Description |
| ---- | ------ | ----------- |
| STATUS_UNSPECIFIED | 0 |  |
| STATUS_TODO | 1 |  |
| STATUS_DOING | 2 |  |
| STATUS_DONE | 3 |  |


 

 


<a name="api-TaskService"></a>

### TaskService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| List | [TaskListRequest](#api-TaskListRequest) | [TaskListResponse](#api-TaskListResponse) |  |
| Get | [TaskGetRequest](#api-TaskGetRequest) | [Task](#api-Task) |  |
| Create | [TaskCreateRequest](#api-TaskCreateRequest) | [TaskListResponse](#api-TaskListResponse) |  |
| Update | [Task](#api-Task) | [TaskListResponse](#api-TaskListResponse) |  |
| Delete | [TaskDeleteRequest](#api-TaskDeleteRequest) | [TaskListResponse](#api-TaskListResponse) |  |

 



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

