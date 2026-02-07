package validation

// Errors colect all validation error
type Errors map[string]string

func New() Errors {
	return make(map[string]string)
}

func (e Errors) Error() string {
	return "validation failed"
}

func (e Errors) Add(field, code string) {
	if code == "" {
		return
	}
	e[field] = code
}

func (e Errors) Empty() bool {
	return len(e) == 0
}
