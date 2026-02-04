package apierrors

// ErrorBody is a response, if something goes wrong
type ErrorBody struct {
	Code    string            `json:"code"`
	Message string            `json:"message,omitempty"`
	Fields  map[string]string `json:"fields,omitempty"`
}

func New(code string, message string) ErrorBody {
	return ErrorBody{Code: code, Message: message}
}
