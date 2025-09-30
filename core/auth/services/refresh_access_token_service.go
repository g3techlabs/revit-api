package services

import (
	"github.com/g3techlabs/revit-api/core/auth/middleware"
	"github.com/g3techlabs/revit-api/core/users/repository"
	"github.com/g3techlabs/revit-api/utils/http"
	"github.com/g3techlabs/revit-api/utils/token"
	"github.com/gofiber/fiber/v2"
)

func RefreshTokens(ctx *fiber.Ctx) error {
	refreshToken, err := middleware.ExtractBearerToken(ctx)
	if err != nil {
		return http.Unauthorized(ctx, err.Error())
	}

	claims, err := token.ValidateRefreshToken(refreshToken)
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

	accessToken, refreshToken, err := token.GenerateAuthTokens(userId)
	if err != nil {
		return http.InternalError(ctx)
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}
