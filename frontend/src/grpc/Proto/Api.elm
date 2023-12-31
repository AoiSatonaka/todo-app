{- !!! DO NOT EDIT THIS FILE MANUALLY !!! -}

module Proto.Api exposing (Task, TaskCreateRequest, TaskDeleteRequest, TaskGetRequest, TaskListRequest, TaskListResponse, decodeTask, decodeTaskCreateRequest, decodeTaskDeleteRequest, decodeTaskGetRequest, decodeTaskListRequest, decodeTaskListResponse, defaultTask, defaultTaskCreateRequest, defaultTaskDeleteRequest, defaultTaskGetRequest, defaultTaskListRequest, defaultTaskListResponse, encodeTask, encodeTaskCreateRequest, encodeTaskDeleteRequest, encodeTaskGetRequest, encodeTaskListRequest, encodeTaskListResponse, fieldNumbersTask, fieldNumbersTaskCreateRequest, fieldNumbersTaskDeleteRequest, fieldNumbersTaskGetRequest, fieldNumbersTaskListRequest, fieldNumbersTaskListResponse)

{-| 
This file was automatically generated by
- [`protoc-gen-elm`](https://www.npmjs.com/package/protoc-gen-elm) 3.2.0
- `protoc` 4.24.4
- the following specification files: `task.proto`

To run it, add a dependency via `elm install` on [`elm-protocol-buffers`](https://package.elm-lang.org/packages/eriktim/elm-protocol-buffers/1.2.0) version latest or higher.


@docs Task, TaskCreateRequest, TaskDeleteRequest, TaskGetRequest, TaskListRequest, TaskListResponse, decodeTask, decodeTaskCreateRequest
@docs decodeTaskDeleteRequest, decodeTaskGetRequest, decodeTaskListRequest, decodeTaskListResponse, defaultTask
@docs defaultTaskCreateRequest, defaultTaskDeleteRequest, defaultTaskGetRequest, defaultTaskListRequest, defaultTaskListResponse
@docs encodeTask, encodeTaskCreateRequest, encodeTaskDeleteRequest, encodeTaskGetRequest, encodeTaskListRequest, encodeTaskListResponse
@docs fieldNumbersTask, fieldNumbersTaskCreateRequest, fieldNumbersTaskDeleteRequest, fieldNumbersTaskGetRequest
@docs fieldNumbersTaskListRequest, fieldNumbersTaskListResponse
-}

import Proto.Api.Internals_
import Protobuf.Decode
import Protobuf.Encode


{-| The field numbers for the fields of `Task`. This is mostly useful for internals, like documentation generation.


-}
fieldNumbersTask : { id : Int, info : Int }
fieldNumbersTask =
    Proto.Api.Internals_.fieldNumbersProto__Api__Task


{-| Default for Task. Should only be used for 'required' decoders as an initial value.


-}
defaultTask : Task
defaultTask =
    Proto.Api.Internals_.defaultProto__Api__Task


{-| Declares how to decode a `Task` from Bytes. To actually perform the conversion from Bytes, you need to use Protobuf.Decode.decode from eriktim/elm-protocol-buffers.


-}
decodeTask : Protobuf.Decode.Decoder Task
decodeTask =
    Proto.Api.Internals_.decodeProto__Api__Task


{-| Declares how to encode a `Task` to Bytes. To actually perform the conversion to Bytes, you need to use Protobuf.Encode.encode from eriktim/elm-protocol-buffers.


-}
encodeTask : Task -> Protobuf.Encode.Encoder
encodeTask =
    Proto.Api.Internals_.encodeProto__Api__Task


{-| `Task` message


-}
type alias Task =
    Proto.Api.Internals_.Proto__Api__Task


{-| The field numbers for the fields of `TaskDeleteRequest`. This is mostly useful for internals, like documentation generation.


-}
fieldNumbersTaskDeleteRequest : { id : Int }
fieldNumbersTaskDeleteRequest =
    Proto.Api.Internals_.fieldNumbersProto__Api__TaskDeleteRequest


{-| Default for TaskDeleteRequest. Should only be used for 'required' decoders as an initial value.


-}
defaultTaskDeleteRequest : TaskDeleteRequest
defaultTaskDeleteRequest =
    Proto.Api.Internals_.defaultProto__Api__TaskDeleteRequest


{-| Declares how to decode a `TaskDeleteRequest` from Bytes. To actually perform the conversion from Bytes, you need to use Protobuf.Decode.decode from eriktim/elm-protocol-buffers.


-}
decodeTaskDeleteRequest : Protobuf.Decode.Decoder TaskDeleteRequest
decodeTaskDeleteRequest =
    Proto.Api.Internals_.decodeProto__Api__TaskDeleteRequest


{-| Declares how to encode a `TaskDeleteRequest` to Bytes. To actually perform the conversion to Bytes, you need to use Protobuf.Encode.encode from eriktim/elm-protocol-buffers.


-}
encodeTaskDeleteRequest : TaskDeleteRequest -> Protobuf.Encode.Encoder
encodeTaskDeleteRequest =
    Proto.Api.Internals_.encodeProto__Api__TaskDeleteRequest


{-| `TaskDeleteRequest` message


-}
type alias TaskDeleteRequest =
    Proto.Api.Internals_.Proto__Api__TaskDeleteRequest


{-| The field numbers for the fields of `TaskCreateRequest`. This is mostly useful for internals, like documentation generation.


-}
fieldNumbersTaskCreateRequest : { info : Int }
fieldNumbersTaskCreateRequest =
    Proto.Api.Internals_.fieldNumbersProto__Api__TaskCreateRequest


{-| Default for TaskCreateRequest. Should only be used for 'required' decoders as an initial value.


-}
defaultTaskCreateRequest : TaskCreateRequest
defaultTaskCreateRequest =
    Proto.Api.Internals_.defaultProto__Api__TaskCreateRequest


{-| Declares how to decode a `TaskCreateRequest` from Bytes. To actually perform the conversion from Bytes, you need to use Protobuf.Decode.decode from eriktim/elm-protocol-buffers.


-}
decodeTaskCreateRequest : Protobuf.Decode.Decoder TaskCreateRequest
decodeTaskCreateRequest =
    Proto.Api.Internals_.decodeProto__Api__TaskCreateRequest


{-| Declares how to encode a `TaskCreateRequest` to Bytes. To actually perform the conversion to Bytes, you need to use Protobuf.Encode.encode from eriktim/elm-protocol-buffers.


-}
encodeTaskCreateRequest : TaskCreateRequest -> Protobuf.Encode.Encoder
encodeTaskCreateRequest =
    Proto.Api.Internals_.encodeProto__Api__TaskCreateRequest


{-| `TaskCreateRequest` message


-}
type alias TaskCreateRequest =
    Proto.Api.Internals_.Proto__Api__TaskCreateRequest


{-| The field numbers for the fields of `TaskGetRequest`. This is mostly useful for internals, like documentation generation.


-}
fieldNumbersTaskGetRequest : { id : Int }
fieldNumbersTaskGetRequest =
    Proto.Api.Internals_.fieldNumbersProto__Api__TaskGetRequest


{-| Default for TaskGetRequest. Should only be used for 'required' decoders as an initial value.


-}
defaultTaskGetRequest : TaskGetRequest
defaultTaskGetRequest =
    Proto.Api.Internals_.defaultProto__Api__TaskGetRequest


{-| Declares how to decode a `TaskGetRequest` from Bytes. To actually perform the conversion from Bytes, you need to use Protobuf.Decode.decode from eriktim/elm-protocol-buffers.


-}
decodeTaskGetRequest : Protobuf.Decode.Decoder TaskGetRequest
decodeTaskGetRequest =
    Proto.Api.Internals_.decodeProto__Api__TaskGetRequest


{-| Declares how to encode a `TaskGetRequest` to Bytes. To actually perform the conversion to Bytes, you need to use Protobuf.Encode.encode from eriktim/elm-protocol-buffers.


-}
encodeTaskGetRequest : TaskGetRequest -> Protobuf.Encode.Encoder
encodeTaskGetRequest =
    Proto.Api.Internals_.encodeProto__Api__TaskGetRequest


{-| `TaskGetRequest` message


-}
type alias TaskGetRequest =
    Proto.Api.Internals_.Proto__Api__TaskGetRequest


{-| The field numbers for the fields of `TaskListResponse`. This is mostly useful for internals, like documentation generation.


-}
fieldNumbersTaskListResponse : { tasks : Int }
fieldNumbersTaskListResponse =
    Proto.Api.Internals_.fieldNumbersProto__Api__TaskListResponse


{-| Default for TaskListResponse. Should only be used for 'required' decoders as an initial value.


-}
defaultTaskListResponse : TaskListResponse
defaultTaskListResponse =
    Proto.Api.Internals_.defaultProto__Api__TaskListResponse


{-| Declares how to decode a `TaskListResponse` from Bytes. To actually perform the conversion from Bytes, you need to use Protobuf.Decode.decode from eriktim/elm-protocol-buffers.


-}
decodeTaskListResponse : Protobuf.Decode.Decoder TaskListResponse
decodeTaskListResponse =
    Proto.Api.Internals_.decodeProto__Api__TaskListResponse


{-| Declares how to encode a `TaskListResponse` to Bytes. To actually perform the conversion to Bytes, you need to use Protobuf.Encode.encode from eriktim/elm-protocol-buffers.


-}
encodeTaskListResponse : TaskListResponse -> Protobuf.Encode.Encoder
encodeTaskListResponse =
    Proto.Api.Internals_.encodeProto__Api__TaskListResponse


{-| `TaskListResponse` message


-}
type alias TaskListResponse =
    Proto.Api.Internals_.Proto__Api__TaskListResponse


{-| The field numbers for the fields of `TaskListRequest`. This is mostly useful for internals, like documentation generation.


-}
fieldNumbersTaskListRequest : {}
fieldNumbersTaskListRequest =
    Proto.Api.Internals_.fieldNumbersProto__Api__TaskListRequest


{-| Default for TaskListRequest. Should only be used for 'required' decoders as an initial value.


-}
defaultTaskListRequest : TaskListRequest
defaultTaskListRequest =
    Proto.Api.Internals_.defaultProto__Api__TaskListRequest


{-| Declares how to decode a `TaskListRequest` from Bytes. To actually perform the conversion from Bytes, you need to use Protobuf.Decode.decode from eriktim/elm-protocol-buffers.


-}
decodeTaskListRequest : Protobuf.Decode.Decoder TaskListRequest
decodeTaskListRequest =
    Proto.Api.Internals_.decodeProto__Api__TaskListRequest


{-| Declares how to encode a `TaskListRequest` to Bytes. To actually perform the conversion to Bytes, you need to use Protobuf.Encode.encode from eriktim/elm-protocol-buffers.


-}
encodeTaskListRequest : TaskListRequest -> Protobuf.Encode.Encoder
encodeTaskListRequest =
    Proto.Api.Internals_.encodeProto__Api__TaskListRequest


{-| `TaskListRequest` message


-}
type alias TaskListRequest =
    Proto.Api.Internals_.Proto__Api__TaskListRequest
