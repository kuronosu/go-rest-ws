package views

import (
	"net/http"

	"github.com/kuronosu/go-rest-ws/common/response"
	"github.com/kuronosu/go-rest-ws/server"
)

type Responder func(r *http.Request, s server.Server) response.IResponse

type IView interface {
	Responder() Responder
}

type FunctionView Responder

func (view FunctionView) Responder() Responder {
	return Responder(view)
}

func HandleView(view IView, s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		view.Responder()(r, s).Write(w)
	}
}

func HandleResponder(responder Responder, s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		responder(r, s).Write(w)
	}
}
