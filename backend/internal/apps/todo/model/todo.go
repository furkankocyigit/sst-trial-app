package todo

type Todo struct {
	ID          string `json:"id" dynamodbav:"id"`
	Title       string `json:"title" dynamodbav:"title"`
	Description string `json:"description" dynamodbav:"description"`
}
