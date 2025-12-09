package service

import (
	"math"

	"github.com/g3techlabs/revit-api/src/core/users/input"
	"github.com/g3techlabs/revit-api/src/core/users/response"
	"github.com/g3techlabs/revit-api/src/response/generics"
	"github.com/g3techlabs/revit-api/src/utils"
)

func (us *UserService) GetUsers(params *input.GetUsersQuery) (*response.GetUsersResponse, error) {
	if err := us.validator.Validate(params); err != nil {
		return nil, err
	}

	users, totalCount, err := us.userRepo.GetUsers(params.Page, params.Limit, params.Nickname)
	if err != nil {
		return nil, generics.InternalError()
	}

	limit := 20
	if params.Limit > 0 {
		limit = int(params.Limit)
	}

	totalPages := uint(0)
	if totalCount > 0 && limit > 0 {
		totalPages = uint(math.Ceil(float64(totalCount) / float64(limit)))
	}

	currentPage := uint(1)
	if params.Page > 0 {
		currentPage = params.Page
	}

	usersResponse := make([]response.GetUserResponseSimple, 0, len(*users))
	for i := range *users {
		user := (*users)[i]
		profilePicUrl := user.ProfilePic
		if profilePicUrl != nil {
			profilePicUrl = utils.MountCloudFrontUrl(*user.ProfilePic)
		}
		usersResponse = append(usersResponse, response.GetUserResponseSimple{
			ID:            user.ID,
			Name:          user.Name,
			Nickname:      user.Nickname,
			ProfilePicUrl: profilePicUrl,
			Since:         &user.CreatedAt,
		})
	}

	return &response.GetUsersResponse{
		Users:       usersResponse,
		CurrentPage: currentPage,
		TotalPages:  totalPages,
	}, nil
}
