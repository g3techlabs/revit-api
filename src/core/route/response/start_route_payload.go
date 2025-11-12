package response

type StartRoutePayload struct {
	RouteID            uint  `json:"routeId"`
	DisplayAtTimestamp int64 `json:"displayAtTimestamp"`
}

func NewStartRoutePayload(routeId uint, displayAtTimestamp int64) StartRoutePayload {
	return StartRoutePayload{
		RouteID:            routeId,
		DisplayAtTimestamp: displayAtTimestamp,
	}
}
