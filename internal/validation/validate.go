package validation

import (
	"github.com/sillkiw/gotube/internal/config"
)

type Validator struct {
	Cfg config.ValidationConfig
}

func New(cfg config.ValidationConfig) Validator {
	return Validator{Cfg: cfg}
}
