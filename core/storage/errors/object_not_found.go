package errors

import "github.com/g3techlabs/revit-api/response"

func ObjectNotFound() error {
	return &response.CustomError{
		Message:    "Object was not found",
		StatusCode: 404,
	}
}
