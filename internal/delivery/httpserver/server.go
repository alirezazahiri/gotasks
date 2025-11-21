package httpserver

import (
	"github.com/alirezazahiri/gotasks/internal/delivery/httpserver/taskhandler"
	"github.com/alirezazahiri/gotasks/internal/services/taskservice"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func New(taskService *taskservice.TaskService) *Server {
	router := gin.Default()

	taskHandler := taskhandler.New(taskService)
	taskGroup := router.Group("/task")
	taskHandler.RegisterRoutes(taskGroup)

	return &Server{router: router}
}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}

func (s *Server) GetRouter() *gin.Engine {
	return s.router
}
