package input

type EmailInput struct {
	Email string `json:"email" validate:"required,email"`
}
