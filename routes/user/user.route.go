package routes

import (
	handler "lets-blog/pkg/handlers/user"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/go-chi/chi"
)

func UserRoutes(ddb *dynamodb.DynamoDB) func(chi.Router) {
	return func(r chi.Router) {
		r.Get("/{userId}", handler.HandleGetUser())
		r.Post("/sign-up", handler.HandleUserSignUp())
		r.Post("/login", handler.HandleUserLogin())
		r.Get("/logout", handler.HandleUserLogout())
	}
}
