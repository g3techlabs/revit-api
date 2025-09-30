package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAuthTokens(id uint) (string, string, error) {
	accessTokenClaims := buildAuthJwtClaims(id, accessTokenExpirationInHours)
	refreshTokenClaims := buildAuthJwtClaims(id, refreshTokenExpirationInHours)

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString([]byte(accessTokenSecret))
	if err != nil {
		return "", "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(refreshTokenSecret))
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

func buildAuthJwtClaims(userId uint, expirationInHours int) JwtClaims {
	return JwtClaims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expirationInHours))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
}

func GenerateResetPassToken(userId uint, expiration int) (string, error) {
	resetPassTokenClaims := buildResetJwtClaims(userId, expiration)

	resetPassToken := jwt.NewWithClaims(jwt.SigningMethodHS256, resetPassTokenClaims)
	resetPassTokenString, err := resetPassToken.SignedString([]byte(resetPasswordTokenSecret))
	if err != nil {
		return "", err
	}

	return resetPassTokenString, nil
}

func buildResetJwtClaims(userId uint, expirationInMinutes int) JwtClaims {
	return JwtClaims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(expirationInMinutes))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
}
