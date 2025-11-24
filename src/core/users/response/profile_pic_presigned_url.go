package response

// ProfilePicPresignedURL representa a URL pré-assinada para upload de foto de perfil
// @Description URL pré-assinada e chave para upload de foto de perfil do usuário
type ProfilePicPresignedURL struct {
	// URL pré-assinada para upload da foto de perfil
	PresignedURL string `json:"presignedUrl" example:"https://s3.amazonaws.com/bucket/users/123/profile-pic.jpg?X-Amz-Signature=..."`
	// Chave da foto no storage
	ObjectKey string `json:"objectKey" example:"users/123/profile-pic.jpg"`
}
