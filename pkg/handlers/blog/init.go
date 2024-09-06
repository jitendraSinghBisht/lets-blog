package blog

import (
	"net/http"
	"encoding/json"
	service "lets-blog/pkg/services/blog"

	"github.com/go-chi/chi"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func HandleGetBlogUsingBlogId(ddb *dynamodb.DynamoDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		blogID := chi.URLParam(r, "blogID")

		res, err := service.GetBlogUsingId(ddb, blogID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resJson, err := json.Marshal(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resJson)
	}
}

func HandleGetBlogsUsingQuery(ddb *dynamodb.DynamoDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()

		category := queryParams.Get("startKey")
		// category := queryParams.Get("status")
	}
}

// Create a success response JSON object
// type Response struct {
// 	Message string `json:"message"`
// }
// resp := Response{Message: "User signed up successfully"}

// // Marshal response object to JSON
// respJSON, err := json.Marshal(resp)
// if err != nil {
// 	http.Error(w, err.Error(), http.StatusInternalServerError)
// 	return
// }

// // Set Content-Type and write response
// w.Header().Set("Content-Type", "application/json")
// w.WriteHeader(http.StatusOK)
// w.Write(respJSON)
