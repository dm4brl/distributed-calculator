package task

import (
	"fmt"
	"sync"

	"github.com/dm4brl/distributed-calculator/pkg/config"
	"github.com/dm4brl/distributed-calculator/pkg/storage"
)

type TaskManager struct {
	cfg *config.Config
	db  storage.Database

	tasks     []*Task
	taskQueue chan *Task
	wg        sync.WaitGroup
}

func NewTaskManager(cfg *config.Config) *TaskManager {
	var db storage.Database
	var err error

	switch cfg.Database.Type {
	case "postgres":
		db, err = storage.NewPostgres(cfg)
	case "redis":
		db, err = storage.NewRedis(cfg)
	default:
		panic("invalid database type")
	}

	if err != nil {
		panic(err)
	}

	return &TaskManager{
		cfg:       cfg,
		db:        db,
		taskQueue: make(chan *Task, cfg.TaskManager.Concurrency),
	}
}

func (tm *TaskManager) RunAgent() {
	for {
		task := <-tm.taskQueue
		tm.wg.Add(1)
