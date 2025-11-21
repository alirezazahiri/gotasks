package taskhandler

import (
	"net/http"

	"github.com/alirezazahiri/gotasks/internal/services/taskservice"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *taskservice.TaskService
}

func New(service *taskservice.TaskService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("", h.CreateTask)
	router.GET("/:id", h.GetTask)
	router.PUT("/:id", h.UpdateTask)
	router.DELETE("/:id", h.DeleteTask)
}

func (h *Handler) CreateTask(c *gin.Context) {
	var req CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := toEntity(&req)
	if err := h.service.CreateTask(task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, toResponse(task))
}

func (h *Handler) GetTask(c *gin.Context) {
	id := c.Param("id")

	task, err := h.service.GetTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	c.JSON(http.StatusOK, toResponse(task))
}

func (h *Handler) UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var req UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.ID = id
	task := toEntityFromUpdate(&req)
	if err := h.service.UpdateTask(task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, toResponse(task))
}

func (h *Handler) DeleteTask(c *gin.Context) {
	id := c.Param("id")

	err := h.service.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
