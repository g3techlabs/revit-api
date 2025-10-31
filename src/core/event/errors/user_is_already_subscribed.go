package errors

import "github.com/g3techlabs/revit-api/src/response"

func UserIsAlreadySubscribed() error {
	return &response.CustomError{
		Message:    "User is already subscribed to this event",
		StatusCode: 409,
	}
}
