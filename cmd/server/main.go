package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dm4brl/distributed-calculator/api/v1"
	"github.com/dm4brl/distributed-calculator/pkg/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	router := v1.NewRouter(cfg)

	log.Printf("Starting server on %s", cfg.Server.Address)
	log.Fatal(http.ListenAndServe(cfg.Server.Address, router))
}
