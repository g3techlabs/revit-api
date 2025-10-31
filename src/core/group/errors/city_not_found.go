package errors

import "github.com/g3techlabs/revit-api/src/response"

func CityNotFound() error {
	return &response.CustomError{
		Message:    "Specified city not found",
		StatusCode: 404,
	}
}
