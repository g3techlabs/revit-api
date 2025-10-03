package token

import "github.com/golang-jwt/jwt/v5"

type JwtClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}
