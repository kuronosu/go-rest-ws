package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kuronosu/go-rest-ws/common/validations"
	"github.com/kuronosu/go-rest-ws/errors"
	"github.com/kuronosu/go-rest-ws/models"
	"github.com/kuronosu/go-rest-ws/repository"
	"github.com/kuronosu/go-rest-ws/server"
	"golang.org/x/crypto/bcrypt"
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

func ResponseError(w http.ResponseWriter, err error, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
}

func SignUpHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		validationResult, err := validations.ValidateSignupBody(r.Body)
		if err != nil {
			ResponseError(w, err, http.StatusBadRequest)
			return
		}

		if !validationResult.IsOk() {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(validationResult)
			return
		}

		email := validationResult.Email()
		password := validationResult.Password()

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			ResponseError(w, err, http.StatusInternalServerError)
			return
		}

		user := models.User{
			Email:    email,
			Password: string(hashedPassword),
		}
		err = repository.InsertUser(r.Context(), &user)
		if err != nil {
			if err, ok := err.(*errors.UserAlreadyExist); ok {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
				return
			}
			ResponseError(w, err, http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(SignUpResponse{
			Id:    user.Id,
			Email: user.Email,
		})
	}
}
