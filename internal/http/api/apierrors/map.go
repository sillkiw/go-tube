package apierrors

import (
	"errors"
	"net/http"

	apivalid "github.com/sillkiw/gotube/internal/http/api/validation"
)

func Map(err error) (int, ErrorBody) {
	if errBody, ok := ValidationMap(err); ok {
		return http.StatusBadRequest, errBody
	}
	return http.StatusInternalServerError, New("internal_error", "internal server error")
}

func ValidationMap(err error) (ErrorBody, bool) {
	var verrs apivalid.Errors
	if errors.As(err, &verrs) {
		return ErrorBody{
			Code:    "validation_error",
			Message: "validation failed",
			Fields:  map[string]string(verrs),
		}, true
	}
	return ErrorBody{}, false
}
