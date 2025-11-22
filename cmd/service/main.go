package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/alirezazahiri/gotasks/internal/config"
	"github.com/alirezazahiri/gotasks/internal/delivery/grpcserver/taskserver"
	"github.com/alirezazahiri/gotasks/internal/delivery/httpserver"
	pb "github.com/alirezazahiri/gotasks/internal/protobuf/go"
	"github.com/alirezazahiri/gotasks/internal/repository/postgresql"
	"github.com/alirezazahiri/gotasks/internal/repository/taskrepo"
	"github.com/alirezazahiri/gotasks/internal/services/taskservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("application failed: %v", err)
	}
}

func run() error {
	cfg := config.Load("config.yml")

	log.Printf("starting application in %s mode", cfg.Env)

	db, err := postgresql.New(cfg)
	if err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return fmt.Errorf("database connection failed: %w", err)
	}

	taskRepo := taskrepo.New(db)
	taskSvc := taskservice.New(taskRepo)

	errChan := make(chan error, 2)
	done := make(chan struct{})

	go startGRPCServer(cfg, taskSvc, errChan)

	grpcClient, err := createGRPCClient(cfg)
	if err != nil {
		return fmt.Errorf("failed to create grpc client: %w", err)
	}

	go startHTTPServer(cfg, grpcClient, errChan)

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint
		log.Println("shutting down gracefully...")
		close(done)
	}()

	select {
	case err := <-errChan:
		return fmt.Errorf("server error: %w", err)
	case <-done:
		return nil
	}
}

func createGRPCClient(cfg config.Config) (pb.TaskServiceClient, error) {
	addr := fmt.Sprintf("localhost:%s", cfg.GRPCServer.Port)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to grpc server: %w", err)
	}

	return pb.NewTaskServiceClient(conn), nil
}

func startHTTPServer(cfg config.Config, grpcClient pb.TaskServiceClient, errChan chan<- error) {
	server := httpserver.New(grpcClient)
	addr := fmt.Sprintf(":%s", cfg.HTTPServer.Port)

	log.Printf("http server listening on %s (proxying to grpc)", addr)

	if err := server.Run(addr); err != nil {
		errChan <- fmt.Errorf("http server failed: %w", err)
	}
}

func startGRPCServer(cfg config.Config, taskSvc *taskservice.TaskService, errChan chan<- error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GRPCServer.Port))
	if err != nil {
		errChan <- fmt.Errorf("failed to create grpc listener: %w", err)
		return
	}

	grpcServer := grpc.NewServer()
	taskServer := taskserver.New(taskSvc)
	pb.RegisterTaskServiceServer(grpcServer, taskServer)

	log.Printf("grpc server listening on :%s (business logic layer)", cfg.GRPCServer.Port)

	if err := grpcServer.Serve(listener); err != nil {
		errChan <- fmt.Errorf("grpc server failed: %w", err)
	}
}
