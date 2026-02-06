package apierrors

import (
	"errors"
	"net/http"

	apivalid "github.com/sillkiw/gotube/internal/http/api/validation"
	"github.com/sillkiw/gotube/internal/storage"
)

func Map(err error) (int, ErrorBody) {
	if errBody, ok := ValidationMap(err); ok {
		return http.StatusBadRequest, errBody
	}
	if errBody, ok := StorageMap(err); ok {
		return http.StatusConflict, errBody
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

func StorageMap(err error) (ErrorBody, bool) {
	if errors.Is(err, storage.ErrTitleExists) {
		return ErrorBody{
			Code:    "exist_error",
			Message: "title exists",
		}, true
	}
	return ErrorBody{}, false
}
