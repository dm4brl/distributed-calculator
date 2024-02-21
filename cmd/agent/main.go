package main

import (
	"log"
	"net/http"

	"github.com/dm4brl/distributed-calculator/pkg/config"
	"github.com/dm4brl/distributed-calculator/pkg/task"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	taskManager := task.NewTaskManager(cfg)

	http.HandleFunc("/task", taskManager.HandleTask)

	log.Printf("Starting server on %s", cfg.Server.Address)
	log.Fatal(http.ListenAndServe(cfg.Server.Address, nil))
}
