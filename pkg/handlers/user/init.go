package user

import (
	"encoding/json"
	"net/http"

	model "lets-blog/pkg/models/user"
	userServices "lets-blog/pkg/services/user"
)

func HandleUserSignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rqBody model.UserReqModel
		err := json.NewDecoder(r.Body).Decode(&rqBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if (rqBody.UserName == "" || rqBody.UserEmail == "" || rqBody.Password == ""){
			http.Error(w, "User Name, Email and password are required.", http.StatusBadRequest)
			return
		}
		if err = userServices.SignUpUser(&rqBody); err != nil {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			return
		}
	}
}

func HandleUserLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func HandleUserLogout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func HandleGetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
