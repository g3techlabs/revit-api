package errors

import "github.com/g3techlabs/revit-api/response"

func GroupNotFound() error {
	return &response.CustomError{
		Message:    "Group not found",
		StatusCode: 404,
	}
}
