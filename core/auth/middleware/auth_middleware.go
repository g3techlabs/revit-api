package middleware

import (
	"strings"

	"github.com/g3techlabs/revit-api/core/users/repository"
	"github.com/g3techlabs/revit-api/utils/http"
	"github.com/g3techlabs/revit-api/utils/token"
	"github.com/gofiber/fiber/v2"
)

func JWTAuth() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		accessToken, err := ExtractBearerToken(ctx)
		if err != nil {
			return http.Unauthorized(ctx, err.Error())
		}

		claims, err := token.ValidateAccessToken(accessToken)
		if err != nil {
			return http.Unauthorized(ctx, "Invalid or expired token")
		}

		userId := claims.UserID
		user, err := repository.FindUserById(userId)
		if err != nil {
			return http.InternalError(ctx)
		} else if user == nil {
			return http.Unauthorized(ctx, "Not authenticated")
		}

		return ctx.Next()
	}
}

func ExtractBearerToken(ctx *fiber.Ctx) (string, error) {
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
