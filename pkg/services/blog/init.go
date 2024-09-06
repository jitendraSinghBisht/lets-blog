package blog

import (
	"fmt"
	"lets-blog/config"
	models "lets-blog/pkg/models/blog"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func GetBlogUsingId(ddb *dynamodb.DynamoDB, blogId string) ( *models.BlogModel, error) {

	var blogData models.BlogModel

	key, err := dynamodbattribute.MarshalMap(map[string]string{"blogId": blogId})
	if err != nil {
		return nil, fmt.Errorf("failed to dynamoAttr marshal map: %w", err)
	}

	input := &dynamodb.GetItemInput{
		TableName: aws.String(config.GetDynamoUserTable()),
		Key:       key,
	}
	output, err := ddb.GetItem(input)
	if err != nil {
		return nil, fmt.Errorf("failed to get blog from db: %w", err)
	}

	if output.Item == nil || len(output.Item) == 0 {
		return nil, nil
	}

	err = dynamodbattribute.UnmarshalMap(output.Item, &blogData)
	if err != nil {
		return nil, fmt.Errorf("failed to dynamoAttr marshal map: %w", err)
	}

	return &blogData, nil
}

func GetBlogsUsingQuery(ddb *dynamodb.DynamoDB) ( *[]models.BlogsWithPagination, error) {

}
