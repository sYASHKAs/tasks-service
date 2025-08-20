package grpc

import (
	"fmt"

	userpb "github.com/sYASHKAs/project-protos/proto/user"
	"google.golang.org/grpc"
)

func NewUserClient(addr string) (userpb.UserServiceClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to dial users-service: %w", err)
	}

	client := userpb.NewUserServiceClient(conn)
	return client, conn, err
}
