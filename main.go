package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	// Recupera o ID do cliente a partir do corpo da requisição
	customerID := request.Body

	// Cria uma sessão para o DynamoDB
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	if err != nil {
		return nil, err
	}

	// Cria um novo serviço do DynamoDB
	svc := dynamodb.New(sess)

	// Consulta o banco de dados para obter as informações do cliente com base no ID
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("customers"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(customerID),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	// Converte o resultado para uma string
	resultString := fmt.Sprintf("%v", result.Item)

	// Retorna as informações do cliente como resposta
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       resultString,
	}, nil
}
