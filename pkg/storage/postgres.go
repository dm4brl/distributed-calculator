package storage

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dm4brl/distributed-calculator/pkg/config"
	"github.com/go-pg/pg/v10"
)

type Postgres struct {
	client *pg.DB
}

func NewPostgres(cfg *config.Config) (*Postgres, error) {
	client := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Database.Host, cfg.Database.Port),
		User:     cfg.Database.Username,
		Password: cfg.Database.Password,
		Database: cfg.Database.Name,
	})

	_, err := client.Exec("SELECT 1")
	if err != nil {
		return nil, err
	}

	return &Postgres{client: client}, nil
}

func (p *Postgres) SaveTask(task *task.Task) error {
	_, err := p.client.Model(task).Insert()
	return err
}

func (p *Postgres) GetTask(id string) (*task.Task, error) {
	task := &task.Task{ID: id}
	err := p.client.Model(task).WherePK().Select()
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (p *Postgres) GetTasks() ([]*task.Task, error) {
	var tasks []*task.Task
	err := p.client.Model(&tasks).Select()
	return tasks, err
}
