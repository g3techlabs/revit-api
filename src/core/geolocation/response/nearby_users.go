package response

type NearbyUsers struct {
	UserID uint    `json:"userId"`
	Lat    float64 `json:"lat"`
	Lng    float64 `json:"lng"`
}
