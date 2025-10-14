package response

type PresignedMainPhotoInfo struct {
	PresignedVehiclePhotoUrl *string `json:"presignedVehiclePhotoUrl"`
	ObjectKey                *string `json:"objectKey"`
	VehicleId                *uint   `json:"vehicleId"`
}
