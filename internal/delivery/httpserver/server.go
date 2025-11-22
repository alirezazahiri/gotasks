package httpserver

import (
	"github.com/alirezazahiri/gotasks/internal/delivery/httpserver/taskhandler"
	pb "github.com/alirezazahiri/gotasks/internal/protobuf/go"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func New(grpcClient pb.TaskServiceClient) *Server {
	router := gin.Default()

	taskHandler := taskhandler.New(grpcClient)
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
