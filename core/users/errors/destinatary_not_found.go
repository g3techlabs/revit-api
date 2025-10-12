package errors

import "github.com/g3techlabs/revit-api/response"

func DestinataryNotFound() error {
	return &response.CustomError{
		Message:    "Destinatary was not found",
		StatusCode: 404,
	}
}
