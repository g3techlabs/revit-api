package services

import (
	"fmt"
	"strings"

	"github.com/g3techlabs/revit-api/config"
	"github.com/g3techlabs/revit-api/core/auth/dto"
	"github.com/g3techlabs/revit-api/core/mail/emails"
	"github.com/g3techlabs/revit-api/utils"
	"github.com/g3techlabs/revit-api/utils/http"
	"github.com/g3techlabs/revit-api/utils/token"
	"github.com/gofiber/fiber/v2"
)

var resetTokenExpirationInMinutes int = config.GetIntVariable("RESET_TOKEN_EXPIRATION")
var appName = config.Get("APP_NAME")

func SendResetPasswordEmailService(ctx *fiber.Ctx) error {
	var identifierDto dto.Identifier

	if err := ctx.BodyParser(&identifierDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body: " + err.Error(),
		})
	}

	if errs := utils.ValidateStruct(identifierDto); len(errs) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errs,
		})
	}

	user, err := findUserByIdentifier(identifierDto.Identifier)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"erro": "erro ao pegar usuario",
		})
	} else if user == nil {
		return http.NoContent(ctx)
	}

	deepLink, err := generateDeepLink(user.ID)
	if err != nil {
		return http.InternalError(ctx)
	}

	if err := emails.SendResetPasswordEmailService(user.Email, user.Name, deepLink, resetTokenExpirationInMinutes); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"erro": "erro ao enviar email: " + err.Error(),
		})
	}

	return http.NoContent(ctx)
}

func generateDeepLink(userId uint) (string, error) {
	resetPassToken, err := token.GenerateResetPassToken(userId, resetTokenExpirationInMinutes)
	if err != nil {
		return "", nil
	}

	deepLink := fmt.Sprintf("%s://reset_password?t=%s", strings.ToLower(appName), resetPassToken)

	return deepLink, nil
}
