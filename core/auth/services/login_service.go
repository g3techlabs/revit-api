package services

import (
	"net/mail"
	"time"

	"github.com/g3techlabs/revit-api/config"
	"github.com/g3techlabs/revit-api/core/auth/dto"
	"github.com/g3techlabs/revit-api/core/users/repository"
	"github.com/g3techlabs/revit-api/db/models"
	"github.com/g3techlabs/revit-api/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(ctx *fiber.Ctx) error {
	var loginDTO dto.Login

	if err := ctx.BodyParser(&loginDTO); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body" + err.Error(),
		})
	}

	if errors := utils.ValidateStruct(loginDTO); len(errors) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errors,
		})
	}

	user, err := findUserByIdentifier(loginDTO.Identifier)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	} else if user == nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"errors": "Invalid credentials",
		})
	}

	if !checkPasswordHash(loginDTO.Password, user.Password) {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"errors": "Invalid credentials",
		})
	}

	accessToken, refreshToken, err := generateTokens(user.ID)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.JSON(fiber.Map{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}

func isAnEmail(value string) bool {
	_, err := mail.ParseAddress(value)
	return err == nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func findUserByIdentifier(identifier string) (*models.User, error) {
	isIdentifierAnEmail := isAnEmail(identifier)

	user, err := new(models.User), *new(error)
	if isIdentifierAnEmail {
		user, err = repository.FindUserByEmail(identifier)

	} else {
		user, err = repository.FindUserByNickname(identifier)
	}

	return user, err
}

func generateTokens(id uint) (string, string, error) {
	var accessTokenSecret string = config.Get("ACCESS_SECRET")
	var refreshTokenSecret string = config.Get("REFRESH_SECRET")
	var accessTokenExpirationInHours int = config.GetIntVariable("ACCESS_TOKEN_EXPIRATION")
	var refreshTokenExpirationInHours int = config.GetIntVariable("REFRESH_TOKEN_EXPIRES_IN")

	accessTokenClaims := buildJwtClaims(id, accessTokenExpirationInHours)
	refreshTokenClaims := buildJwtClaims(id, refreshTokenExpirationInHours)

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString([]byte(accessTokenSecret))
	if err != nil {
		return "", "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(refreshTokenSecret))
	if err != nil {
		return "", "", nil
	}

	return accessTokenString, refreshTokenString, nil
}

func buildJwtClaims(userId uint, expirationInHours int) dto.JwtClaims {
	return dto.JwtClaims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expirationInHours))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
}
