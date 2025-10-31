package middleware

import (
	"strings"

	"github.com/g3techlabs/revit-api/src/core/users/repository"
	"github.com/g3techlabs/revit-api/src/infra/token"
	"github.com/g3techlabs/revit-api/src/response/generics"
	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	userRepository repository.UserRepository
	tokenService   token.ITokenService
}

func NewAuthMiddleware(userRepository repository.UserRepository, tokenService token.ITokenService) *AuthMiddleware {
	return &AuthMiddleware{
		userRepository: userRepository, tokenService: tokenService,
	}
}

func (m *AuthMiddleware) Auth() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		accessToken, err := ExtractBearerToken(ctx)
		if err != nil {
			return err
		}

		claims, err := m.tokenService.ValidateAccessToken(accessToken)
		if err != nil {
			return generics.Unauthorized("Invalid or expired token")
		}

		userId := claims.UserID
		user, err := m.userRepository.FindUserById(userId)
		if err != nil {
			return generics.InternalError()
		} else if user == nil {
			return generics.Unauthorized("Not authenticated")
		}
		ctx.Locals("userId", user.ID)

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
