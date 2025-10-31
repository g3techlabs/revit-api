package token

import "github.com/g3techlabs/revit-api/src/config"

type ITokenService interface {
	GenerateAuthTokens(userId uint) (string, string, error)
	GenerateResetPassToken(userId uint) (string, error)
	ValidateAccessToken(tokenString string) (*JwtClaims, error)
	ValidateRefreshToken(tokenString string) (*JwtClaims, error)
	ValidateResetPassToken(tokenString string) (*JwtClaims, error)
}

type TokenService struct {
	accessTokenSecret                 string
	refreshTokenSecret                string
	resetPasswordTokenSecret          string
	accessTokenExpirationInHours      int
	refreshTokenExpirationInHours     int
	resetPassTokenExpirationInMinutes int
}

func NewTokenService() ITokenService {
	return &TokenService{
		accessTokenSecret:                 config.Get("ACCESS_SECRET"),
		refreshTokenSecret:                config.Get("REFRESH_SECRET"),
		resetPasswordTokenSecret:          config.Get("RESET_PASSWORD_SECRET"),
		accessTokenExpirationInHours:      config.GetIntVariable("ACCESS_TOKEN_EXPIRATION"),
		refreshTokenExpirationInHours:     config.GetIntVariable("REFRESH_TOKEN_EXPIRATION"),
		resetPassTokenExpirationInMinutes: config.GetIntVariable("RESET_TOKEN_EXPIRATION"),
	}
}
