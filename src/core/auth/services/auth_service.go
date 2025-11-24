package services

import (
	"github.com/g3techlabs/revit-api/src/core/auth/input"
	"github.com/g3techlabs/revit-api/src/core/auth/response"
	"github.com/g3techlabs/revit-api/src/core/users/repository"
	"github.com/g3techlabs/revit-api/src/infra/mail"
	"github.com/g3techlabs/revit-api/src/infra/token"
	"github.com/g3techlabs/revit-api/src/validation"
	"github.com/sirupsen/logrus"
)

type IAuthService interface {
	RegisterUser(input *input.CreateUser) (*response.UserCreatedResponse, error)
	Login(credentials *input.LoginCredentials) (*response.AuthTokensResponse, error)
	RefreshTokens(refreshToken string) (*response.AuthTokensResponse, error)
	ResetPassword(input *input.ResetPassword) error
	SendPassResetEmail(input *input.Identifier) error
}

type AuthService struct {
	userRepo     repository.UserRepository
	emailService mail.IEmailService
	tokenService token.ITokenService
	validator    validation.IValidator
	log          *logrus.Logger
}

func NewAuthService(validate validation.IValidator, userRepo repository.UserRepository, emailService mail.IEmailService, tokenService token.ITokenService) IAuthService {
	return &AuthService{validator: validate, userRepo: userRepo, emailService: emailService, tokenService: tokenService}
}
