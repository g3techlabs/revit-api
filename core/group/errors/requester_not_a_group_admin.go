package errors

import "github.com/g3techlabs/revit-api/response"

func RequesterIsNotAnAdmin() error {
	return &response.CustomError{
		Message:    "Requester is not an admin",
		StatusCode: 403,
	}
}
