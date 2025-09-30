package token

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateAccessToken(tokenString string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}
		return []byte(accessTokenSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims := token.Claims.(*JwtClaims)
	return claims, nil
}

func ValidateRefreshToken(tokenString string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}
		return []byte(refreshTokenSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims := token.Claims.(*JwtClaims)
	return claims, nil
}

func ValidateResetPassToken(tokenString string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}
		return []byte(resetPasswordTokenSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims := token.Claims.(*JwtClaims)
	return claims, nil
}
