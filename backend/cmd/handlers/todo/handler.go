package handler

import (
	model "backend/internal/apps/todo/model"
	"backend/internal/apps/todo/service"
	"encoding/json"
	"fmt"

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
			Headers: map[string]string{"Content-Type": "application/json"},
			 Body: err.Error(),
			 StatusCode: 500,
		 },nil
	 }

	 body, err := json.Marshal(todos)
	 if err != nil {
		 return events.APIGatewayProxyResponse{
			Headers: map[string]string{"Content-Type": "application/json"},
			 Body: err.Error(),
			 StatusCode: 500,
		 },nil
	 }

	 return events.APIGatewayProxyResponse{
		Headers: map[string]string{"Content-Type": "application/json"},
		 Body: string(body),
		 StatusCode: 200,
	 },nil
}

func (t *TodoHandler) CreateTodo(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse,error) {
	fmt.Println("Request body:", request.Body)
	todoRequest := request.Body
	var todopayLoad model.Todo
	err := json.Unmarshal([]byte(todoRequest), &todopayLoad)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Headers: map[string]string{"Content-Type": "application/json"},
			Body: "Invalid request body" + err.Error(),
			StatusCode: 500,
		},nil
	}

	todo, err := t.service.CreateTodo(todopayLoad)

	if err != nil {
		return events.APIGatewayProxyResponse{
			Headers: map[string]string{"Content-Type": "application/json"},
			Body: err.Error(),
			StatusCode: 500,
		},nil
	}

	body, err := json.Marshal(todo)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Headers: map[string]string{"Content-Type": "application/json"},
			Body: err.Error(),
			StatusCode: 500,
		},nil
	}

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{"Content-Type": "application/json"},
		Body: string(body),
		StatusCode: 200,
	},nil

}

func (t *TodoHandler) FindTodoById(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse,error) {
	id := request.PathParameters["id"]
	todo, err := t.service.FindTodoById(id)

	if err != nil {
		return events.APIGatewayProxyResponse{
			Headers: map[string]string{"Content-Type": "application/json"},
			Body: err.Error(),
			StatusCode: 500,
		},nil
	}

	body, err := json.Marshal(todo)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Headers: map[string]string{"Content-Type": "application/json"},
			Body: err.Error(),
			StatusCode: 500,
		},nil
	}

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{"Content-Type": "application/json"},
		Body: string(body),
		StatusCode: 200,
	},nil
}

func (t *TodoHandler) DeleteTodo(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse,error) {
	id := request.PathParameters["id"]
	isDeleted, err := t.service.DeleteTodo (id)

	if err != nil {
		return events.APIGatewayProxyResponse{
			Headers: map[string]string{"Content-Type": "application/json"},
			Body: err.Error(),
			StatusCode: 500,
		},nil
	}

	body, err := json.Marshal(isDeleted)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Headers: map[string]string{"Content-Type": "application/json"},
			Body: err.Error(),
			StatusCode: 500,
		},nil
	}

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{"Content-Type": "application/json"},
		Body: fmt.Sprintf("todo delete status: %s", body),
		StatusCode: 200,
	},nil
}
