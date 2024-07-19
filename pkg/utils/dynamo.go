package utils

import (
	"lets-blog/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func ConnectToDynamodb() *dynamodb.DynamoDB {
	cfg := config.GetInstance()
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,

		Config: aws.Config{
			Region:   &cfg.AwsConfig.Region,
			Endpoint: &cfg.DynamoEndpoint,
			Credentials: credentials.NewStaticCredentials(
				cfg.AwsConfig.AccessKey,
				cfg.AwsConfig.SecretKey,
				"",
			),
		},
	}))

	return dynamodb.New(sess)
}

