package service

import (
	"github.com/g3techlabs/revit-api/core/users/errors"
	"github.com/g3techlabs/revit-api/core/users/response"
	"github.com/g3techlabs/revit-api/response/generics"
)

func (us *UserService) GetUser(userId uint) (*response.GetUserResponse, error) {
	user, err := us.userRepo.FindUserById(userId)
	if err != nil {
		us.Log.Errorf("Error searching for user %d", userId)
		return nil, generics.InternalError()
	} else if user == nil {
		return nil, errors.UserNotFound("User not found")
	}

	response := user.ToGetUserResponse()
	return response, nil
}
