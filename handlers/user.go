package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kuronosu/go-rest-ws/errors"
	"github.com/kuronosu/go-rest-ws/models"
	"github.com/kuronosu/go-rest-ws/repository"
	"github.com/kuronosu/go-rest-ws/server"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
}

func SignUpHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var requestData SignUpRequest
		err := json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil {
			w.Header().Set("X-Content-Type-Options", "nosniff")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{
				Message: "Invalida data",
			})
			return
		}

		user := models.User{
			Email:    requestData.Email,
			Password: requestData.Password,
		}
		err = repository.InsertUser(r.Context(), &user)
		if err != nil {
			if err, ok := err.(*errors.UserAlreadyExist); ok {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(SignUpResponse{
			Id:    user.Id,
			Email: user.Email,
		})
	}
}
