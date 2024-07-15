package routes

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/go-chi/chi"
)

func UserRoutes(ddb *dynamodb.DynamoDB) func(chi.Router) {
	return func(r chi.Router) {}
}
