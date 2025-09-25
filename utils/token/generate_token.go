package token

import (
	"time"

	"github.com/g3techlabs/revit-api/config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateTokens(id uint) (string, string, error) {
	var accessTokenSecret string = config.Get("ACCESS_SECRET")
	var refreshTokenSecret string = config.Get("REFRESH_SECRET")
	var accessTokenExpirationInHours int = config.GetIntVariable("ACCESS_TOKEN_EXPIRATION")
	var refreshTokenExpirationInHours int = config.GetIntVariable("REFRESH_TOKEN_EXPIRES_IN")

	accessTokenClaims := buildJwtClaims(id, accessTokenExpirationInHours)
	refreshTokenClaims := buildJwtClaims(id, refreshTokenExpirationInHours)

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString([]byte(accessTokenSecret))
	if err != nil {
		return "", "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(refreshTokenSecret))
	if err != nil {
		return "", "", nil
	}

	return accessTokenString, refreshTokenString, nil
}

func buildJwtClaims(userId uint, expirationInHours int) JwtClaims {
	return JwtClaims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expirationInHours))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
}
