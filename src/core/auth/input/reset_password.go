package input

// ResetPassword representa os dados para redefinição de senha
// @Description Dados necessários para redefinir a senha do usuário
type ResetPassword struct {
	// Token de redefinição de senha recebido por email
	ResetToken string `json:"resetToken" validate:"required" example:"abc123def456"`
	// Nova senha do usuário
	NewPassword string `json:"newPassword" validate:"required,password" example:"novaSenha123"`
}
