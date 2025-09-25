package middleware

import (
	"strings"

	"github.com/g3techlabs/revit-api/config"
	"github.com/g3techlabs/revit-api/core/auth/dto"
	"github.com/g3techlabs/revit-api/core/users/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		accessToken, err := extractBearerToken(ctx)
		if err != nil {
			return unauthorized(ctx, err.Error())
		}

		claims, err := validateJwt(accessToken)
		if err != nil {
			return unauthorized(ctx, "Invalid or expired token")
		}

		userId := claims.UserID
		user, err := repository.FindUserById(userId)
		if err != nil {
			return internalError(ctx)
		} else if user == nil {
			return unauthorized(ctx, "Not authenticated")
		}

		return ctx.Next()
	}
}

func extractBearerToken(ctx *fiber.Ctx) (string, error) {
	authHeader := ctx.Get("Authorization")

	if authHeader == "" {
		return "", fiber.NewError(fiber.StatusUnauthorized, "Missing Authorization header")
	}

	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) != 2 || strings.ToLower(bearerToken[0]) != "bearer" {
		return "", fiber.NewError(fiber.StatusUnauthorized, "Invalid Authorization header format")
	}

	return bearerToken[1], nil
}

func validateJwt(tokenString string) (*dto.JwtClaims, error) {
	var accessSecret = []byte(config.Get("ACCESS_SECRET"))

	token, err := jwt.ParseWithClaims(tokenString, &dto.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}
		return accessSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims := token.Claims.(*dto.JwtClaims)
	return claims, nil
}

func unauthorized(ctx *fiber.Ctx, message string) error {
	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": message,
	})
}

func internalError(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusInternalServerError)
}
