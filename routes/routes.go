package routes

import (
	"net/http"

	user "lets-blog/routes/user"
	blog "lets-blog/routes/blog"

	"github.com/go-chi/chi"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func NewRoutes(ddb *dynamodb.DynamoDB) http.Handler {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("All okay"))
		return
	})

	router.Route("/user", loadUserRoutes(ddb))
	router.Route("/blog", loadBlogRoutes(ddb))

	return router
}

func loadUserRoutes(db *dynamodb.DynamoDB) func(chi.Router) {
	return user.UserRoutes(db)
}

func loadBlogRoutes(db *dynamodb.DynamoDB) func(chi.Router) {
	return blog.BlogRoutes(db)
}