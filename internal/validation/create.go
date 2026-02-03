package validation

import (
	"fmt"
	"unicode/utf8"
)

type ValidatableCreateRequest interface {
	GetTitle() string
	GetContentType() string
	GetSize() int64
}

func (v *Validator) CreateRequest(req ValidatableCreateRequest) error {
	if err := v.title(req.GetTitle()); err != nil {
		return fmt.Errorf("title: %w", err)
	}
	if err := v.contentType(req.GetContentType()); err != nil {
		return fmt.Errorf("content type: %w", err)
	}
	if err := v.size(req.GetSize()); err != nil {
		return fmt.Errorf("size: %w", err)
	}
	return nil
}

func (v *Validator) title(title string) error {
	lenght := utf8.RuneCountInString(title)
	if lenght == 0 {
		return ErrTitleRequired
	}
	if lenght > v.Cfg.Title.MaxLen {
		return ErrTitleToLong
	}
	if lenght < v.Cfg.Title.MinLen {
		return ErrTitleToShort
	}
	return nil
}

func (v *Validator) contentType(cntType string) error {
	if cntType == "" {
		return ErrContentTypeRequired
	}
	for _, t := range v.Cfg.UplLimit.AllowedContent {
		if cntType == t {
			return nil
		}
	}
	return ErrNotAllowedContent
}

func (v *Validator) size(size int64) error {
	if size <= 0 {
		return ErrInvalidSize
	}
	if size > v.Cfg.UplLimit.MaxSize {
		return ErrLargeSize
	}
	if size < v.Cfg.UplLimit.MinSize {
		return ErrSmallSize
	}
	return nil
}
