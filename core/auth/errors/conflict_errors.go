package errors

import (
	"github.com/g3techlabs/revit-api/response"
)

func NewConflictError(emailTaken, nicknameTaken bool) error {
	errs := make(map[string]string)
	if emailTaken {
		errs["email"] = "already taken"
	}
	if nicknameTaken {
		errs["nickname"] = "already taken"
	}
	if len(errs) == 0 {
		return nil
	}
	return &response.CustomError{
		StatusCode: 409,
		Details:    errs,
	}
}
