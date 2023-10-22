# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [todo.proto](#todo-proto)
    - [ToDo](#api-todo-ToDo)
    - [ToDo.Info](#api-todo-ToDo-Info)
    - [ToDoAddRequest](#api-todo-ToDoAddRequest)
    - [ToDoDeleteRequest](#api-todo-ToDoDeleteRequest)
    - [ToDoGetRequest](#api-todo-ToDoGetRequest)
    - [ToDoListRequest](#api-todo-ToDoListRequest)
    - [ToDoListResponse](#api-todo-ToDoListResponse)
  
    - [ToDo.Info.Priority](#api-todo-ToDo-Info-Priority)
    - [ToDo.Info.Status](#api-todo-ToDo-Info-Status)
  
    - [ToDoService](#api-todo-ToDoService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="todo-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## todo.proto



<a name="api-todo-ToDo"></a>

### ToDo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| info | [ToDo.Info](#api-todo-ToDo-Info) |  |  |






<a name="api-todo-ToDo-Info"></a>

### ToDo.Info



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| title | [string](#string) |  |  |
| description | [string](#string) |  |  |
| limit | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| priority | [ToDo.Info.Priority](#api-todo-ToDo-Info-Priority) |  |  |
| status | [ToDo.Info.Status](#api-todo-ToDo-Info-Status) |  |  |
| labels | [string](#string) | repeated |  |






<a name="api-todo-ToDoAddRequest"></a>

### ToDoAddRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| info | [ToDo.Info](#api-todo-ToDo-Info) |  |  |






<a name="api-todo-ToDoDeleteRequest"></a>

### ToDoDeleteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |






<a name="api-todo-ToDoGetRequest"></a>

### ToDoGetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |






<a name="api-todo-ToDoListRequest"></a>

### ToDoListRequest







<a name="api-todo-ToDoListResponse"></a>

### ToDoListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| todos | [ToDo](#api-todo-ToDo) | repeated |  |





 


<a name="api-todo-ToDo-Info-Priority"></a>

### ToDo.Info.Priority


| Name | Number | Description |
| ---- | ------ | ----------- |
| PRIORITY_UNSPECIFIED | 0 |  |
| PRIORITY_LOW | 1 |  |
| PRIORITY_NORMAL_MIDDLE | 2 |  |
| PRIORITY_HIGH | 3 |  |



<a name="api-todo-ToDo-Info-Status"></a>

### ToDo.Info.Status


| Name | Number | Description |
| ---- | ------ | ----------- |
| STATUS_UNSPECIFIED | 0 |  |
| STATUS_NOT_YET | 1 |  |
| STATUS_DOING | 2 |  |
| STATUS_DONE | 3 |  |


 

 


<a name="api-todo-ToDoService"></a>

### ToDoService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| List | [ToDoListRequest](#api-todo-ToDoListRequest) | [ToDoListResponse](#api-todo-ToDoListResponse) |  |
| Get | [ToDoGetRequest](#api-todo-ToDoGetRequest) | [ToDo](#api-todo-ToDo) |  |
| Add | [ToDoAddRequest](#api-todo-ToDoAddRequest) | [ToDoListResponse](#api-todo-ToDoListResponse) |  |
| Update | [ToDo](#api-todo-ToDo) | [ToDoListResponse](#api-todo-ToDoListResponse) |  |
| Delete | [ToDoDeleteRequest](#api-todo-ToDoDeleteRequest) | [ToDoListResponse](#api-todo-ToDoListResponse) |  |

 



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

