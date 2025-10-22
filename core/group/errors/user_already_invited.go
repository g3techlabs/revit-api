package errors

import "github.com/g3techlabs/revit-api/response"

func UserAlreadyInvitedOrMember() error {
	return &response.CustomError{
		Message:    "Invited user is already invited or a member",
		StatusCode: 409,
	}
}
