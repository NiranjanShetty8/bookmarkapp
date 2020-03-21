package web

import (
	"encoding/json"
	"net/http"
)

func writeToHeader(w *http.ResponseWriter, statusCode int, payload interface{}) {
	(*w).WriteHeader(statusCode)
	(*w).Write(payload.([]byte))
}

//Set response with status code
func RespondJSON(w *http.ResponseWriter, statusCode int, content interface{}) {
	response, err := json.Marshal(content)
	if err != nil {
		writeToHeader(w, http.StatusInternalServerError, err.Error())
		return
	}
	(*w).Header().Set("Content-Type", "application/json")
	writeToHeader(w, statusCode, response)
}

//RespondErrorMessage Writes Error to Respond Writer
func RespondErrorMessage(w *http.ResponseWriter, code int, message string) {
	RespondJSON(w, code, message)
}

func RespondError(w *http.ResponseWriter, err error) {
	switch err.(type) {
	case ValidationError, *ValidationError:
		err, _ := err.(*ValidationError)
		RespondJSON(w, http.StatusBadRequest, err.ErrorKey+": "+err.Errors["error"])
	case HTTPError, *HTTPError:
		httpError := err.(*HTTPError)
		RespondErrorMessage(w, httpError.HTTPStatus, httpError.ErrorKey)
	default:
		RespondErrorMessage(w, http.StatusInternalServerError, err.Error())
	}
}
