package errors

import "github.com/g3techlabs/revit-api/src/response"

func UserIsAlreadyMember() error {
	return &response.CustomError{
		Message:    "User is already member of this group",
		StatusCode: 409,
	}
}
