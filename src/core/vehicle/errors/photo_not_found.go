package errors

import "github.com/g3techlabs/revit-api/src/response"

func PhotoNotFound() error {
	return &response.CustomError{
		Message:    "Photo not found",
		StatusCode: 404,
	}
}
