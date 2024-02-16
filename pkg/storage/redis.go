package storage

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dm4brl/distributed-calculator/pkg/config"
	"github.com/go-redis/redis/v8"
)

type Redis struct {
	client *redis.Client
}

func NewRedis(cfg *config.Config) (*Redis, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Database.Host, cfg.Database.Port),
		Password: cfg.Database.Password,
		DB:       cfg.Database.Name,
	})

	_, err := client.Ping(client.Context()).Result()
	if err != nil {
		return nil, err
	}

	return &Redis{client: client}, nil
}

func (r *Redis) SaveTask(task *task.Task) error {
	err := r.client.HSet(context.Background(), "tasks", task.ID, task.Result).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) GetTask(id string) (*task.Task, error) {
	result, err := r.client.HGet(context.Background(), "tasks", id).Result()
	if err != nil {
		return nil, err
	}

	resultFloat, err := strconv.ParseFloat(result, 64)
	if err != nil {
		return nil, err
	}

	return &task.Task{ID: id, Result: resultFloat}, nil
}

func (r *Redis) GetTasks() ([]*task.Task, error) {
	tasks := make(map[string]string)
	err := r.client.HGetAll(context.Background(), "tasks").Scan(&tasks)
	if err != nil {
		return nil, err
	}

	var resultTasks []*task.Task
	for id, result := range tasks {
		resultFloat, err := strconv.ParseFloat(result, 64)
		if err != nil {
			return nil, err
		}
		resultTasks =
