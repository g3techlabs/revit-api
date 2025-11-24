package response

// PresignedPhotoInfo representa a URL pré-assinada para upload de foto
// @Description URL pré-assinada e chave para upload de foto do veículo
type PresignedPhotoInfo struct {
	// URL pré-assinada para upload da foto do veículo
	PresignedVehiclePhotoUrl *string `json:"presignedVehiclePhotoUrl" example:"https://s3.amazonaws.com/bucket/vehicles/123/photo.jpg?X-Amz-Signature=..."`
	// Chave da foto no storage
	ObjectKey *string `json:"objectKey" example:"vehicles/123/photo.jpg"`
	// ID do veículo
	VehicleId *uint `json:"vehicleId" example:"123"`
}
