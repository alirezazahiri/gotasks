package taskhandler

import (
	"time"

	"github.com/alirezazahiri/gotasks/internal/entity"
	"github.com/google/uuid"
)

func toEntity(req *CreateTaskRequest) *entity.Task {
	now := time.Now().Unix()
	
	var dueDate *int64
	if req.DueDateUnix != nil {
		dueDate = req.DueDateUnix
	}

	return &entity.Task{
		ID:          uuid.New().String(),
		Title:       req.Title,
		Description: req.Description,
		Status:      "todo",
		Priority:    "medium",
		DueDate:     dueDate,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func toEntityFromUpdate(req *UpdateTaskRequest) *entity.Task {
	task := &entity.Task{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		Priority:    req.Priority,
		UpdatedAt:   time.Now().Unix(),
	}

	if req.DueDateUnix != nil {
		task.DueDate = req.DueDateUnix
	}

	return task
}

func toResponse(task *entity.Task) *TaskResponse {
	return &TaskResponse{
		ID:              task.ID,
		Title:           task.Title,
		Description:     task.Description,
		Status:          task.Status,
		Priority:        task.Priority,
		DueDateUnix:     *task.DueDate,
		CompletedAtUnix: task.CompletedAt,
		CreatedAtUnix:   task.CreatedAt,
		UpdatedAtUnix:   task.UpdatedAt,
	}
}
