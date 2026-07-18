package util

import (
	"context"
	"fmt"

	"github.com/AoiSatonaka/todo-app/backend/internal/constant"
	"github.com/AoiSatonaka/todo-app/backend/internal/domain/task"
	"github.com/AoiSatonaka/todo-app/backend/internal/domain/task/priority"
	"github.com/AoiSatonaka/todo-app/backend/internal/domain/task/status"
	"github.com/AoiSatonaka/todo-app/backend/internal/domain/user"
	"github.com/AoiSatonaka/todo-app/backend/internal/presentation/grpc/protobuf"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertToProtoPriority(p priority.Priority) (protobuf.Task_Info_Priority, error) {
	_, ok := protobuf.Task_Info_Priority_name[int32(p)]
	if !ok && p == priority.Priority(protobuf.Task_Info_STATUS_UNSPECIFIED) {
		return protobuf.Task_Info_PRIORITY_UNSPECIFIED,
			fmt.Errorf(constant.ERR_CONVERT_TO_PROTO_PRIORITY)
	}

	return protobuf.Task_Info_Priority(p), nil
}

func convertToProtoStatus(s status.Status) (protobuf.Task_Info_Status, error) {
	_, ok := protobuf.Task_Info_Status_name[int32(s)]
	if !ok && s == status.Status(protobuf.Task_Info_STATUS_UNSPECIFIED) {
		return protobuf.Task_Info_STATUS_UNSPECIFIED,
			fmt.Errorf(constant.ERR_CONVERT_TO_PROTO_STATUS)
	}

	return protobuf.Task_Info_Status(s), nil
}

func ConvertToProtoTask(t *task.Task) (*protobuf.Task, error) {
	pp, err := convertToProtoPriority(t.Priority)
	if err != nil {
		return nil, err
	}

	ps, err := convertToProtoStatus(t.Status)
	if err != nil {
		return nil, err
	}

	pt := timestamppb.New(t.Limit)

	return &protobuf.Task{
		Id: t.Id,
		Info: &protobuf.Task_Info{
			Title:       t.Title,
			Description: t.Description,
			Limit:       pt,
			Priority:    pp,
			Status:      ps,
			Labels:      t.Labels,
		},
	}, nil
}

func ConvertToDomainTask(t *protobuf.Task) (*task.Task, error) {
	in := t.GetInfo()
	ip, is := int32(in.GetPriority()), int32(in.GetStatus())
	return task.New(
		t.GetId(),
		in.GetTitle(),
		in.GetDescription(),
		in.GetLimit().AsTime(),
		ip,
		is,
		in.GetLabels(),
	)
}

func GetUid(ctx context.Context) string {
	return ctx.Value(UserKey{}).(user.User).Id
}
