package services

import (
	"github.com/g3techlabs/revit-api/core/auth/input"
	"github.com/g3techlabs/revit-api/core/auth/response"
	"github.com/g3techlabs/revit-api/core/mail"
	"github.com/g3techlabs/revit-api/core/token"
	usersInput "github.com/g3techlabs/revit-api/core/users/input"
	"github.com/g3techlabs/revit-api/core/users/repository"
	usersResponse "github.com/g3techlabs/revit-api/core/users/response"
	"github.com/go-playground/validator/v10"
)

type IAuthService interface {
	RegisterUser(input *usersInput.CreateUser) (*usersResponse.UserCreatedResponse, error)
	Login(credentials *input.LoginCredentials) (*response.AuthTokensResponse, error)
	RefreshTokens(refreshToken string) (*response.AuthTokensResponse, error)
	ResetPassword(input *input.ResetPassword) error
	SendPassResetEmail(input *input.Identifier) error
}

type AuthService struct {
	userRepo     repository.UserRepository
	emailService mail.IEmailService
	tokenService token.ITokenService
	validator    *validator.Validate
}

func NewAuthService(validate *validator.Validate, userRepo repository.UserRepository, emailService mail.IEmailService, tokenService token.ITokenService) IAuthService {
	return &AuthService{validator: validate, userRepo: userRepo, emailService: emailService, tokenService: tokenService}
}
