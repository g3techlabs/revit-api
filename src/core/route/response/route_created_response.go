package response

// RouteCreatedReponse representa a resposta de criação de rota
// @Description Resposta retornada quando uma rota é criada com sucesso
type RouteCreatedReponse struct {
	// ID da rota criada
	RouteID uint `json:"routeId" example:"123"`
}
