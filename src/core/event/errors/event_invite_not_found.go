package errors

import "github.com/g3techlabs/revit-api/src/response"

func EventInviteNotFound() error {
	return &response.CustomError{
		Message:    "Event invite not found",
		StatusCode: 404,
	}
}
