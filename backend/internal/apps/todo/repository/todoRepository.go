package repository

import (
	model "backend/internal/apps/todo/model"
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
)

type TodoRepositoryInterface interface {
	GetAllTodos() ([]model.Todo, error)
	CreateTodo(model.Todo) (model.Todo, error)
	FindTodoById(string) (model.Todo, error)
	DeleteTodo(string) (bool, error)
}


type TodoRepository struct {
	TableName string
	dynamoDbClient *dynamodb.Client
	
}

func NewTodoRepository() *TodoRepository {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = "us-east-1"
		return nil
	})

	if err != nil {
		panic(err)
	}
	// create dynamodb client
	dynamoDbClient := dynamodb.NewFromConfig(cfg)
	tableName := "FKocyigi-sst-trial-app-todoTable"  //dev table name	
	return &TodoRepository{ TableName: tableName, dynamoDbClient: dynamoDbClient}
}

func (l *TodoRepository) GetAllTodos() ([]model.Todo, error) {
	resp,err := l.dynamoDbClient.Scan(context.TODO(),&dynamodb.ScanInput{
		TableName: aws.String(l.TableName),
	})
	if err != nil {
		log.Printf("Got error calling Scan: %s", err)
		panic(err)
	}

	if len(resp.Items) == 0 {
		return nil, errors.New("No todos found")
	}

	todos := []model.Todo{}
	err = attributevalue.UnmarshalListOfMaps(resp.Items, &todos)
	if err != nil {
		log.Printf("Got error unmarshalling: %s", err)
	}
	return todos, nil
}

func (l *TodoRepository) CreateTodo(todo model.Todo) (model.Todo, error) {
	item,err := attributevalue.MarshalMap(todo)
	if err != nil {
		log.Printf("Got error marshalling new todo item: %s", err)
		panic( err)
	}

	_, err = l.dynamoDbClient.PutItem(context.TODO(),&dynamodb.PutItemInput{
		TableName: aws.String(l.TableName),
		Item: item,
	})

	if err != nil {
		log.Printf("Couldn't add item to table. Here's why: %v\n", err)
	}

	return todo ,err
}


func (l *TodoRepository) FindTodoById(id string) (model.Todo, error) {
	response,err := l.dynamoDbClient.GetItem(context.TODO(),&dynamodb.GetItemInput{
		TableName: aws.String(l.TableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},	
	})

	if err != nil {
		log.Printf("Got error calling GetItem: %s", err)
		panic(err)
	}
	if len(response.Item) == 0 {
		return model.Todo{}, fmt.Errorf("No item found for ID: %s", id)
	}
	todo := model.Todo{}
	err = attributevalue.UnmarshalMap(response.Item, &todo)
	if err != nil {
		log.Printf("Got error unmarshalling: %s", err)
	}
	return todo, nil
}


func (l *TodoRepository) DeleteTodo (id string) (bool, error) {
	_,err := l.dynamoDbClient.DeleteItem(context.TODO(),&dynamodb.DeleteItemInput{
		TableName: aws.String(l.TableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},	
	})

	if err != nil {
		log.Printf("Got error calling DeleteItem: %s", err)
		panic(err)
	}

	return true, nil
}