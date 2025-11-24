package controller

import (
	"strconv"

	geoinput "github.com/g3techlabs/revit-api/src/core/geolocation/geo_input"
	"github.com/g3techlabs/revit-api/src/core/route/input"
	_ "github.com/g3techlabs/revit-api/src/core/route/response"
	"github.com/g3techlabs/revit-api/src/core/route/service"
	"github.com/gofiber/fiber/v2"
)

type RouteController struct {
	routeService service.IRouteService
}

func NewRouteController(routeService service.IRouteService) *RouteController {
	return &RouteController{
		routeService: routeService,
	}
}

// ValidationErrorResponse representa a resposta de erro de validação
type ValidationErrorResponse struct {
	Errors map[string]string `json:"errors"`
}

// ErrorMessageResponse representa uma resposta de erro simples
type ErrorMessageResponse struct {
	Message string `json:"message"`
}

// CoordinatesRequest representa coordenadas geográficas (apenas para Swagger)
type CoordinatesRequest struct {
	Lat  float64 `json:"lat" validate:"required,number,gte=-85.05112878,lte=85.05112878"`
	Long float64 `json:"long" validate:"required,longitude"`
}

// CreateRoute godoc
// @Summary Criar rota
// @Description Cria uma nova rota com localização de início e destino
// @Tags Routes
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param route body input.CreateRouteInput true "Dados da rota (localização inicial e destino)"
// @Success 201 {object} response.RouteCreatedReponse "Rota criada com sucesso"
// @Failure 400 {object} ValidationErrorResponse "Erro na validação dos dados. Possíveis campos: startLocation, destination"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Router /api/route [post]
func (c *RouteController) CreateRoute(ctx *fiber.Ctx) error {
	var data input.CreateRouteInput

	if err := ctx.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body request")
	}

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid user ID")
	}

	response, err := c.routeService.CreateRoute(userId, &data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

// InviteUsers godoc
// @Summary Convidar usuários para rota
// @Description Envia convites para uma lista de usuários participarem de uma rota
// @Tags Routes
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param routeId path int true "ID da rota"
// @Param invite body input.UsersToInviteInput true "Lista de IDs dos usuários a serem convidados"
// @Success 204 "Convites enviados com sucesso"
// @Failure 400 {object} ValidationErrorResponse "Erro na validação dos dados. Possíveis campos: idsToInvite"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Failure 403 {object} ErrorMessageResponse "Sem permissão para convidar usuários para esta rota"
// @Failure 404 {object} ErrorMessageResponse "Rota não encontrada"
// @Router /api/route/{routeId}/invite [put]
func (c *RouteController) InviteUsers(ctx *fiber.Ctx) error {
	var data input.UsersToInviteInput

	if err := ctx.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body request")
	}

	routeParam := ctx.Params("routeId")
	routeIdUint64, err := strconv.ParseUint(routeParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid route ID",
		})
	}
	routeId := uint(routeIdUint64)

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid user ID")
	}

	if err := c.routeService.InviteUsers(userId, routeId, &data); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

// GetOnlineFriendsToInvite godoc
// @Summary Listar amigos online para convite
// @Description Retorna a lista de amigos online que podem ser convidados para uma rota
// @Tags Routes
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param limit query int true "Limite de resultados por página (1-99)"
// @Param page query int true "Número da página (maior que 0)"
// @Success 200 {array} response.OnlineFriendsResponse "Lista de amigos online"
// @Failure 400 {object} ValidationErrorResponse "Erro na validação dos parâmetros. Possíveis campos: limit, page"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Router /api/route/friends [get]
func (c *RouteController) GetOnlineFriendsToInvite(ctx *fiber.Ctx) error {
	var query input.GetOnlineFriendsToInviteQuery

	if err := ctx.QueryParser(&query); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body request")
	}

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid user ID")
	}

	response, err := c.routeService.GetOnlineFriendsToInvite(userId, &query)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

// GetNearbyUsersToRouteInvite godoc
// @Summary Listar usuários próximos para convite
// @Description Retorna a lista de usuários próximos a uma localização que podem ser convidados para uma rota
// @Tags Routes
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param lat query number true "Latitude da localização (-85.05112878 a 85.05112878)"
// @Param long query number true "Longitude da localização"
// @Param limit query int true "Limite de resultados por página (1-99)"
// @Param page query int true "Número da página (maior que 0)"
// @Success 200 {array} response.NearbyUserToRouteResponse "Lista de usuários próximos"
// @Failure 400 {object} ValidationErrorResponse "Erro na validação dos parâmetros. Possíveis campos: lat, long, limit, page"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Router /api/route/nearby [get]
func (c *RouteController) GetNearbyUsersToRouteInvite(ctx *fiber.Ctx) error {
	var data input.GetNearbyUsersToInviteQuery

	if err := ctx.QueryParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body request")
	}

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid user ID")
	}

	response, err := c.routeService.GetNearbyUsersToInvite(userId, &data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

// AcceptRouteInvite godoc
// @Summary Aceitar convite de rota
// @Description Aceita um convite para participar de uma rota, fornecendo as coordenadas iniciais
// @Tags Routes
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param routeId path int true "ID da rota"
// @Param coordinates body CoordinatesRequest true "Coordenadas geográficas iniciais do usuário"
// @Success 204 "Convite aceito com sucesso"
// @Failure 400 {object} ValidationErrorResponse "Erro na validação dos dados. Possíveis campos: lat, long"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Failure 404 {object} ErrorMessageResponse "Rota ou convite não encontrado"
// @Router /api/route/{routeId}/invite [patch]
func (c *RouteController) AcceptRouteInvite(ctx *fiber.Ctx) error {
	var coordinates geoinput.Coordinates

	if err := ctx.BodyParser(&coordinates); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body request")
	}

	routeParam := ctx.Params("routeId")
	routeIdUint64, err := strconv.ParseUint(routeParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid route ID",
		})
	}
	routeId := uint(routeIdUint64)

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid user ID")
	}

	if err := c.routeService.AcceptRouteInvite(userId, routeId, &coordinates); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
