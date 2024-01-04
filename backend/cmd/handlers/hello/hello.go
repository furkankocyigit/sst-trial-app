package main

import "github.com/aws/aws-lambda-go/lambda"

func HelloHandler() (string,error) {
	return "Hello from lambda",nil
}

func main(){
	lambda.Start(HelloHandler)
}