package service

import (
	"github.com/g3techlabs/revit-api/core/storage"
	"github.com/g3techlabs/revit-api/core/users/input"
	"github.com/g3techlabs/revit-api/core/users/repository"
	"github.com/g3techlabs/revit-api/core/users/response"
	"github.com/go-playground/validator/v10"
)

type IUserService interface {
	Update(id uint, input *input.UpdateUser) error
	PresignProfilePic(id uint, input *input.PresignProfilePic) (*response.ProfilePicPresignedURL, error)
	UpdateProfilePic(id uint, input *input.UpdateProfilePic) error
}

type UserService struct {
	userRepo       repository.UserRepository
	storageService storage.StorageService
	validator      *validator.Validate
}

func NewUserService(validator *validator.Validate, userRepo repository.UserRepository, storageService storage.StorageService) IUserService {
	return &UserService{
		userRepo:       userRepo,
		storageService: storageService,
		validator:      validator,
	}
}
