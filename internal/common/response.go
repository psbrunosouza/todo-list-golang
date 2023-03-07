package common

type Responser interface {
	Response(data interface{}, ok bool, status int, errorMessage error) *Response
}

type Response struct {
	Data   interface{}
	Ok     bool
	Status int
	Error  error
}

func (response *Response) Response(data interface{}, ok bool, status int, errorMessage error) *Response {
	return &Response{
		Data:   data,
		Ok:     ok,
		Status: status,
		Error:  errorMessage,
	}
}
