package errors

import "github.com/g3techlabs/revit-api/response"

func UserNotFound() error {
	return &response.CustomError{
		Message:    "User not found",
		StatusCode: 404,
	}
}
