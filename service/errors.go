package service

import (
	"net/http"
)

type RestError struct {
	Status int   `json:"Status"`
	Error  error `json:"Error"`
}

func InternalServerError(err error) *RestError {
	return &RestError{
		Status: http.StatusInternalServerError,
		Error:  err,
	}
}
