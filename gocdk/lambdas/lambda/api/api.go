package api

import (
	"context"
	"encoding/json"
	"fmt"
	"lambda/auth"
	"lambda/database"
	"lambda/helpers"
	"lambda/types"

	"github.com/aws/aws-lambda-go/events"
)

type ApiHandler struct {
	dbStore database.UserStorer
}

func NewApiHandler(dbsStore database.UserStorer) *ApiHandler {
	return &ApiHandler{
		dbStore: dbsStore,
	}
}

func (a *ApiHandler) RegisterUserHandler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var registerUser types.RegisterUser

	err := json.Unmarshal([]byte(request.Body), &registerUser)
	if err != nil {
		return helpers.Response(400, "Invalid request"), fmt.Errorf("error unmarshaling request: %w", err)
	}

	// Validate the request
	if registerUser.Username == "" || registerUser.Password == "" {
		return helpers.Response(400, "Invalid request missing fields"), nil
	}

	// does user exist already?
	exists, err := a.dbStore.DoesUserExist(registerUser.Username)

	if err != nil {
		return helpers.Response(500, "Internal server error"), fmt.Errorf("error checking user existence: %w", err)
	}

	if exists {
		return helpers.Response(409, "Username already exists"), nil
	}

	user, err := auth.GeneratePassword(registerUser)
	if err != nil {
		return helpers.Response(500, "Internal server error"), fmt.Errorf("error generating password: %w", err)
	}
	// Insert the user into the database
	err = a.dbStore.InsertUser(user) // Dereference the pointer to pass the user value

	// Handle any errors
	if err != nil {
		return helpers.Response(500, "Internal server error"), fmt.Errorf("error inserting user into database: %w", err)
	}

	// Return the response
	return helpers.Response(201, "User created"), nil
}

func (a *ApiHandler) LoginUserHandler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var loginUser types.User

	err := json.Unmarshal([]byte(request.Body), &loginUser)
	if err != nil {
		return helpers.Response(400, "Invalid request"), fmt.Errorf("error unmarshaling request: %w", err)
	}

	// Validate the request
	if loginUser.Username == "" || loginUser.PasswordHash == "" {
		return helpers.Response(400, "Invalid request missing fields"), nil
	}

	// does user exist already?
	exists, err := a.dbStore.DoesUserExist(loginUser.Username)
	if err != nil {
		return helpers.Response(500, "Internal server error"), fmt.Errorf("error checking user existence: %w", err)
	}

	if !exists {
		return helpers.Response(404, "User not found"), fmt.Errorf("error user not found")
	}

	user, err := a.dbStore.GetUser(loginUser.Username)
	if err != nil {
		return helpers.Response(500, "Internal server error"), fmt.Errorf("error getting user: %w", err)
	}

	isValid := auth.ValidatePassword(loginUser.PasswordHash, user.PasswordHash)
	if !isValid {
		return helpers.Response(401, "Invalid password"), fmt.Errorf("error invalid password")
	}

	if isValid {
		return &events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       user.Username,
		}, nil
	}

	// if anything else returns an error
	return helpers.Response(500, "Internal server error"), nil
}
