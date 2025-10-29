package errors

import "github.com/g3techlabs/revit-api/response"

func UserIsNotSubscribed() error {
	return &response.CustomError{
		Message:    "User is not subscribed in this event",
		StatusCode: 400,
	}
}
