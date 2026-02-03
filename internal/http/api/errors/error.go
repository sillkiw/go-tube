package apierrors

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/sillkiw/gotube/internal/validation"
)

const (
	titleRequired = "title_required"
	invalidTitle = "invalid_title"
	invalidSize = "invalid_size"
)



func CreateRequestValidation(err error, v validation.Validator) (code int, errCode string, message string) {
	switch {
	case errors.Is(err, validation.ErrTitleRequired):
		return http.StatusBadRequest, titleRequired, "title is required"
	case errors.Is(err, validation.ErrTitleToShort):
		msg := fmt.Sprintf("title lengthg must be >= %d symbols", v.Cfg.Title.MinLen)
		return http.StatusBadRequest, invalidTitle, msg
	case errors.Is(err, validation.ErrTitleToLong):
		msg := fmt.Sprintf("title length must be <= %d symbols", v.Cfg.Title.MaxLen)
		return http.StatusBadRequest, invalidTitle, msg
	case errors.Is(err, validation.ErrLargeSize):
		msg := fmt.Sprintf("video size must be <= %d bytes", v.Cfg.UplLimit.MaxSize)
		return http.StatusBadRequest, invalidSize, msg 
		
	
}
