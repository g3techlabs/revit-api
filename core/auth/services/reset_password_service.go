package services

import (
	"fmt"

	"github.com/g3techlabs/revit-api/core/auth/dto"
	"github.com/g3techlabs/revit-api/core/users/repository"
	"github.com/g3techlabs/revit-api/utils"
	"github.com/g3techlabs/revit-api/utils/http"
	"github.com/g3techlabs/revit-api/utils/token"
	"github.com/gofiber/fiber/v2"
)

func ResetPasswordService(ctx *fiber.Ctx) error {
	var resetPasswordDto dto.ResetPassword

	if err := ctx.BodyParser(&resetPasswordDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body: " + err.Error(),
		})
	}

	if errs := utils.ValidateStruct(resetPasswordDto); len(errs) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errs,
		})
	}

	claims, err := token.ValidateResetPassToken(resetPasswordDto.ResetToken)
	if err != nil {
		return http.Unauthorized(ctx, "Invalid or expired token")
	}

	user, err := repository.FindUserById(claims.UserID)
	if err != nil {
		return http.InternalError(ctx)
	} else if user == nil {
		return http.Unauthorized(ctx, "User does not exist")
	}

	newPassword, err := hashPassword(resetPasswordDto.NewPassword)
	if err != nil {
		return http.InternalError(ctx)
	}

	if err = repository.UpdateUserPassword(claims.UserID, newPassword); err != nil {
		fmt.Println(err.Error())
		return http.InternalError(ctx)
	}

	return http.NoContent(ctx)
}
