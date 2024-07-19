package user

import (
	"lets-blog/config"
	
	"lets-blog/pkg/utils"
	model "lets-blog/pkg/models/user"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func GetUserUsingEmail(userEmail string) (*model.UserModel, error) {
	awsClient := utils.ConnectToDynamodb()

	input := &dynamodb.QueryInput{
		TableName:              aws.String(config.GetDynamoUserTable()),
		IndexName:              aws.String(config.GetDynamoUserGsiTable()),
		KeyConditionExpression: aws.String("#gsiKey = :gsiValue"),
		ExpressionAttributeNames: map[string]*string{
			"#gsiKey": aws.String("userEmail"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":gsiValue": {
				S: aws.String(userEmail),
			},
		},
	}

	result, err := awsClient.Query(input)
	if err != nil {
		return nil, err
	}
	var user *model.UserModel
	for _, item := range result.Items{
		if err = dynamodbattribute.UnmarshalMap(item, &user);err != nil {
			return nil, err
		}
	}
	return user, nil
}

func CreateNewUser(user model.UserModel) error {
	awsClient := utils.ConnectToDynamodb()

	data, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		TableName: aws.String(config.GetDynamoUserTable()),
		Item:      data,
	}
	if _, err = awsClient.PutItem(input); err != nil {
		return err
	}
	return nil
}