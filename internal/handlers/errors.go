package handlers

import (
	"errors"
)

var (
	ErrInvalidYear        = errors.New("invalid value for year")
	ErrInvalidMonth       = errors.New("A Month specifies a month of the year (January = 1, ...)")
	ErrInvalidPhoneNumber = errors.New("phone number cant be null")
)

type errorResponse struct {
	Message string `json:"message"`
}

func getErrorResponseBody(err error) errorResponse {
	return errorResponse{
		Message: err.Error(),
	}
}
