package errors

import (
	"net/http"
)

// RestErr common struct of error response
type RestErr struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Error      string `json:"error"`
}

// BadRequestErr is fucntion to return BAD_REQUEST response
func BadRequestErr(message string) *RestErr {
	return &RestErr{
		Message:    message,
		StatusCode: http.StatusBadRequest,
		Error:      "BAD_REQUEST",
	}
}

// NotFoundErr is fucntion to return BAD_REQUEST response
func NotFoundErr(message string) *RestErr {
	return &RestErr{
		Message:    message,
		StatusCode: http.StatusNotFound,
		Error:      "NOT_FOUND",
	}
}

// AlreadyExistErr is fucntion to return BAD_REQUEST response
func ConflictErr(message string) *RestErr {
	return &RestErr{
		Message:    message,
		StatusCode: http.StatusConflict,
		Error:      "CONFLICT",
	}
}
