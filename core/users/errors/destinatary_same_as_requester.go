package errors

import "github.com/g3techlabs/revit-api/response"

func DestinatarySameAsRequester() error {
	return &response.CustomError{
		Message:    "Destinatary is the same as the user",
		StatusCode: 400,
	}
}
