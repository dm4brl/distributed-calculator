package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/dm4brl/distributed-calculator/pkg/config"
	"github.com/dm4brl/distributed-calculator/pkg/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func saveTask(cfg *config.Config, id, result string) error {
	var storage storage.Storage
	var err error

	switch cfg.Database.Type {
	case "postgres":
		storage, err = storage.NewPostgres(cfg)
	case "redis":
		storage, err = storage.NewRedis(cfg)
	default:
		return errors.New("unknown database type")
	}

	if err != nil {
		return err
	}

	task := &storage.Task{
		ID:       id,
		Result:   result,
		CreatedAt: time.Now(),
	}

	err = storage.SaveTask(task)
	if err != nil {
		return err
	}

	return nil
}

func getTask(cfg *config.Config, id string) (*storage.Task, error) {
	var storage storage.Storage
	var err error

	switch cfg.Database.Type {
	case "postgres":
		storage, err = storage.NewPostgres(cfg)
	case "redis":
		storage, err = storage.NewRedis(cfg)
	default:
		return nil, errors.New("unknown database type")
	}

	if err != nil {
		return nil, err
	}

	task, err := storage.GetTask(id)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func getTasks(cfg *config.Config) ([]*storage.Task, error) {
	var storage storage.Storage
	var err error

	switch cfg.Database.Type {
	case "postgres":
		storage, err = storage.NewPostgres(cfg)
	case "redis":
		storage, err = storage.NewRedis(cfg)
	default:
		return nil, errors.New("unknown database type")
	}

	if err != nil {
		return nil, err
	}

	tasks, err := storage.GetTasks()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
