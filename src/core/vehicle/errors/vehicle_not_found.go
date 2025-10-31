package errors

import "github.com/g3techlabs/revit-api/src/response"

func VehicleNotFound() error {
	return &response.CustomError{
		Message:    "Vehicle not found",
		StatusCode: 404,
	}
}
