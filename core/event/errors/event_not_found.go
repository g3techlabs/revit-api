package errors

import "github.com/g3techlabs/revit-api/response"

func EventNotFound() error {
	return &response.CustomError{
		Message:    "Event not found",
		StatusCode: 404,
	}
}
