package response

// AuthTokensResponse representa os tokens de autenticação retornados após login ou refresh
// @Description Tokens de autenticação JWT
type AuthTokensResponse struct {
	// Token de acesso JWT (válido por tempo limitado)
	AccessToken string `json:"accessToken" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	// Token de refresh JWT (usado para obter novos tokens)
	RefreshToken string `json:"refreshToken" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}
