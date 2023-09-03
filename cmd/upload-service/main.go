package main

import (
	"context"
	"stream-x/internal/server"

	ddlambda "github.com/DataDog/datadog-lambda-go"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	service, err := server.New()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error initialzing server",
		}, err
	}
	return service.UploadVideo(ctx, request)
}

func main() {
	lambda.Start(ddlambda.WrapFunction(Handler, nil))
}
