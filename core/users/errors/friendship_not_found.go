package errors

import "github.com/g3techlabs/revit-api/response"

func FriendshipNotFound() error {
	return &response.CustomError{
		Message:    "Friendship between the specified users was not found",
		StatusCode: 404,
	}
}
