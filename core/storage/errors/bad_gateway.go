package errors

import "github.com/g3techlabs/revit-api/response"

func BadGateway() error {
	return &response.CustomError{
		Message:    "Bad gateway",
		StatusCode: 502,
	}
}
