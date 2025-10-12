package service

import (
	"github.com/g3techlabs/revit-api/core/storage"
	"github.com/g3techlabs/revit-api/core/users/input"
	"github.com/g3techlabs/revit-api/core/users/repository"
	"github.com/g3techlabs/revit-api/core/users/response"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type IUserService interface {
	Update(id uint, input *input.UpdateUser) error
	RequestProfilePicUpdate(id uint, input *input.RequestProfilePicUpdate) (*response.ProfilePicPresignedURL, error)
	ConfirmNewProfilePic(id uint, input *input.ConfirmNewProfilePic) error
	GetUsers(params *input.GetUsersQuery) (*[]response.GetUserResponse, error)
	GetUser(userId uint) (*response.GetUserResponse, error)
	RequestFriendship(userId, destinataryId uint) error
	AnswerFriendshipRequest(userId, requesterId uint, answer *input.FriendshipRequestAnswer) (*response.FriendshipRequestAnswered, error)
}

type UserService struct {
	userRepo       repository.UserRepository
	storageService storage.StorageService
	validator      *validator.Validate
	Log            *logrus.Logger
}

func NewUserService(validator *validator.Validate, userRepo repository.UserRepository, storageService storage.StorageService) IUserService {
	return &UserService{
		userRepo:       userRepo,
		storageService: storageService,
		validator:      validator,
	}
}
