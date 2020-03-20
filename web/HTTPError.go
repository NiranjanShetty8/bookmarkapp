package web

type HTTPError struct {
	HTTPStatus int
	ErrorKey   string
}

func (err HTTPError) Error() string {
	return err.ErrorKey
}

func NewHTTPError(err string, statusCode int) *HTTPError {
	return &HTTPError{
		HTTPStatus: statusCode,
		ErrorKey:   err,
	}
}
