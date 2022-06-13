package handlers

import (
	"net/http"

	"github.com/kuronosu/go-rest-ws/common/response"
	"github.com/kuronosu/go-rest-ws/server"
)

type HomeResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func HomeHandler(r *http.Request, s server.Server) response.IResponse {
	return response.NewJsonResponse(HomeResponse{
		Message: "Welcome",
		Status:  http.StatusOK,
	}, http.StatusOK, nil)
}
