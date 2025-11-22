package taskhandler

import (
	pb "github.com/alirezazahiri/gotasks/internal/protobuf/go"
)

// toResponseFromProto converts gRPC proto Task to HTTP TaskResponse
func toResponseFromProto(protoTask *pb.Task) *TaskResponse {
	resp := &TaskResponse{
		ID:            protoTask.Id,
		Title:         protoTask.Title,
		Description:   protoTask.Description,
		Status:        protoTask.Status,
		Priority:      protoTask.Priority,
		DueDateUnix:   protoTask.DueDateUnix,
		CreatedAtUnix: protoTask.CreatedAtUnix,
		UpdatedAtUnix: protoTask.UpdatedAtUnix,
	}

	if protoTask.CompletedAtUnix > 0 {
		resp.CompletedAtUnix = &protoTask.CompletedAtUnix
	}

	return resp
}


func toResponseFromProtoList(protoTasks []*pb.Task) []*TaskResponse {
	resp := make([]*TaskResponse, len(protoTasks))
	for i, protoTask := range protoTasks {
		resp[i] = toResponseFromProto(protoTask)
	}
	return resp
}