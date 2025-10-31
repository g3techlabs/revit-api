package response

type PresignedPhotoInfo struct {
	PresignedVehiclePhotoUrl *string `json:"presignedVehiclePhotoUrl"`
	ObjectKey                *string `json:"objectKey"`
	VehicleId                *uint   `json:"vehicleId"`
}
