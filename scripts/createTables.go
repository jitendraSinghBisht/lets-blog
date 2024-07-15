package scripts

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"lets-blog/config"
)

func CreateDynamoDBTables(cfg *config.Config) {
	log.Println("Script: CreateDynamoDBTables")

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

	svc := dynamodb.New(sess)

	createUserTable(svc)
	createBlogTable(svc)
}

func createUserTable(svc *dynamodb.DynamoDB) {
	input := &dynamodb.CreateTableInput{
		TableName: aws.String(config.GetDynamoUserTable()),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("userID"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("userName"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("userEmail"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("password"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("createdAt"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("userID"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("createdAt"),
				KeyType:       aws.String("RANGE"),
			},
		},
		GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex{
			{
				IndexName: aws.String(config.GetDynamoUserGsiTable()),
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String("userEmail"),
						KeyType:       aws.String("HASH"),
					},
					{
						AttributeName: aws.String("userName"),
						KeyType:       aws.String("RANGE"),
					},
				},
				Projection: &dynamodb.Projection{
					ProjectionType: aws.String("ALL"),
				},
				ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(5),
					WriteCapacityUnits: aws.Int64(5),
				},
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
	}

	_, err := svc.CreateTable(input)
	if err != nil {
		log.Fatalf("Failed to create user table: %v", err)
	}

	log.Println("Successfully created user table")
}

func createBlogTable(svc *dynamodb.DynamoDB) {
	input := &dynamodb.CreateTableInput{
		TableName: aws.String(config.GetDynamoBlogTable()),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("blogID"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("blogTitle"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("blogContent"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("createdBy"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("createdAt"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("blogID"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("createdAt"),
				KeyType:       aws.String("RANGE"),
			},
		},
		GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex{
			{
				IndexName: aws.String(config.GetDynamoBlogGsiTable()),
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String("createdBy"),
						KeyType:       aws.String("HASH"),
					},
					{
						AttributeName: aws.String("blogId"),
						KeyType:       aws.String("RANGE"),
					},
				},
				Projection: &dynamodb.Projection{
					ProjectionType: aws.String("ALL"),
				},
				ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(5),
					WriteCapacityUnits: aws.Int64(5),
				},
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
	}

	_, err := svc.CreateTable(input)
	if err != nil {
		log.Fatalf("Failed to create blog table: %v", err)
	}

	log.Println("Successfully created blog table")
}