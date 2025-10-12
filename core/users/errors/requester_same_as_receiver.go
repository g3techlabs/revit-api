package errors

import "github.com/g3techlabs/revit-api/response"

func RequesterSameAsReceiver() error {
	return &response.CustomError{
		Message:    "Requester is the same as the receiver user",
		StatusCode: 400,
	}
}
