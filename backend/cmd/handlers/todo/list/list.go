package main

import (
	handler "backend/cmd/handlers/todo"
	"backend/internal/apps/todo/repository"
	"backend/internal/apps/todo/service"

	"github.com/aws/aws-lambda-go/lambda"
)

func main(){
	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(todoRepository)
	todoHandler := handler.NewTodoHandler(todoService)

	lambda.Start(todoHandler.GetAllTodos)
}