package token

import (
	"github.com/g3techlabs/revit-api/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateAccessToken(tokenString string) (*JwtClaims, error) {
	var accessSecret = []byte(config.Get("ACCESS_SECRET"))

	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}
		return accessSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims := token.Claims.(*JwtClaims)
	return claims, nil
}

func ValidateRefreshToken(tokenString string) (*JwtClaims, error) {
	var refreshSecret = []byte(config.Get("REFRESH_SECRET"))

	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}
		return refreshSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims := token.Claims.(*JwtClaims)
	return claims, nil
}
