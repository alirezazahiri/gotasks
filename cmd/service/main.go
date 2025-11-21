package main

import (
	"log"

	"github.com/alirezazahiri/gotasks/internal/config"
	"github.com/alirezazahiri/gotasks/internal/repository/postgresql"
)

func main() {
	cfg := config.Load("config.yml")
	
	repo, err := postgresql.New(cfg)
	if err != nil {
		log.Fatalf("failed to create repository: %v", err)
	}
	defer repo.Close()
}