package geoinput

type Coordinates struct {
	Lat  float64 `json:"lat" validate:"required,number,gte=-85.05112878,lte=85.05112878"`
	Long float64 `json:"long" validate:"required,longitude"`
}
