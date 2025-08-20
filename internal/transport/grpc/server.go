package grpc

import (
	"fmt"
	"net"

	taskpb "github.com/sYASHKAs/project-protos/proto/task"
	userpb "github.com/sYASHKAs/project-protos/proto/user"
	"github.com/sYASHKAs/tasks-service/internal/task"
	"google.golang.org/grpc"
)

func RunGRPC(svc task.TaskService, uc userpb.UserServiceClient) error {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	grpcServer := grpc.NewServer()
	taskpb.RegisterTaskServiceServer(grpcServer, NewHandler(svc, uc))

	fmt.Println("gRPC server is running on port 50052")
	return grpcServer.Serve(lis)
}
