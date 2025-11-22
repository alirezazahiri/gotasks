package taskserver

import (
	"github.com/alirezazahiri/gotasks/internal/entity"
	pb "github.com/alirezazahiri/gotasks/internal/protobuf/go"
)

func toProto(task *entity.Task) *pb.Task {
	var completedAtUnix int64
	if task.CompletedAt != nil {
		completedAtUnix = *task.CompletedAt
	}

	var dueDateUnix int64
	if task.DueDate != nil {
		dueDateUnix = *task.DueDate
	}

	return &pb.Task{
		Id:              task.ID,
		Title:           task.Title,
		Description:     task.Description,
		Status:          task.Status,
		Priority:        task.Priority,
		DueDateUnix:     dueDateUnix,
		CompletedAtUnix: completedAtUnix,
		CreatedAtUnix:   task.CreatedAt,
		UpdatedAtUnix:   task.UpdatedAt,
	}
}


func toProtoList(tasks []*entity.Task) []*pb.Task {
	protoTasks := make([]*pb.Task, len(tasks))
	for i, task := range tasks {
		protoTasks[i] = toProto(task)
	}
	return protoTasks
}