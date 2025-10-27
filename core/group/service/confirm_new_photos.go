package service

import (
	"github.com/g3techlabs/revit-api/core/group/errors"
	"github.com/g3techlabs/revit-api/core/group/input"
	"github.com/g3techlabs/revit-api/response/generics"
)

func (gs *GroupService) ConfirmNewPhotos(userId, groupId uint, data *input.ConfirmNewPhotos) error {
	if err := gs.validator.Validate(data); err != nil {
		return err
	}

	if data.BannerKey != nil {
		if err := gs.storageService.DoesObjectExist(*data.BannerKey); err != nil {
			return err
		}

		if err := gs.groupRepo.UpdateBanner(userId, groupId, *data.BannerKey); err != nil {
			if err.Error() == "group not found or user not allowed" {
				return errors.RequesterIsNotAnAdmin()
			}
			return generics.InternalError()
		}

	}
	if data.MainPhotoKey != nil {
		if err := gs.storageService.DoesObjectExist(*data.MainPhotoKey); err != nil {
			return err
		}

		if err := gs.groupRepo.UpdateMainPhoto(userId, groupId, *data.MainPhotoKey); err != nil {
			if err.Error() == "group not found or user not allowed" {
				return errors.RequesterIsNotAnAdmin()
			}
			return generics.InternalError()
		}
	}

	return nil
}
