package handlers

import (
	"net/http"

	"github.com/kuronosu/go-rest-ws/common/response"
	"github.com/kuronosu/go-rest-ws/common/validations"
	"github.com/kuronosu/go-rest-ws/errors"
	"github.com/kuronosu/go-rest-ws/models"
	"github.com/kuronosu/go-rest-ws/repository"
	"github.com/kuronosu/go-rest-ws/server"
	"golang.org/x/crypto/bcrypt"
)

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
}

func SignUpHandler(r *http.Request, s server.Server) response.IResponse {
	validationResult, err := validations.ValidateSignupBody(r.Body)
	if err != nil {
		return response.ErrorJsonResponse(err.Error(), http.StatusBadRequest, nil)
	}

	if !validationResult.IsOk() {
		return response.ErrorJsonResponse(validationResult, http.StatusBadRequest, nil)
		// return response.NewJsonResponse(validationResult, http.StatusBadRequest, nil)
	}

	email := validationResult.Email()
	password := validationResult.Password()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return response.ErrorJsonResponse(err.Error(), http.StatusInternalServerError, nil)
	}

	user := models.User{
		Email:    email,
		Password: string(hashedPassword),
	}
	err = repository.InsertUser(r.Context(), &user)
	if err != nil {
		code := http.StatusInternalServerError
		if _, ok := err.(*errors.UserAlreadyExist); ok {
			code = http.StatusBadRequest
		}
		return response.ErrorJsonResponse(err.Error(), code, nil)
	}
	return response.NewJsonResponse(SignUpResponse{
		Id:    user.Id,
		Email: user.Email,
	}, http.StatusCreated, nil)
}
