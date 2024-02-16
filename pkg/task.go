package task

type Task struct {
	ID          string `json:"id"`
	Expression  string `json:"expression"`
	Result      float64 `json:"result"`
	Status      string `json:"status"`
}

func NewTask(id, expression string) *Task {
	return &Task{
		ID:         id,
		Expression: expression,
		Status:     "pending",
	}
}
