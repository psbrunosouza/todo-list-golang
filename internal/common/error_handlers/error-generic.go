package error_handlers

type GenericError struct {
	Message interface{}
	Status  int
}

func NewGenericError(status int, message interface{}) GenericError {
	return GenericError{
		Message: message,
		Status:  status,
	}
}
