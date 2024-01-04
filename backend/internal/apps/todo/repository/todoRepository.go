package repository

import model "backend/internal/apps/todo/model"

type TodoRepositoryInterface interface {
	GetAllTodos() ([]model.Todo, error)
}


type TodoRepository struct {
}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{}
}

func (l *TodoRepository) GetAllTodos() ([]model.Todo, error) {
	todos := []model.Todo{
		{
			ID:          "1",
			Title:       "My first todo",
			Description: "My first todo description",
		},
		{
			ID:          "2",
			Title:       "My second todo",
			Description: "My second todo description",
		},
	}

	return todos, nil
}