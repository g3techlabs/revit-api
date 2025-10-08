package response

type ProfilePicPresignedURL struct {
	PresignedURL string `json:"presignedUrl"`
	ObjectKey    string `json:"objectKey"`
}
