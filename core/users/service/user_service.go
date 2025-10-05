package service

import (
	"github.com/g3techlabs/revit-api/core/users/input"
	"github.com/g3techlabs/revit-api/core/users/repository"
	"github.com/go-playground/validator/v10"
)

type IUserService interface {
	Update(id uint, input *input.UpdateUser) error
}

type UserService struct {
	userRepo  repository.UserRepository
	validator *validator.Validate
}

func NewUserService(validator *validator.Validate, userRepo repository.UserRepository) IUserService {
	return &UserService{
		userRepo:  userRepo,
		validator: validator,
	}
}
