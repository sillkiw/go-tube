package httpjson

import (
	"net/http"

	"github.com/go-chi/render"
)

func WriteJSON(w http.ResponseWriter, r *http.Request, code int, body any) {
	render.Status(r, code)
	render.JSON(w, r, body)
}
