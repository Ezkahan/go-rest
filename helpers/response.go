package helpers

import "strings"

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

type EmptyObj struct{}

func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Data:    data,
		Errors:  nil,
	}

	return res
}

func BuildErrorResponse(message string, errors string, data interface{}) Response {
	splittedErrors := strings.Split(errors, "\n")

	res := Response{
		Status:  false,
		Message: message,
		Errors:  splittedErrors,
		Data:    data,
	}

	return res
}
