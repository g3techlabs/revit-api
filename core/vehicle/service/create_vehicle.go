package service

import (
	"fmt"
	"strings"

	"github.com/g3techlabs/revit-api/core/vehicle/errors"
	"github.com/g3techlabs/revit-api/core/vehicle/input"
	"github.com/g3techlabs/revit-api/core/vehicle/response"
	"github.com/g3techlabs/revit-api/response/generics"
)

func (vs *VehicleService) CreateVehicle(userId uint, data *input.CreateVehicle) (*response.VehicleCreated, error) {
	if err := vs.validator.Validate(data); err != nil {
		return nil, err
	}

	vs.lowerStrings(data)
	vehicleModel := data.ToVehicleModel(userId)
	if err := vs.vehicleRepo.CreateVehicle(vehicleModel); err != nil {
		return nil, generics.InternalError()
	}

	response := new(response.VehicleCreated)
	if data.MainPhotoContentTye == nil {
		return response, nil
	}

	response, err := vs.buildResponse(userId, vehicleModel.ID, data.MainPhotoContentTye)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (vs *VehicleService) lowerStrings(i *input.CreateVehicle) {
	i.Nickname = strings.ToLower(i.Nickname)
	i.Brand = strings.ToLower(i.Brand)
	i.Model = strings.ToLower(i.Model)
	if i.Version != nil {
		lowered := strings.ToLower(*i.Version)
		i.Version = &lowered
	}
}

func (vs *VehicleService) buildResponse(userId, vehicleId uint, contentType *string) (*response.VehicleCreated, error) {
	response := new(response.VehicleCreated)
	if contentType == nil {
		return response, nil
	}

	objectKey, err := vs.generatePresignedMainPhotoUrl(userId, vehicleId, *contentType, response)
	if err != nil {
		return nil, err
	}

	response.ObjectKey = &objectKey
	response.VehicleId = &vehicleId
	return response, nil
}

func (vs *VehicleService) generatePresignedMainPhotoUrl(userId, vehicleId uint, contentType string, r *response.VehicleCreated) (string, error) {
	extension := vs.mapContentTypeToExtension(contentType)
	if extension == "" {
		return "", errors.InvalidFileExtension()
	}
	mainPhotoKey := fmt.Sprintf("users/%d/vehicles/%d/main%s", userId, vehicleId, extension)
	presignedUrl, err := vs.storageService.PresignPutObjectURL(mainPhotoKey, contentType)
	if err != nil {
		return "", generics.InternalError()
	}

	r.PresignedVehiclePhotoUrl = &presignedUrl

	return mainPhotoKey, nil
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
