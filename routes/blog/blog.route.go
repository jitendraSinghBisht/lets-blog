package routes

import (
	handler "lets-blog/pkg/handlers/blog"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/go-chi/chi"
)

func BlogRoutes(ddb *dynamodb.DynamoDB) func(chi.Router) {
	return func(r chi.Router) {
		r.Get("/{blogId}", handler.HandleGetBlogUsingBlogId(ddb))
		r.Get("",handler.HandleGetBlogsUsingQuery(ddb))
	}
}
