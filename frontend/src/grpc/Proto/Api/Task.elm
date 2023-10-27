{- !!! DO NOT EDIT THIS FILE MANUALLY !!! -}

module Proto.Api.Task exposing (Task, TaskAddRequest, TaskDeleteRequest, TaskGetRequest, TaskListRequest, TaskListResponse, decodeTask, decodeTaskAddRequest, decodeTaskDeleteRequest, decodeTaskGetRequest, decodeTaskListRequest, decodeTaskListResponse, defaultTask, defaultTaskAddRequest, defaultTaskDeleteRequest, defaultTaskGetRequest, defaultTaskListRequest, defaultTaskListResponse, encodeTask, encodeTaskAddRequest, encodeTaskDeleteRequest, encodeTaskGetRequest, encodeTaskListRequest, encodeTaskListResponse, fieldNumbersTask, fieldNumbersTaskAddRequest, fieldNumbersTaskDeleteRequest, fieldNumbersTaskGetRequest, fieldNumbersTaskListRequest, fieldNumbersTaskListResponse)

{-| 
This file was automatically generated by
- [`protoc-gen-elm`](https://www.npmjs.com/package/protoc-gen-elm) 3.2.0
- `protoc` 4.24.4
- the following specification files: `task.proto`

To run it, add a dependency via `elm install` on [`elm-protocol-buffers`](https://package.elm-lang.org/packages/eriktim/elm-protocol-buffers/1.2.0) version latest or higher.


@docs Task, TaskAddRequest, TaskDeleteRequest, TaskGetRequest, TaskListRequest, TaskListResponse, decodeTask, decodeTaskAddRequest
@docs decodeTaskDeleteRequest, decodeTaskGetRequest, decodeTaskListRequest, decodeTaskListResponse, defaultTask, defaultTaskAddRequest
@docs defaultTaskDeleteRequest, defaultTaskGetRequest, defaultTaskListRequest, defaultTaskListResponse, encodeTask, encodeTaskAddRequest
@docs encodeTaskDeleteRequest, encodeTaskGetRequest, encodeTaskListRequest, encodeTaskListResponse, fieldNumbersTask
@docs fieldNumbersTaskAddRequest, fieldNumbersTaskDeleteRequest, fieldNumbersTaskGetRequest, fieldNumbersTaskListRequest
@docs fieldNumbersTaskListResponse
-}

import Proto.Api.Task.Internals_
import Protobuf.Decode
import Protobuf.Encode


{-| The field numbers for the fields of `Task`. This is mostly useful for internals, like documentation generation.


-}
fieldNumbersTask : { id : Int, info : Int }
fieldNumbersTask =
    Proto.Api.Task.Internals_.fieldNumbersProto__Api__Task__Task


{-| Default for Task. Should only be used for 'required' decoders as an initial value.


-}
defaultTask : Task
defaultTask =
    Proto.Api.Task.Internals_.defaultProto__Api__Task__Task


{-| Declares how to decode a `Task` from Bytes. To actually perform the conversion from Bytes, you need to use Protobuf.Decode.decode from eriktim/elm-protocol-buffers.


-}
decodeTask : Protobuf.Decode.Decoder Task
decodeTask =
    Proto.Api.Task.Internals_.decodeProto__Api__Task__Task


{-| Declares how to encode a `Task` to Bytes. To actually perform the conversion to Bytes, you need to use Protobuf.Encode.encode from eriktim/elm-protocol-buffers.


-}
encodeTask : Task -> Protobuf.Encode.Encoder
encodeTask =
    Proto.Api.Task.Internals_.encodeProto__Api__Task__Task


{-| `Task` message


-}
type alias Task =
    Proto.Api.Task.Internals_.Proto__Api__Task__Task


{-| The field numbers for the fields of `TaskDeleteRequest`. This is mostly useful for internals, like documentation generation.


-}
fieldNumbersTaskDeleteRequest : { id : Int }
fieldNumbersTaskDeleteRequest =
    Proto.Api.Task.Internals_.fieldNumbersProto__Api__Task__TaskDeleteRequest


{-| Default for TaskDeleteRequest. Should only be used for 'required' decoders as an initial value.


-}
defaultTaskDeleteRequest : TaskDeleteRequest
defaultTaskDeleteRequest =
    Proto.Api.Task.Internals_.defaultProto__Api__Task__TaskDeleteRequest


{-| Declares how to decode a `TaskDeleteRequest` from Bytes. To actually perform the conversion from Bytes, you need to use Protobuf.Decode.decode from eriktim/elm-protocol-buffers.


-}
decodeTaskDeleteRequest : Protobuf.Decode.Decoder TaskDeleteRequest
decodeTaskDeleteRequest =
    Proto.Api.Task.Internals_.decodeProto__Api__Task__TaskDeleteRequest


{-| Declares how to encode a `TaskDeleteRequest` to Bytes. To actually perform the conversion to Bytes, you need to use Protobuf.Encode.encode from eriktim/elm-protocol-buffers.


-}
encodeTaskDeleteRequest : TaskDeleteRequest -> Protobuf.Encode.Encoder
encodeTaskDeleteRequest =
    Proto.Api.Task.Internals_.encodeProto__Api__Task__TaskDeleteRequest


{-| `TaskDeleteRequest` message


-}
type alias TaskDeleteRequest =
    Proto.Api.Task.Internals_.Proto__Api__Task__TaskDeleteRequest


{-| The field numbers for the fields of `TaskAddRequest`. This is mostly useful for internals, like documentation generation.


-}
fieldNumbersTaskAddRequest : { info : Int }
fieldNumbersTaskAddRequest =
    Proto.Api.Task.Internals_.fieldNumbersProto__Api__Task__TaskAddRequest


{-| Default for TaskAddRequest. Should only be used for 'required' decoders as an initial value.


-}
defaultTaskAddRequest : TaskAddRequest
defaultTaskAddRequest =
    Proto.Api.Task.Internals_.defaultProto__Api__Task__TaskAddRequest


{-| Declares how to decode a `TaskAddRequest` from Bytes. To actually perform the conversion from Bytes, you need to use Protobuf.Decode.decode from eriktim/elm-protocol-buffers.


-}
decodeTaskAddRequest : Protobuf.Decode.Decoder TaskAddRequest
decodeTaskAddRequest =
    Proto.Api.Task.Internals_.decodeProto__Api__Task__TaskAddRequest


{-| Declares how to encode a `TaskAddRequest` to Bytes. To actually perform the conversion to Bytes, you need to use Protobuf.Encode.encode from eriktim/elm-protocol-buffers.


-}
encodeTaskAddRequest : TaskAddRequest -> Protobuf.Encode.Encoder
encodeTaskAddRequest =
    Proto.Api.Task.Internals_.encodeProto__Api__Task__TaskAddRequest


{-| `TaskAddRequest` message


-}
type alias TaskAddRequest =
    Proto.Api.Task.Internals_.Proto__Api__Task__TaskAddRequest


{-| The field numbers for the fields of `TaskGetRequest`. This is mostly useful for internals, like documentation generation.


-}
fieldNumbersTaskGetRequest : { id : Int }
fieldNumbersTaskGetRequest =
    Proto.Api.Task.Internals_.fieldNumbersProto__Api__Task__TaskGetRequest


{-| Default for TaskGetRequest. Should only be used for 'required' decoders as an initial value.


-}
defaultTaskGetRequest : TaskGetRequest
defaultTaskGetRequest =
    Proto.Api.Task.Internals_.defaultProto__Api__Task__TaskGetRequest


{-| Declares how to decode a `TaskGetRequest` from Bytes. To actually perform the conversion from Bytes, you need to use Protobuf.Decode.decode from eriktim/elm-protocol-buffers.


-}
decodeTaskGetRequest : Protobuf.Decode.Decoder TaskGetRequest
decodeTaskGetRequest =
    Proto.Api.Task.Internals_.decodeProto__Api__Task__TaskGetRequest


{-| Declares how to encode a `TaskGetRequest` to Bytes. To actually perform the conversion to Bytes, you need to use Protobuf.Encode.encode from eriktim/elm-protocol-buffers.


-}
encodeTaskGetRequest : TaskGetRequest -> Protobuf.Encode.Encoder
encodeTaskGetRequest =
    Proto.Api.Task.Internals_.encodeProto__Api__Task__TaskGetRequest


{-| `TaskGetRequest` message


-}
type alias TaskGetRequest =
    Proto.Api.Task.Internals_.Proto__Api__Task__TaskGetRequest


{-| The field numbers for the fields of `TaskListResponse`. This is mostly useful for internals, like documentation generation.


-}
fieldNumbersTaskListResponse : { tasks : Int }
fieldNumbersTaskListResponse =
    Proto.Api.Task.Internals_.fieldNumbersProto__Api__Task__TaskListResponse


{-| Default for TaskListResponse. Should only be used for 'required' decoders as an initial value.


-}
defaultTaskListResponse : TaskListResponse
defaultTaskListResponse =
    Proto.Api.Task.Internals_.defaultProto__Api__Task__TaskListResponse


{-| Declares how to decode a `TaskListResponse` from Bytes. To actually perform the conversion from Bytes, you need to use Protobuf.Decode.decode from eriktim/elm-protocol-buffers.


-}
decodeTaskListResponse : Protobuf.Decode.Decoder TaskListResponse
decodeTaskListResponse =
    Proto.Api.Task.Internals_.decodeProto__Api__Task__TaskListResponse


{-| Declares how to encode a `TaskListResponse` to Bytes. To actually perform the conversion to Bytes, you need to use Protobuf.Encode.encode from eriktim/elm-protocol-buffers.


-}
encodeTaskListResponse : TaskListResponse -> Protobuf.Encode.Encoder
encodeTaskListResponse =
    Proto.Api.Task.Internals_.encodeProto__Api__Task__TaskListResponse


{-| `TaskListResponse` message


-}
type alias TaskListResponse =
    Proto.Api.Task.Internals_.Proto__Api__Task__TaskListResponse


{-| The field numbers for the fields of `TaskListRequest`. This is mostly useful for internals, like documentation generation.


-}
fieldNumbersTaskListRequest : {}
fieldNumbersTaskListRequest =
    Proto.Api.Task.Internals_.fieldNumbersProto__Api__Task__TaskListRequest


{-| Default for TaskListRequest. Should only be used for 'required' decoders as an initial value.


-}
defaultTaskListRequest : TaskListRequest
defaultTaskListRequest =
    Proto.Api.Task.Internals_.defaultProto__Api__Task__TaskListRequest


{-| Declares how to decode a `TaskListRequest` from Bytes. To actually perform the conversion from Bytes, you need to use Protobuf.Decode.decode from eriktim/elm-protocol-buffers.


-}
decodeTaskListRequest : Protobuf.Decode.Decoder TaskListRequest
decodeTaskListRequest =
    Proto.Api.Task.Internals_.decodeProto__Api__Task__TaskListRequest


{-| Declares how to encode a `TaskListRequest` to Bytes. To actually perform the conversion to Bytes, you need to use Protobuf.Encode.encode from eriktim/elm-protocol-buffers.


-}
encodeTaskListRequest : TaskListRequest -> Protobuf.Encode.Encoder
encodeTaskListRequest =
    Proto.Api.Task.Internals_.encodeProto__Api__Task__TaskListRequest


{-| `TaskListRequest` message


-}
type alias TaskListRequest =
    Proto.Api.Task.Internals_.Proto__Api__Task__TaskListRequest
