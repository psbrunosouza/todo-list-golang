package common

type AppError struct {
	Message interface{}
	Status  int
}

// NewAppError provides a new app error pattern
func NewAppError(status int, message interface{}) *AppError {
	return &AppError{
		Message: message,
		Status:  status,
	}
}
