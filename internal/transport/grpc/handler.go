package grpc

import (
	"context"
	"fmt"

	taskpb "github.com/sYASHKAs/project-protos/proto/task"
	userpb "github.com/sYASHKAs/project-protos/proto/user"
	"github.com/sYASHKAs/tasks-service/internal/task"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	svc        task.TaskService
	userClient userpb.UserServiceClient
	taskpb.UnimplementedTaskServiceServer
}

func NewHandler(svc task.TaskService, uc userpb.UserServiceClient) *Handler {
	return &Handler{svc: svc, userClient: uc}
}

func (h *Handler) CreateTask(ctx context.Context, req *taskpb.CreateTaskRequest) (*taskpb.CreateTaskResponse, error) {
	if _, err := h.userClient.GetUser(ctx, &userpb.GetUserRequest{Id: req.UserId}); err != nil {
		return nil, fmt.Errorf("user %d not found: %w", req.UserId, err)
	}

	t, err := h.svc.CreateTask(task.TaskRequest{
		Task:   req.Title,
		IsDone: req.IsDone,
		UserID: req.UserId,
	})
	if err != nil {
		return nil, err
	}

	return &taskpb.CreateTaskResponse{
		Task: &taskpb.Task{
			Id:     t.ID,
			Title:  t.Task,
			IsDone: t.IsDone,
			UserId: t.UserID,
		},
	}, nil
}

func (h *Handler) GetTask(ctx context.Context, req *taskpb.GetTaskRequest) (*taskpb.GetTaskResponse, error) {
	t, err := h.svc.GetTaskByID(req.GetId())
	if err != nil {
		return nil, err
	}

	return &taskpb.GetTaskResponse{
		Task: &taskpb.Task{
			Id:     t.ID,
			Title:  t.Task,
			IsDone: t.IsDone,
			UserId: t.UserID,
		},
	}, nil
}

func (h *Handler) ListTasks(ctx context.Context, _ *taskpb.ListTasksRequest) (*taskpb.ListTasksResponse, error) {
	tasks, err := h.svc.GetAllTasks()
	if err != nil {
		return nil, err
	}

	var pbTasks []*taskpb.Task
	for _, t := range tasks {
		pbTasks = append(pbTasks, &taskpb.Task{
			Id:     t.ID,
			Title:  t.Task,
			IsDone: t.IsDone,
			UserId: t.UserID,
		})
	}

	return &taskpb.ListTasksResponse{Tasks: pbTasks}, nil
}

func (h *Handler) ListTasksByUser(ctx context.Context, req *taskpb.ListTasksByUserRequest) (*taskpb.ListTasksByUserResponse, error) {
	tasks, err := h.svc.GetTasksForUser(req.UserId)
	if err != nil {
		return nil, err
	}

	var pbTasks []*taskpb.Task
	for _, t := range tasks {
		pbTasks = append(pbTasks, &taskpb.Task{
			Id:     t.ID,
			Title:  t.Task,
			IsDone: t.IsDone,
			UserId: t.UserID,
		})
	}

	return &taskpb.ListTasksByUserResponse{Tasks: pbTasks}, nil
}

func (h *Handler) UpdateTask(ctx context.Context, req *taskpb.UpdateTaskRequest) (*taskpb.UpdateTaskResponse, error) {
	if _, err := h.userClient.GetUser(ctx, &userpb.GetUserRequest{Id: req.UserId}); err != nil {
		return nil, fmt.Errorf("user %d not found: %w", req.UserId, err)
	}

	t, err := h.svc.UpdateTask(req.GetId(), task.TaskRequest{
		Task:   req.Title,
		IsDone: req.IsDone,
		UserID: req.UserId,
	})
	if err != nil {
		return nil, err
	}

	return &taskpb.UpdateTaskResponse{
		Task: &taskpb.Task{
			Id:     t.ID,
			Title:  t.Task,
			IsDone: t.IsDone,
			UserId: t.UserID,
		},
	}, nil
}

func (h *Handler) DeleteTask(ctx context.Context, req *taskpb.DeleteTaskRequest) (*emptypb.Empty, error) {
	if _, err := h.userClient.GetUser(ctx, &userpb.GetUserRequest{Id: req.UserId}); err != nil {
		return nil, fmt.Errorf("user %d not found: %w", req.UserId, err)
	}

	err := h.svc.DeleteTask(req.GetId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
