package response

type StartRouteEvent struct {
	Event   string            `json:"event"`
	Payload StartRoutePayload `json:"payload"`
}

type StartRoutePayload struct {
	RouteID            uint  `json:"routeId"`
	DisplayAtTimestamp int64 `json:"displayAtTimestamp"`
}

func NewStartRouteEvent(routeId uint, displayAtTimestamp int64) StartRouteEvent {
	return StartRouteEvent{
		Event: "start-route",
		Payload: StartRoutePayload{
			RouteID:            routeId,
			DisplayAtTimestamp: displayAtTimestamp,
		},
	}
}
