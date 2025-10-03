package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (ts TokenService) GenerateAuthTokens(userId uint) (string, string, error) {
	accessTokenClaims := ts.buildAuthJwtClaims(userId, ts.accessTokenExpirationInHours)
	refreshTokenClaims := ts.buildAuthJwtClaims(userId, ts.refreshTokenExpirationInHours)

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString([]byte(ts.accessTokenSecret))
	if err != nil {
		return "", "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(ts.refreshTokenSecret))
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

func (ts TokenService) buildAuthJwtClaims(userId uint, expirationInHours int) JwtClaims {
	return JwtClaims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expirationInHours))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
}

func (ts TokenService) GenerateResetPassToken(userId uint) (string, error) {
	resetPassTokenClaims := ts.buildResetJwtClaims(userId, ts.resetPassTokenExpirationInMinutes)

	resetPassToken := jwt.NewWithClaims(jwt.SigningMethodHS256, resetPassTokenClaims)
	resetPassTokenString, err := resetPassToken.SignedString([]byte(ts.resetPasswordTokenSecret))
	if err != nil {
		return "", err
	}

	return resetPassTokenString, nil
}

func (ts TokenService) buildResetJwtClaims(userId uint, expirationInMinutes int) JwtClaims {
	return JwtClaims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(expirationInMinutes))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
}
