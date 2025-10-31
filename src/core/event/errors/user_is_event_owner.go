package errors

import "github.com/g3techlabs/revit-api/src/response"

func UserIsEventOwner() error {
	return &response.CustomError{
		Message:    "User is event owner, forbidden operation",
		StatusCode: 403,
	}
}
