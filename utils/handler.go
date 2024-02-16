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

type Handler struct {
	storage storage.Storage
}

func NewHandler(cfg *config.Config) (*Handler, error) {
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

	return &Handler{storage: storage}, nil
}

func (h *Handler) SaveTask(w http.ResponseWriter, r *http.Request) {
	task := &task.Task{
		ID:       chi.URLParam(r, "id"),
		Result:   chi.URLParam(r, "result"),
		CreatedAt: time.Now(),
	}

	err := h.storage.SaveTask(task)
	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	render.JSON(w, r, task)
}

func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	task, err := h.storage.GetTask(id)
	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	render.JSON(w, r, task)
}

func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.storage.GetTasks()
	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	render.JSON(w, r, tasks)
}
