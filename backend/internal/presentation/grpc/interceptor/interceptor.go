package interceptor

import (
	"github.com/AoiSatonaka/todo-app/backend/internal/application"
	"google.golang.org/grpc"
)

func GetInterceptors(aua application.UserService) []grpc.UnaryServerInterceptor {
	return []grpc.UnaryServerInterceptor{AuthInterceptor(aua)}
}
