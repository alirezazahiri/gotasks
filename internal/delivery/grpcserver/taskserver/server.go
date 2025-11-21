package taskserver

import (
	"context"
	"time"

	"github.com/alirezazahiri/gotasks/internal/entity"
	pb "github.com/alirezazahiri/gotasks/internal/protobuf/go"
	"github.com/alirezazahiri/gotasks/internal/services/taskservice"
	"github.com/google/uuid"
)

type Server struct {
	pb.UnimplementedTaskServiceServer
	service *taskservice.TaskService
}

func New(service *taskservice.TaskService) *Server {
	return &Server{service: service}
}

func (s *Server) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	now := time.Now().Unix()
	var dueDate *int64
	if req.DueDateUnix > 0 {
		dueDate = &req.DueDateUnix
	}
	task := &entity.Task{
		ID:          uuid.New().String(),
		Title:       req.Title,
		Description: req.Description,
		Status:      "todo",
		Priority:    "medium",
		DueDate:     dueDate,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	err := s.service.CreateTask(task)
	if err != nil {
		return nil, err
	}

	return &pb.CreateTaskResponse{
		Task: toProto(task),
	}, nil
}

func (s *Server) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.GetTaskResponse, error) {
	task, err := s.service.GetTask(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetTaskResponse{
		Task: toProto(task),
	}, nil
}

func (s *Server) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.UpdateTaskResponse, error) {
	task := &entity.Task{
		ID:          req.Id,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		Priority:    req.Priority,
		UpdatedAt:   time.Now().Unix(),
	}

	err := s.service.UpdateTask(task)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateTaskResponse{
		Task: toProto(task),
	}, nil
}

func (s *Server) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	err := s.service.DeleteTask(req.Id)
	if err != nil {
		return &pb.DeleteTaskResponse{Success: false}, err
	}

	return &pb.DeleteTaskResponse{Success: true}, nil
}
