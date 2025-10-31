package token

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func (ts TokenService) ValidateAccessToken(tokenString string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}
		return []byte(ts.accessTokenSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims := token.Claims.(*JwtClaims)
	return claims, nil
}

func (ts TokenService) ValidateRefreshToken(tokenString string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}
		return []byte(ts.refreshTokenSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims := token.Claims.(*JwtClaims)
	return claims, nil
}

func (ts TokenService) ValidateResetPassToken(tokenString string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}
		return []byte(ts.resetPasswordTokenSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims := token.Claims.(*JwtClaims)
	return claims, nil
}
