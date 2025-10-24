package errors

import "github.com/g3techlabs/revit-api/response"

func CityNotFound() error {
	return &response.CustomError{
		Message:    "Specified city not found",
		StatusCode: 404,
	}
}
