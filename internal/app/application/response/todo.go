package response

import (
	"time"

	"github.com/tappn/docker-k8s-sample/internal/app/domain/model"
)

type Todo struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
}

func NewTodo(todo *model.Todo) *Todo {
	return &Todo{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Deadline:    todo.Deadline,
	}
}
