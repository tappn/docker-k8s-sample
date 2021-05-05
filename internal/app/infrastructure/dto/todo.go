package dto

import (
	"database/sql"
	"time"

	"github.com/tappn/docker-k8s-sample/internal/app/domain/model"
)

type Todo struct {
	ID          int64        `db:"id"`
	Title       string       `db:"title"`
	Description string       `db:"description"`
	Deadline    time.Time    `db:"deadline"`
	CreatedAt   time.Time    `db:"created_at"`
	UpdatedAt   time.Time    `db:"updated_at"`
	DeletedAt   sql.NullTime `db:"deleted_at"`
}

func (t *Todo) ToModel() *model.Todo {
	return &model.Todo{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Deadline:    t.Deadline,
	}
}
