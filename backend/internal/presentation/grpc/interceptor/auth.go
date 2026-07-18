package interceptor

import (
	"context"
	"strings"

	"github.com/AoiSatonaka/todo-app/backend/internal/application"
	"github.com/AoiSatonaka/todo-app/backend/internal/constant"
	"github.com/AoiSatonaka/todo-app/backend/internal/presentation/grpc/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc/metadata"
)

func AuthInterceptor(ua application.UserService) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// get idtoken
		idt, err := getAuthHeader(ctx)
		if err != nil {
			return nil, err
		}

		// idtoken authorization
		user, err := ua.Authorize(idt)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}

		// setting user to context
		newCtx := context.WithValue(ctx, util.UserKey{}, user)

		// service exec
		res, err := handler(newCtx, req)
		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func getAuthHeader(ctx context.Context) (string, error) {
	// format metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Error(codes.DataLoss, constant.ERR_METADATA_NOT_SET)
	}

	header := md.Get("authorization")[0]
	if len(header) < 1 {
		return "", status.Error(codes.DataLoss, constant.ERR_AUTH_HEADER_NOT_SET)
	}

	splits := strings.Split(header, " ")
	// length and format check
	if len(splits) != 2 || splits[0] != "bearer" {
		return "", status.Error(codes.DataLoss, constant.ERR_AUTH_HEADER_NOT_EXPECTED_FORMAT)
	}

	return splits[1], nil
}
