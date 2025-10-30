package service

import (
	"fmt"

	"github.com/g3techlabs/revit-api/core/group/input"
	"github.com/g3techlabs/revit-api/core/group/response"
	"github.com/g3techlabs/revit-api/response/generics"
	"github.com/g3techlabs/revit-api/utils"
)

func (gs *GroupService) CreateGroup(userId uint, data *input.CreateGroup) (*response.PresignedGroupPhotosInfo, error) {
	if err := gs.validator.Validate(data); err != nil {
		return nil, err
	}

	modelData := data.ToGroupModel()
	if err := gs.groupRepo.CreateGroup(userId, modelData); err != nil {
		return nil, generics.InternalError()
	}

	response, err := gs.buildResponse(modelData.ID, data.MainPhotoContentType, data.BannerContentType)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (gs *GroupService) buildResponse(groupId uint, mainPhotoContentType, bannerContentType *string) (*response.PresignedGroupPhotosInfo, error) {
	response := new(response.PresignedGroupPhotosInfo)

	if mainPhotoContentType == nil && bannerContentType == nil {
		return response, nil
	}

	response.GroupId = &groupId

	if mainPhotoContentType != nil {
		if err := gs.makePresignedMainPhotoUrl(groupId, *mainPhotoContentType, response); err != nil {
			return nil, err
		}
	}

	if bannerContentType != nil {
		if err := gs.makePresignedBannerUrl(groupId, *bannerContentType, response); err != nil {
			return nil, err
		}
	}

	return response, nil
}

func (gs *GroupService) makePresignedMainPhotoUrl(groupId uint, contentType string, r *response.PresignedGroupPhotosInfo) error {
	const MAIN_PHOTO_KEY = "groups/%d/main%s"

	extension := utils.MapImageMIMEToExtension(contentType)
	if extension == "" {
		return generics.InvalidFileExtension()
	}

	photoKey := fmt.Sprintf(MAIN_PHOTO_KEY, groupId, extension)
	presignedUrl, err := gs.storageService.PresignPutObjectURL(photoKey, contentType)
	if err != nil {
		return generics.InternalError()
	}

	r.PresignedMainPhotoUrl = &presignedUrl
	r.MainPhotoKey = &photoKey

	return nil
}

func (gs *GroupService) makePresignedBannerUrl(groupId uint, contentType string, r *response.PresignedGroupPhotosInfo) error {
	const BANNER_KEY = "groups/%d/banner%s"

	extension := utils.MapImageMIMEToExtension(contentType)
	if extension == "" {
		return generics.InvalidFileExtension()
	}

	bannerKey := fmt.Sprintf(BANNER_KEY, groupId, extension)
	presignedBannerUrl, err := gs.storageService.PresignPutObjectURL(bannerKey, contentType)
	if err != nil {
		return generics.InternalError()
	}

	r.PresignedBannerUrl = &presignedBannerUrl
	r.BannerKey = &bannerKey

	return nil
}
