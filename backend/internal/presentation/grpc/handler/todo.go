package handler

import (
	"context"

	"github.com/AoiSatonaka/todo-app/backend/internal/application"
	"github.com/AoiSatonaka/todo-app/backend/internal/constant"
	"github.com/AoiSatonaka/todo-app/backend/internal/presentation/grpc/protobuf"
	"github.com/AoiSatonaka/todo-app/backend/internal/presentation/grpc/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type ToDoServer struct {
	protobuf.UnimplementedToDoServiceServer
	ta *application.ToDoService
}

var _ protobuf.ToDoServiceServer = &ToDoServer{}

func NewToDoServer(ata *application.ToDoService) *ToDoServer {
	return &ToDoServer{ta: ata}
}

func (ts *ToDoServer) List(ctx context.Context, in *protobuf.TaskListRequest) (*protobuf.TaskListResponse, error) {
	// get tasks
	tasks, err := ts.ta.List(util.GetUid(ctx))
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}

	// create response
	var ptsl []*protobuf.Task
	for _, t := range *tasks {
		pt, err := util.ConvertToProtoTask(&t)
		if err != nil {
			return nil, status.Error(codes.Aborted, err.Error())
		}
		ptsl = append(ptsl, pt)
	}

	return &protobuf.TaskListResponse{Tasks: ptsl}, nil
}

func (ts *ToDoServer) Get(ctx context.Context, in *protobuf.TaskGetRequest) (*protobuf.TaskGetResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.DataLoss, constant.ERR_METADATA_NOT_SET)
	}
	md.Get("Id")
	task, err := ts.ta.Get(in.Id, util.GetUid(ctx))
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}
	// create response
	t, err := util.ConvertToProtoTask(task)
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}
	return &protobuf.TaskGetResponse{Task: t}, nil
}

func (ts *ToDoServer) Create(ctx context.Context, in *protobuf.TaskCreateRequest) (*protobuf.TaskCreateResponse, error) {
	in.GetInfo()
	in.Info.ProtoMessage()
	return &protobuf.TaskCreateResponse{}, nil
}

func (ts *ToDoServer) Update(ctx context.Context, in *protobuf.TaskUpdateRequest) (*protobuf.TaskUpdateResponse, error) {
	return &protobuf.TaskUpdateResponse{}, nil
}

func (ts *ToDoServer) Delete(ctx context.Context, in *protobuf.TaskDeleteRequest) (*protobuf.TaskDeleteResponse, error) {
	return &protobuf.TaskDeleteResponse{}, nil
}
