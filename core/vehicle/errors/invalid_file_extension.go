package errors

import "github.com/g3techlabs/revit-api/response"

func InvalidFileExtension() error {
	return &response.CustomError{
		Message:    "Invalid file extension. Only .jpg, .jpeg and .png are allowed.",
		StatusCode: 400,
	}
}
