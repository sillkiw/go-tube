package apierrors

import (
	"errors"
	"fmt"
	"net/http"

	apivalid "github.com/sillkiw/gotube/internal/http/api/validation"
	"github.com/sillkiw/gotube/internal/storage"
)

func Map(err error) (int, ErrorBody) {
	if errBody, ok := ValidationMap(err); ok {
		return http.StatusBadRequest, errBody
	}
	if code, errBody, ok := StorageMap(err); ok {
		return code, errBody
	}
	return http.StatusInternalServerError, New("internal_error", "internal server error")
}

func ValidationMap(err error) (ErrorBody, bool) {
	var verrs apivalid.Errors
	fmt.Println(err)
	if errors.As(err, &verrs) {
		return ErrorBody{
			Code:    "validation_error",
			Message: "validation failed",
			Fields:  map[string]string(verrs),
		}, true
	}
	return ErrorBody{}, false
}

func StorageMap(err error) (int, ErrorBody, bool) {
	if errors.Is(err, storage.ErrTitleExists) {
		return http.StatusConflict, ErrorBody{
			Code:    "exist_error",
			Message: "title exists",
		}, true
	}
	if errors.Is(err, storage.ErrIdNotFound) {
		return http.StatusNotFound, ErrorBody{
			Code:    "not_found",
			Message: "id not found",
		}, true
	}
	return 0, ErrorBody{}, false
}
