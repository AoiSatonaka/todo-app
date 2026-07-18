package grpc

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/AoiSatonaka/todo-app/backend/internal/application"
	"github.com/AoiSatonaka/todo-app/backend/internal/presentation"
	"github.com/AoiSatonaka/todo-app/backend/internal/presentation/grpc/handler"
	"github.com/AoiSatonaka/todo-app/backend/internal/presentation/grpc/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GrpcServer struct {
	// todopb
	s grpc.Server
}

var _ presentation.IPresenter = &GrpcServer{}

func New(ata *application.ToDoService) (*GrpcServer, error) {
	gp := &GrpcServer{s: *grpc.NewServer()}

	// register reflections
	reflection.Register(&gp.s)

	// register services
	protobuf.RegisterToDoServiceServer(&gp.s, handler.NewToDoServer(ata))

	return gp, nil
}

func (g *GrpcServer) Serve() error {
	port := fmt.Sprintf(":%s", os.Getenv("GRPC_PORT"))
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	log.Printf("[presentation] start GRPC server (port:%s)", port)

	g.s.Serve(listener)
	return nil
}
