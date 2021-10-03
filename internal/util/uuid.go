package util

import (
	"net/http"

	custerror "github.com/lin-sel/baston-eventos/error"
	uuid "github.com/satori/go.uuid"
)

// ParseID Used to Parse Id From String
func ParseID(id string) (uuid.UUID, *custerror.CustomeError) {
	uid, err := uuid.FromString(id)
	if err != nil {
		er := custerror.CreateCustomeError("Failed to parse id", err, http.StatusBadRequest)
		return uid, &er
	}
	return uid, nil
}

// CreateID Used to Generate New ID
func CreateID() string {
	id := uuid.NewV4()
	return id.String()
}

// ValidateIDFormat Used to Validate ID Format
func ValidateIDFormat(id string) *custerror.CustomeError {
	_, err := ParseID(id)
	if err != nil {
		return err
	}
	return nil
}
