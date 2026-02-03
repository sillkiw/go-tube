package validation

import "errors"

var (
	ErrTitleRequired       = errors.New("title is required")
	ErrTitleToLong         = errors.New("title longer than available")
	ErrTitleToShort        = errors.New("title shorter than available")
	ErrLargeSize           = errors.New("video size is larger than available")
	ErrSmallSize           = errors.New("video size is smaller than available")
	ErrInvalidSize         = errors.New("size is invalid")
	ErrNotAllowedContent   = errors.New("content type is not allowed")
	ErrContentTypeRequired = errors.New("content type is required")
)
