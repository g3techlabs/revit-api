package errors

import "github.com/g3techlabs/revit-api/src/response"

func RequesterIsNotAnAdmin() error {
	return &response.CustomError{
		Message:    "Requester is not an admin of this group",
		StatusCode: 403,
	}
}
