package dto

import (
	"encoding/json"
	"time"
)

type OperationBody struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

type ErrorObj struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

func NewErrorObj(err error) string {

	e := ErrorObj{
		Message: err.Error(),
		Time:    time.Now(),
	}

	data, err := json.Marshal(e)

	if err != nil {
		return err.Error()
	}

	return string(data)
}
