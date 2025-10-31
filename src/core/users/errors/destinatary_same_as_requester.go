package errors

import "github.com/g3techlabs/revit-api/src/response"

func RequesterAndDestinataryAreTheSame() error {
	return &response.CustomError{
		Message:    "Requester and destinatary are the same",
		StatusCode: 400,
	}
}
