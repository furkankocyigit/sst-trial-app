package service

import (
	model "backend/internal/apps/todo/model"
	"backend/internal/apps/todo/repository"
)

type TodoServiceInterface interface {
	GetAllTodos() ([]model.Todo, error)
	CreateTodo(model.Todo) (model.Todo, error)
	FindTodoById(string) (model.Todo, error)
	DeleteTodo(string) (bool, error)
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

func (l *TodoService) CreateTodo(todo model.Todo) (model.Todo, error) {
	createdTodo, err := l.repository.CreateTodo(todo)
	if err != nil {
		return model.Todo{}, err
	}

	return createdTodo, nil
}

func (l *TodoService) FindTodoById(id string) (model.Todo, error) {
	todo, err := l.repository.FindTodoById(id)
	if err != nil {
		return model.Todo{}, err
	}

	return todo, nil
}

func (l *TodoService) DeleteTodo (id string) (bool, error) {
	isDeleted, err := l.repository.DeleteTodo(id)
	
	if err != nil {
		return false, err
	}

	return isDeleted, nil
}