package web

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	custerror "github.com/lin-sel/baston-eventos/error"
)

func RequestParse(r *http.Request, target interface{}) *custerror.CustomeError {
	if r.Body == nil {
		err := custerror.CreateCustomeError("Request Body is Empty", errors.New(""),
			http.StatusBadRequest)
		return &err
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err := custerror.CreateCustomeError("Failed to Read Request Body", err,
			http.StatusBadRequest)
		return &err
	}

	if len(body) == 0 {
		err := custerror.CreateCustomeError("Request Body is Empty", errors.New(""),
			http.StatusBadRequest)
		return &err
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		err := custerror.CreateCustomeError("Failed to Unmarshal Request Body", errors.New(""),
			http.StatusInternalServerError)
		return &err
	}
	return nil
}
