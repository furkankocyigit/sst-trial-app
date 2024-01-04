package service

import (
	model "backend/internal/apps/todo/model"
	"backend/internal/apps/todo/repository"
)

type TodoServiceInterface interface {
	GetAllTodos() ([]model.Todo, error)
}

type TodoService struct {
	repository repository.TodoRepositoryInterface
}

func NewTodoService(repository repository.TodoRepositoryInterface) *TodoService {
	return &TodoService{
		repository: repository,
	}
}

func (l *TodoService) GetAllTodos() ([]model.Todo, error) {
	
	todos, err := l.repository.GetAllTodos()

	if err != nil {
		return nil, err
	}

	return todos, nil
}