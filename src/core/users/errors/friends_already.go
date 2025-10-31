package errors

import "github.com/g3techlabs/revit-api/src/response"

func FriendsAlready() error {
	return &response.CustomError{
		Message:    "Users either have a pending friendship request or are friends already",
		StatusCode: 409,
	}
}
