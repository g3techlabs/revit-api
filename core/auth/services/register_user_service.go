package services

import (
	"fmt"

	"github.com/g3techlabs/revit-api/core/users/dto"
	usersRepository "github.com/g3techlabs/revit-api/core/users/repository"
	"github.com/g3techlabs/revit-api/db/models"
	"github.com/g3techlabs/revit-api/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type User = models.User

func RegisterUser(ctx *fiber.Ctx) error {
	var createUserDTO dto.CreateUser

	if err := ctx.BodyParser(&createUserDTO); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body:" + err.Error(),
		})
	}

	if errors := utils.ValidateStruct(createUserDTO); len(errors) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errors,
		})
	}

	if nicknameErr, emailErr := verifyUniqueFieldsAvailability(createUserDTO.Nickname, createUserDTO.Email); nicknameErr != nil || emailErr != nil {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"errors": utils.CollectErrors(nicknameErr, emailErr),
		})
	}

	hashedPassword, err := hashPassword(createUserDTO.Password)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	createUserDTO.Password = hashedPassword

	user := parseDtoToUserEntity(createUserDTO)
	err = usersRepository.RegisterUser(&user)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	userResponse := parseEntityToUserResponse(user)

	return ctx.Status(fiber.StatusCreated).JSON(userResponse)
}

func verifyUniqueFieldsAvailability(nickname string, email string) (error, error) {
	var emailError error
	var nicknameError error

	if !isNicknameAvailable(nickname) {
		nicknameError = fmt.Errorf("nickname already in use")
	}

	if !isEmailAvailable(email) {
		emailError = fmt.Errorf("email already in use")
	}

	return nicknameError, emailError
}

func isNicknameAvailable(nickname string) bool {
	nicknameAlreadyInUse, _ := usersRepository.FindUserByNickname(nickname)
	return nicknameAlreadyInUse == nil
}

func isEmailAvailable(email string) bool {
	emailAlreadyInUse, _ := usersRepository.FindUserByEmail(email)
	return emailAlreadyInUse == nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	return string(bytes), err
}

func parseDtoToUserEntity(dto dto.CreateUser) User {
	user := User{
		Name:     dto.Name,
		Email:    dto.Email,
		Nickname: dto.Nickname,
		Password: dto.Password,
	}

	return user
}

func parseEntityToUserResponse(user User) dto.UserCreatedResponse {
	userResponse := dto.UserCreatedResponse{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		Nickname:   user.Nickname,
		ProfilePic: user.ProfilePic,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.CreatedAt,
	}

	return userResponse
}
