package errors

import "github.com/g3techlabs/revit-api/src/response"

func GroupNotFound() error {
	return &response.CustomError{
		Message:    "Group not found",
		StatusCode: 404,
	}
}
