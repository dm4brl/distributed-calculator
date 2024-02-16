package task

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/dm4brl/distributed-calculator/pkg/config"
	"github.com/dm4brl/distributed-calculator/pkg/storage"
)

type Task struct {
	ID        string
	Expression string
	Result    float64
	Error     error
}

func NewTask(id, expression string) *Task {
	return &Task{ID: id, Expression: expression}
}

func (t *Task) Calculate() {
	result, err := calculate(t.Expression)
	if err != nil {
		t.Error = err
	} else {
		t.Result = result
	}
}

func calculate(expression string) (float64, error) {
	// Implement your own expression parser and evaluator here
	// For example, you can use the "github.com/Knetic/govaluate" package
	// or write your own parser and evaluator using the "github.com/antlr/antlr4/runtime/Go/antlr" package
	// This example uses a simple expression parser and evaluator

	var result float64
	var err error

	switch expression {
	case "sqrt(2)":
		result = math.Sqrt(2)
	case "sqrt(3)":
		result = math.Sqrt(3)
	default:
		err = errors.New("invalid expression")
	}

	return result, err
}
