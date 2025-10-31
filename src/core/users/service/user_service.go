package service

import (
	"github.com/g3techlabs/revit-api/src/core/users/input"
	"github.com/g3techlabs/revit-api/src/core/users/repository"
	"github.com/g3techlabs/revit-api/src/core/users/response"
	"github.com/g3techlabs/revit-api/src/infra/storage"
	"github.com/g3techlabs/revit-api/src/validation"
)

type IUserService interface {
	Update(id uint, input *input.UpdateUser) error
	RequestProfilePicUpdate(id uint, input *input.RequestProfilePicUpdate) (*response.ProfilePicPresignedURL, error)
	ConfirmNewProfilePic(id uint, input *input.ConfirmNewProfilePic) error
	GetUsers(params *input.GetUsersQuery) (*[]response.GetUserResponse, error)
	GetUser(userId uint) (*response.GetUserResponse, error)
	GetFriends(userId uint, params *input.GetUsersQuery) (*[]response.Friend, error)
	RequestFriendship(userId, destinataryId uint) error
	AnswerFriendshipRequest(userId, requesterId uint, answer *input.FriendshipRequestAnswer) (*response.FriendshipRequestAnswered, error)
	RemoveFriendship(userId, friendId uint) error
	GetFriendshipRequests(userId uint, query *input.GetFriendshipRequestsQuery) (*[]response.FriendshipRequest, error)
}

type UserService struct {
	userRepo       repository.UserRepository
	storageService storage.StorageService
	validator      validation.IValidator
}

func NewUserService(validator validation.IValidator, userRepo repository.UserRepository, storageService storage.StorageService) IUserService {
	return &UserService{
		userRepo:       userRepo,
		storageService: storageService,
		validator:      validator,
	}
}
