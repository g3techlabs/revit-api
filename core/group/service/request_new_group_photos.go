package service

import (
	"github.com/g3techlabs/revit-api/core/group/input"
	"github.com/g3techlabs/revit-api/core/group/response"
	"github.com/g3techlabs/revit-api/response/generics"
)

func (gs *GroupService) RequestNewGroupPhotos(userId, groupId uint, data *input.RequestNewGroupPhotos) (*response.PresignedGroupPhotosInfo, error) {
	if err := gs.validator.Validate(data); err != nil {
		return nil, err
	}

	isUserAdmin, err := gs.groupRepo.IsUserAdmin(userId, groupId)
	if err != nil {
		return nil, generics.InternalError()
	}

	if !isUserAdmin {
		return nil, generics.Forbidden("User does not have permission of editing this group")
	}

	response, err := gs.buildResponse(groupId, data.MainPhotoContentType, data.BannerContentType)
	if err != nil {
		return nil, err
	}

	return response, nil
}
