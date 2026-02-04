package videosvalidation

import (
	"unicode/utf8"

	apivalid "github.com/sillkiw/gotube/internal/http/api/validation"
	"github.com/sillkiw/gotube/internal/http/api/videos/dto"
)

func (v *Validator) CreateRequest(req dto.CreateRequest) apivalid.Errors {
	errs := apivalid.Errors{}
	v.title(errs, req.Title)
	v.contentType(errs, req.ContentType)
	v.size(errs, req.Size)

	return errs
}

func (v *Validator) title(errs apivalid.Errors, title string) {
	lenght := utf8.RuneCountInString(title)
	if lenght == 0 {
		errs.Add("title", "required")
	} else {
		if lenght > v.Cfg.Title.MaxLen {
			errs.Add("title", "too_long")
		}
		if lenght < v.Cfg.Title.MinLen {
			errs.Add("title", "too_short")
		}
	}
}

func (v *Validator) contentType(errs apivalid.Errors, cntType string) {
	if cntType == "" {
		errs.Add("content_type", "required")
	}
	for _, t := range v.Cfg.UplLimit.AllowedContent {
		if cntType == t {
			return
		}
	}
	errs.Add("content_type", "not_allowed")
}

func (v *Validator) size(errs apivalid.Errors, size int64) {
	if size <= 0 {
		errs.Add("size", "invalid")
	} else {
		if size > v.Cfg.UplLimit.MaxSize {
			errs.Add("size", "too_large")
		}
		if size < v.Cfg.UplLimit.MinSize {
			errs.Add("size", "too_small")
		}
	}

}
