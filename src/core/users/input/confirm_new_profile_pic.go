package input

// ConfirmNewProfilePic representa os dados para confirmar upload de nova foto de perfil
// @Description Chave da foto que foi enviada com sucesso para confirmar o upload
type ConfirmNewProfilePic struct {
	// Chave da foto no storage
	ObjectKey string `json:"objectKey" validate:"required" example:"users/123/profile-pic.jpg"`
}
