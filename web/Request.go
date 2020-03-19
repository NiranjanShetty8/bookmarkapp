package web

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

//UnmarshalJSON checks for empty body and then parses JSON into the target
//target should be a pointer
func UnmarshalJSON(r *http.Request, target interface{}) error {
	if r.Body == nil {
		return errors.New("There is problem while reading data")
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errors.New("Can't handle data")
	}

	if len(body) == 0 {
		return errors.New("Empty Data")
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		return errors.New("Unavle to Parse Data")
	}
	return nil
}
