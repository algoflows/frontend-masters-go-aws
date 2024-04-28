package main

import (
	"context"
	"lambda/app"
	"lambda/helpers"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	myApp := app.NewApp()
	lambda.Start(func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		switch request.Path {
		case "/register":
			response, err := myApp.ApiHandler.RegisterUserHandler(ctx, request)
			if err != nil {
				return *helpers.Response(500, "Internal Server Error"), err
			}
			return *response, nil
		case "/login":
			response, err := myApp.ApiHandler.LoginUserHandler(ctx, request)
			if err != nil {
				return *helpers.Response(500, "Internal Server Error"), err
			}
			return *response, nil
		default:
			return *helpers.Response(404, "Not Found"), nil
		}
	})
}
