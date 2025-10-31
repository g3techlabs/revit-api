package errors

import "github.com/g3techlabs/revit-api/src/response"

func UserNotAdmin() error {
	return &response.CustomError{
		Message:    "User is not an admin of this group",
		StatusCode: 403,
	}
}
