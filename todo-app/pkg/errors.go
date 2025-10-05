package pkg

import (
	"encoding/json"
	"errors"
	"time"
)

var ErrValidation = errors.New("validation error -")

func ErrorAddition(err error, additionText string) error {

	return errors.Join(err, errors.New(additionText))
}

type HTTPError struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

func HttpError(err error) string {
	er := HTTPError{
		Message: err.Error(),
		Time:    time.Now(),
	}

	data, e := json.Marshal(er)

	if e != nil {
		return ""
	}

	return string(data)
}
