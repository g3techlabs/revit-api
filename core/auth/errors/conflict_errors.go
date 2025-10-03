package errors

type ConflictErrors struct {
	Errors map[string]string
}

func (e *ConflictErrors) Error() string {
	return "Conflict errors"
}

func NewConflictErrors(errs map[string]string) *ConflictErrors {
	return &ConflictErrors{Errors: errs}
}
