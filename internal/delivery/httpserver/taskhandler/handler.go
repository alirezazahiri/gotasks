package taskhandler

import (
	"context"
	"net/http"
	"strconv"

	pb "github.com/alirezazahiri/gotasks/internal/protobuf/go"
	"github.com/alirezazahiri/gotasks/pkg/envelope"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	grpcClient pb.TaskServiceClient
}

func New(grpcClient pb.TaskServiceClient) *Handler {
	return &Handler{grpcClient: grpcClient}
}

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("", h.CreateTask)
	router.GET("/:id", h.GetTask)
	router.GET("/list", h.ListTasks)
	router.PUT("/:id", h.UpdateTask)
	router.DELETE("/:id", h.DeleteTask)
}

func (h *Handler) CreateTask(c *gin.Context) {
	var req CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		envelope.ValidationError(c, "Invalid request data", err.Error())
		return
	}

	grpcReq := &pb.CreateTaskRequest{
		Title:       req.Title,
		Description: req.Description,
	}
	if req.DueDateUnix != nil {
		grpcReq.DueDateUnix = *req.DueDateUnix
	}

	resp, err := h.grpcClient.CreateTask(context.Background(), grpcReq)
	if err != nil {
		envelope.InternalServerError(c, err.Error())
		return
	}

	envelope.Created(c, toResponseFromProto(resp.Task))
}

func (h *Handler) GetTask(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.grpcClient.GetTask(context.Background(), &pb.GetTaskRequest{Id: id})
	if err != nil {
		envelope.NotFound(c, "Task not found")
		return
	}

	envelope.OK(c, toResponseFromProto(resp.Task))
}

func (h *Handler) ListTasks(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("page_size", "10")

	pageInt, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		envelope.BadRequest(c, "Invalid page parameter")
		return
	}
	pageSizeInt, err := strconv.ParseInt(pageSize, 10, 64)
	if err != nil {
		envelope.BadRequest(c, "Invalid page_size parameter")
		return
	}

	grpcReq := &pb.ListTasksRequest{
		Page:     pageInt,
		PageSize: pageSizeInt,
	}

	resp, err := h.grpcClient.ListTasks(context.Background(), grpcReq)
	if err != nil {
		envelope.InternalServerError(c, err.Error())
		return
	}

	pagination := &envelope.Pagination{
		Page:       resp.Page,
		PageSize:   resp.PageSize,
		Total:      resp.Total,
		TotalPages: resp.TotalPages,
	}

	envelope.SuccessWithPagination(c, http.StatusOK, toResponseFromProtoList(resp.Tasks), pagination)
}

func (h *Handler) UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var req UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		envelope.ValidationError(c, "Invalid request data", err.Error())
		return
	}

	grpcReq := &pb.UpdateTaskRequest{
		Id:          id,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		Priority:    req.Priority,
	}
	if req.DueDateUnix != nil {
		grpcReq.DueDateUnix = *req.DueDateUnix
	}

	resp, err := h.grpcClient.UpdateTask(context.Background(), grpcReq)
	if err != nil {
		envelope.InternalServerError(c, err.Error())
		return
	}

	envelope.OK(c, toResponseFromProto(resp.Task))
}

func (h *Handler) DeleteTask(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.grpcClient.DeleteTask(context.Background(), &pb.DeleteTaskRequest{Id: id})
	if err != nil {
		envelope.InternalServerError(c, err.Error())
		return
	}

	envelope.OK(c, gin.H{"deleted": resp.Success})
}
