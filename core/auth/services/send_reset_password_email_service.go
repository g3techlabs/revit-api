package services

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/g3techlabs/revit-api/config"
	"github.com/g3techlabs/revit-api/core/auth/dto"
	"github.com/g3techlabs/revit-api/core/mail/emails"
	"github.com/g3techlabs/revit-api/core/reset_token/repository"
	"github.com/g3techlabs/revit-api/db/models"
	"github.com/g3techlabs/revit-api/utils"
	"github.com/g3techlabs/revit-api/utils/http"
	"github.com/gofiber/fiber/v2"
)

var resetTokenExpirationInMinutes = config.GetIntVariable("RESET_TOKEN_EXPIRATION")

func SendResetPasswordEmailService(ctx *fiber.Ctx) error {
	var identifierDto dto.Identifier

	if err := ctx.BodyParser(identifierDto); err != nil {
		return http.BadRequest(ctx, "Invalid request body"+err.Error())
	}

	if errors := utils.ValidateStruct(identifierDto); len(errors) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errors,
		})
	}

	user, err := findUserByIdentifier(identifierDto.Identifier)
	if err != nil {
		return http.InternalError(ctx)
	} else if user == nil {
		return ctx.SendStatus(fiber.StatusNoContent)
	}

	token, err := generateToken()
	if err != nil {
		return http.InternalError(ctx)
	}

	resetToken := models.ResetToken{
		UserID:    user.ID,
		Token:     token,
		ExpiresAt: time.Now().Add(time.Minute * time.Duration(resetTokenExpirationInMinutes)),
	}
	if err := repository.RegisterResetToken(&resetToken); err != nil {
		return http.InternalError(ctx)
	}

	if err := emails.SendResetPasswordEmailService(user.Email, user.Name, token, resetTokenExpirationInMinutes); err != nil {
		return http.InternalError(ctx)
	}

	return http.NoContent(ctx)
}

func generateToken() (string, error) {
	max := big.NewInt(1000000)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}
	return fmt.Sprint("%06d", n.Int64()), nil
}
