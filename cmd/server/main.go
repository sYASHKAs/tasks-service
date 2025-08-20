package main

import (
	"log"

	"github.com/sYASHKAs/tasks-service/internal/database"
	"github.com/sYASHKAs/tasks-service/internal/task"
	transportgrpc "github.com/sYASHKAs/tasks-service/internal/transport/grpc"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	repo := task.NewTaskRepository(db)
	svc := task.NewTaskService(repo)

	userClient, conn, err := transportgrpc.NewUserClient("localhost:50051")
	if err != nil {
		log.Fatalf("failed to connect to users: %v", err)
	}
	defer conn.Close()

	if err := transportgrpc.RunGRPC(svc, userClient); err != nil {
		log.Fatalf("Tasks gRPC server error: %v", err)
	}
}
