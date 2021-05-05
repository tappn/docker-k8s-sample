package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/tappn/docker-k8s-sample/internal/app/domain/model"
	"github.com/tappn/docker-k8s-sample/internal/app/infrastructure/dto"
)

type Repository struct {
	db *sqlx.DB
}

func NewRespository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) FindAll() ([]*model.Todo, error) {
	dto := []*dto.Todo{}
	err := r.db.Select(&dto, "SELECT * FROM todos WHERE deleted_at is null")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	todos := []*model.Todo{}
	for _, v := range dto {
		todos = append(todos, v.ToModel())
	}

	return todos, nil
}

func (r *Repository) FindByID(id int64) (*model.Todo, error) {
	dto := new(dto.Todo)
	err := r.db.Get(dto,
		`
    SELECT 
      * 
    FROM 
      todos 
    WHERE 
      id = ?
      AND
      deleted_at is null`,
		id,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	todo := dto.ToModel()

	return todo, nil
}

func (r *Repository) Store(todo *model.Todo) error {
	now := time.Now()

	dto := &dto.Todo{
		Title:       todo.Title,
		Description: todo.Description,
		Deadline:    todo.Deadline,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	tx := r.db.MustBegin()
	result, err := tx.NamedExec(`
  INSERT INTO
    todos(title, description, deadline, created_at, updated_at)
  VALUES
    (:title, :description, :deadline, :created_at, :updated_at)`,
		dto)

	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	todo.ID, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) Update(todo *model.Todo) error {
	now := time.Now()

	dto := &dto.Todo{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Deadline:    todo.Deadline,
		UpdatedAt:   now,
	}

	tx := r.db.MustBegin()
	_, err := tx.NamedExec(`
  UPDATE
    todos
  SET
    title = :title,
    description = :description,
    deadline = :deadline,
    updated_at = :updated_at
  WHERE
    id = :id
    AND
    deleted_at is null`,
		dto)

	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *Repository) Delete(id int64) error {
	now := time.Now()

	dto := &dto.Todo{
		ID:        id,
		UpdatedAt: now,
		DeletedAt: sql.NullTime{Time: now, Valid: true},
	}

	tx := r.db.MustBegin()
	_, err := tx.NamedExec(`
  UPDATE
    todos
  SET
    updated_at = :updated_at,
    deleted_at = :deleted_at
  WHERE
    id = :id
    AND
    deleted_at is null`,
		dto)

	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
