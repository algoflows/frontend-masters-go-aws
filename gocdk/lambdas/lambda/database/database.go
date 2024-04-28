package database

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"lambda/types"
)

const (
	TABLE_NAME = "userTable"
)

type UserStorer interface {
	DoesUserExist(username string) (bool, error)
	InsertUser(user *types.User) error
	GetUser(username string) (types.User, error)
}

type DynamoDBClient struct {
	databaseStore *dynamodb.DynamoDB
}

func NewDynamoDBClient() *DynamoDBClient {
	dbSession := session.Must(session.NewSession())
	db := dynamodb.New(dbSession)
	return &DynamoDBClient{
		databaseStore: db,
	}
}

func (d *DynamoDBClient) DoesUserExist(username string) (bool, error) {
	// Check if the user exists in the database
	// If the user exists, return true
	// If the user does not exist, return false
	input := &dynamodb.GetItemInput{
		TableName: aws.String(TABLE_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(username),
			},
		},
	}

	result, err := d.databaseStore.GetItem(input)

	if err != nil {
		// If there is an error, return false and the error
		// This is a catch all for any other errors
		return true, err
	}

	if result.Item == nil {
		// The user does not exist
		// Return false
		return false, nil
	}

	return true, nil
}

func (d *DynamoDBClient) InsertUser(user *types.User) error {
	item := &dynamodb.PutItemInput{
		TableName: aws.String(TABLE_NAME),
		Item: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(user.Username),
			},
			"password": {
				S: aws.String(user.PasswordHash),
			},
		},
	}

	_, err := d.databaseStore.PutItem(item)

	if err != nil {
		return err
	}

	return nil
}

func (d *DynamoDBClient) GetUser(username string) (types.User, error) {
	user := types.User{}

	// construct input
	input := &dynamodb.GetItemInput{
		TableName: aws.String(TABLE_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(username),
			},
		},
	}

	// get item
	result, err := d.databaseStore.GetItem(input)

	// Error getting user
	if err != nil {
		return user, err
	}

	// user not found
	if result.Item == nil {
		return user, nil
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &user)

	if err != nil {
		return user, err
	}

	return user, nil
}
