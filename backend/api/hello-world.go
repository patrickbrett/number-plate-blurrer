// package main

// import (
// 	"context"
// 	"fmt"

// 	"github.com/aws/aws-lambda-go/events"
// 	"github.com/aws/aws-lambda-go/lambda"
// )

// type MyEvent struct {
// 	queryStringParameters struct {
// 		name string
// 	}
// }

// func HandleRequest(ctx context.Context, event events.APIGatewayProxyRequest) (interface{}, error) {
// 	message := fmt.Sprintf("Hello %s!", event.QueryStringParameters["name"])

// 	// return {
// 	// 	status: 200,
// 	// 	body: message
// 	// }

// 	m := map[string]interface{}{"statusCode": 200, "body": message}

// 	return m, nil
// }

// func main() {
// 	lambda.Start(HandleRequest)
// }
