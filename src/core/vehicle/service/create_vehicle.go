package service

import (
	"fmt"

	"github.com/g3techlabs/revit-api/src/core/vehicle/input"
	"github.com/g3techlabs/revit-api/src/core/vehicle/response"
	"github.com/g3techlabs/revit-api/src/response/generics"
	"github.com/google/uuid"
)

func (vs *VehicleService) CreateVehicle(userId uint, data *input.CreateVehicle) (*response.PresignedPhotoInfo, error) {
	if err := vs.validator.Validate(data); err != nil {
		return nil, err
	}

	vehicleModel := data.ToVehicleModel(userId)
	if err := vs.vehicleRepo.CreateVehicle(vehicleModel); err != nil {
		return nil, generics.InternalError()
	}

	response, err := vs.buildPresignedPhotoResponse(userId, vehicleModel.ID, data.MainPhotoContentTye, "main")
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (vs *VehicleService) buildPresignedPhotoResponse(userId, vehicleId uint, contentType *string, photoType string) (*response.PresignedPhotoInfo, error) {
	response := new(response.PresignedPhotoInfo)
	if contentType == nil {
		return response, nil
	}

	if err := vs.generatePresignedPhotoUrl(userId, vehicleId, *contentType, photoType, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (vs *VehicleService) generatePresignedPhotoUrl(userId, vehicleId uint, contentType, photoType string, r *response.PresignedPhotoInfo) error {
	const MAIN_PHOTO_KEY = "users/%d/vehicles/%d/main%s"
	const FEED_PHOTO_KEY = "users/%d/vehicles/%d/photos/%s%s"

	extension := vs.mapContentTypeToExtension(contentType)
	if extension == "" {
		return generics.InvalidFileExtension()
	}

	switch photoType {
	case "main":
		photoKey := fmt.Sprintf(MAIN_PHOTO_KEY, userId, vehicleId, extension)
		presignedUrl, err := vs.storageService.PresignPutObjectURL(photoKey, contentType)
		if err != nil {
			return generics.InternalError()
		}

		r.PresignedVehiclePhotoUrl = &presignedUrl
		r.ObjectKey = &photoKey
	case "feed":
		photoUUID := uuid.New().String()
		photoKey := fmt.Sprintf(FEED_PHOTO_KEY, userId, vehicleId, photoUUID, extension)
		presignedUrl, err := vs.storageService.PresignPutObjectURL(photoKey, contentType)
		if err != nil {
			return generics.InternalError()
		}

		r.PresignedVehiclePhotoUrl = &presignedUrl
		r.ObjectKey = &photoKey
	}

	r.VehicleId = &vehicleId

	return nil
}

func (s *VehicleService) mapContentTypeToExtension(contentType string) string {
	switch contentType {
	case "image/jpeg":
		return ".jpg"
	case "image/png":
		return ".png"
	case "image/webp":
		return ".webp"
	default:
		return ""
	}
}
