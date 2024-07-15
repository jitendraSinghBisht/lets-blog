package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func NewRoutes(ddb *dynamodb.DynamoDB) http.Handler {
	router := chi.NewRouter()

	return router
}