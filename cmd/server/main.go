package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dm4brl/distributed-calculator/pkg/config"
	"github.com/dm4brl/distributed-calculator/pkg/storage"
	"github.com/dm4brl/distributed-calculator/pkg/task"
)

func main() {
	cfg := config.LoadConfig()
	storage := storage.NewPostgresStorage(cfg.DatabaseURL)
	taskManager := task.NewTaskManager(storage)

	http.HandleFunc("/calculate", func(w http.ResponseWriter, r *http.Request) {
		// Обработка запроса на вычисление
		taskManager.HandleTask(w, r)
	})

	port := fmt.Sprintf(":%d", cfg.Port)
	log.Printf("Starting server on port %d", cfg.Port)
	log.Fatal(http.ListenAndServe(port, nil))
}
