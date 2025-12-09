package input

type GetNearbyUsersToInviteQuery struct {
	Lat   float64 `json:"lat" validate:"required,number,gte=-85.05112878,lte=85.05112878"`
	Long  float64 `json:"long" validate:"required,longitude"`
	Limit uint    `json:"limit" validate:"required,number,gt=0,lt=100"`
	Page  uint    `json:"page" validate:"required,number,gt=0"`
}
