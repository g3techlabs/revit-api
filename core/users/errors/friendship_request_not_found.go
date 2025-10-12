package errors

import "github.com/g3techlabs/revit-api/response"

func FriendshipRequestNotFound() error {
	return &response.CustomError{
		Message:    "The pending friendship request was not found",
		StatusCode: 400,
	}
}
