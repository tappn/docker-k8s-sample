package service

import (
	"github.com/tappn/docker-k8s-sample/internal/app/domain/model"
	"github.com/tappn/docker-k8s-sample/internal/app/infrastructure/repository"
)

type Service struct {
	repositor *repository.Repository
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		repositor: repository,
	}
}

func (s *Service) GetAll() ([]*model.Todo, error) {

	todos, err := s.repositor.FindAll()
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (s *Service) GetByID(id int64) (*model.Todo, error) {
	todo, err := s.repositor.FindByID(id)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *Service) Create(todo *model.Todo) (*model.Todo, error) {
	if err := s.repositor.Store(todo); err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *Service) Edit(todo *model.Todo) (*model.Todo, error) {
	if err := s.repositor.Update(todo); err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *Service) Delete(id int64) error {
	if err := s.repositor.Delete(id); err != nil {
		return err
	}

	return nil
}
