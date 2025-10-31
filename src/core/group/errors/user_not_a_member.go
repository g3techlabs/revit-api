package errors

import "github.com/g3techlabs/revit-api/src/response"

func UserIsNotAMember() error {
	return &response.CustomError{
		Message:    "User is not a member",
		StatusCode: 400,
	}
}
