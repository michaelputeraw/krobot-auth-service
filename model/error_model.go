package model

type RequestError struct {
	Error interface{} `json:"error"`
}

type HttpCustomError struct {
	StatusCode int    `json:"code"`
	Message    string `json:"message"`
}

func (e HttpCustomError) Error() string {
	return e.Message
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type HttpCustomErrors struct {
	Message string            `json:"message"`
	Errors  []ValidationError `json:"errors"`
}

func NewHttpCustomError(statusCode int, err error) *HttpCustomError {
	return &HttpCustomError{
		StatusCode: statusCode,
		Message:    err.Error(),
	}
}
