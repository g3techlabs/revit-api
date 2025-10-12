package errors

import "github.com/g3techlabs/revit-api/response"

func FriendsAlready() error {
	return &response.CustomError{
		Message:    "Users are friends already",
		StatusCode: 409,
	}
}
