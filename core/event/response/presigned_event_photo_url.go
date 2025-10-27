package response

type PresginedEventPhotoResponse struct {
	EventId   *uint   `json:"eventId"`
	Url       *string `json:"url"`
	ObjectKey *string `json:"key"`
}
