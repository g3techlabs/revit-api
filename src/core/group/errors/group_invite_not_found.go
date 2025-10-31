package errors

import "github.com/g3techlabs/revit-api/src/response"

func GroupInviteNotFound() error {
	return &response.CustomError{
		Message:    "Group invite not found",
		StatusCode: 404,
	}
}
