package errors

import "github.com/g3techlabs/revit-api/src/response"

func RouteNotFound() error {
	return &response.CustomError{
		Message:    "Route not found",
		StatusCode: 404,
	}
}
