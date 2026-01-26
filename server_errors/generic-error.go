package servererrors

type RequestError struct {
	Code    int
	Message string
}

func (e *RequestError) Error() string {
	return e.Message
}
