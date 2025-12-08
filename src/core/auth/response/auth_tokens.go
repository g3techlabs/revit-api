package response

type AuthTokensResponse struct {
	AccessToken   string  `json:"accessToken" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	RefreshToken  string  `json:"refreshToken" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	ID            uint    `json:"id" example:"1"`
	ProfilePicUrl *string `json:"profilePicUrl" example:"https://example.com/users/1/profile.jpg"`
	Name          string  `json:"name" example:"João Silva"`
	Nickname      string  `json:"nickname" example:"joaosilva"`
}
