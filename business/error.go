package business

import "errors"


type ErrorMessage struct{
	Message error		`json:"message"`
	Data	interface{}	`json:"data,omitempty"`
}
var (
	ErrIDNotFound = errors.New("id not found")
	ErrBadRequest = errors.New("bad request")
)