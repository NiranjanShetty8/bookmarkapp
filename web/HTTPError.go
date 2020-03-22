package web

//Implements error
type HTTPError struct {
	HTTPStatus int
	ErrorKey   string
}

//Error method to implement error
func (err HTTPError) Error() string {
	return err.ErrorKey
}

//returns instance of HTTPError
func NewHTTPError(err string, statusCode int) *HTTPError {
	return &HTTPError{
		HTTPStatus: statusCode,
		ErrorKey:   err,
	}
}
