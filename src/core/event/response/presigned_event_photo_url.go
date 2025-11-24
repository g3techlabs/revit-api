package response

// PresginedEventPhotoResponse representa a URL pré-assinada para upload de foto
// @Description URL pré-assinada e chave para upload de foto do evento
type PresginedEventPhotoResponse struct {
	// ID do evento
	EventId *uint `json:"eventId" example:"456"`
	// URL pré-assinada para upload da foto
	Url *string `json:"url" example:"https://s3.amazonaws.com/bucket/events/456/photo.jpg?X-Amz-Signature=..."`
	// Chave da foto no storage
	ObjectKey *string `json:"key" example:"events/456/photo.jpg"`
}
