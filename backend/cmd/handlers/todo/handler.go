package handler

import (
	"backend/internal/apps/todo/service"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

type TodoHandler struct {
	service service.TodoServiceInterface
}

func NewTodoHandler(service service.TodoServiceInterface) *TodoHandler {
	return &TodoHandler{ service }
}

func (t *TodoHandler) GetAllTodos() (events.APIGatewayProxyResponse,error) {
	todos, err := t.service.GetAllTodos()
	
	if err != nil {
		 return events.APIGatewayProxyResponse{
			 Body: err.Error(),
			 StatusCode: 500,
		 },nil
	 }

	 body, err := json.Marshal(todos)
	 if err != nil {
		 return events.APIGatewayProxyResponse{
			 Body: err.Error(),
			 StatusCode: 500,
		 },nil
	 }

	 return events.APIGatewayProxyResponse{
		 Body: string(body),
		 StatusCode: 200,
	 },nil
}