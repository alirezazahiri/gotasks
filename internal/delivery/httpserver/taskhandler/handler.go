package taskhandler

import (
	"context"
	"net/http"
	"strconv"

	pb "github.com/alirezazahiri/gotasks/internal/protobuf/go"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, toResponseFromProto(resp.Task))
}

func (h *Handler) GetTask(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.grpcClient.GetTask(context.Background(), &pb.GetTaskRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	c.JSON(http.StatusOK, toResponseFromProto(resp.Task))
}

func (h *Handler) ListTasks(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("page_size", "10")

	pageInt, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page"})
		return
	}
	pageSizeInt, err := strconv.ParseInt(pageSize, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page size"})
		return
	}

	grpcReq := &pb.ListTasksRequest{
		Page:     pageInt,
		PageSize: pageSizeInt,
	}

	resp, err := h.grpcClient.ListTasks(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tasks":       toResponseFromProtoList(resp.Tasks),
		"total":       resp.Total,
		"page":        resp.Page,
		"page_size":   resp.PageSize,
		"total_pages": resp.TotalPages,
	})
}

func (h *Handler) UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var req UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, toResponseFromProto(resp.Task))
}

func (h *Handler) DeleteTask(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.grpcClient.DeleteTask(context.Background(), &pb.DeleteTaskRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": resp.Success})
}
