package models

import "encoding/json"

type APIError struct {
	Code    int
	Message string
}

func NewAPIError(code int, message string) *APIError {
	return &APIError{
		Code:    code,
		Message: message,
	}
}

func (a *APIError) Jsonify() []byte {
	result, _ := json.Marshal(a)
	return result
}
