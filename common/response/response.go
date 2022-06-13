package response

import (
	"encoding/json"
	"net/http"
)

type IResponse interface {
	Write(http.ResponseWriter) error
	WriteHeaders(http.ResponseWriter)
}
type BaseResponse[T any] struct {
	StatusCode int
	Headers    http.Header
	Body       T
}

func (res *BaseResponse[T]) WriteHeaders(w http.ResponseWriter) {
	w.WriteHeader(res.StatusCode)
	for key := range res.Headers {
		w.Header().Set(key, res.Headers.Get(key))
	}

}

type JsonResponse[T any] struct {
	BaseResponse[T]
}

func (res *JsonResponse[T]) Write(w http.ResponseWriter) error {
	res.WriteHeaders(w)
	return json.NewEncoder(w).Encode(res.Body)
}

func NewJsonResponse[T any](body T, statusCode int, herders http.Header) *JsonResponse[T] {
	if herders == nil {
		herders = http.Header{}
	}
	herders.Set("Content-Type", "application/json")
	return &JsonResponse[T]{
		BaseResponse[T]{
			StatusCode: statusCode,
			Headers:    herders,
			Body:       body,
		},
	}
}

type ErrorResponseBody struct {
	Message string `json:"message"`
}

func ErrorJsonResponse(
	body interface{},
	statusCode int,
	herders http.Header,
) *JsonResponse[any] {
	if herders == nil {
		herders = http.Header{}
	}
	herders.Set("Content-Type", "application/json")
	switch message := body.(type) {
	case string:
		body = ErrorResponseBody{message}
	}
	return &JsonResponse[any]{
		BaseResponse[any]{
			StatusCode: statusCode,
			Headers:    herders,
			Body:       body,
		},
	}
}
