package httpjson

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

func WriteJSON(w http.ResponseWriter, r *http.Request, code int, body any) {
	render.Status(r, code)
	render.JSON(w, r, body)
}

func WriteError(w http.ResponseWriter, r *http.Request, code int, errCode string, message string) {
	errBody := ErrorResponse{
		Error:   errCode,
		Message: message,
	}
	WriteJSON(w, r, code, errBody)
}
