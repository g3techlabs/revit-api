package errors

import "github.com/g3techlabs/revit-api/response"

func UsersAreTheSame() error {
	return &response.CustomError{
		Message:    "Requester and invite target are the same",
		StatusCode: 400,
	}
}
