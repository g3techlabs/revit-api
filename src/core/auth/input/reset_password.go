package input

type ResetPassword struct {
	ResetToken  string `json:"resetToken" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required,password"`
}
