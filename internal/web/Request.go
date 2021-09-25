package web

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	custerror "github.com/lin-sel/baston-eventos/error"
)

func RequestParse(r *http.Request, target interface{}) error {
	if r.Body == nil {
		return custerror.CreateCustomeError("Request Body is Empty", errors.New(""),
			http.StatusBadRequest)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return custerror.CreateCustomeError("Failed to Read Request Body", err,
			http.StatusBadRequest)
	}

	if len(body) == 0 {
		return custerror.CreateCustomeError("Request Body is Empty", errors.New(""),
			http.StatusBadRequest)
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		return custerror.CreateCustomeError("Failed to Unmarshal Request Body", errors.New(""),
			http.StatusInternalServerError)
	}
	return nil
}
