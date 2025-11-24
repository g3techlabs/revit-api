package controller

import (
	"strconv"

	"github.com/g3techlabs/revit-api/src/core/event/input"
	_ "github.com/g3techlabs/revit-api/src/core/event/response"
	"github.com/g3techlabs/revit-api/src/core/event/service"
	"github.com/g3techlabs/revit-api/src/response/generics"
	"github.com/gofiber/fiber/v2"
)

type EventController struct {
	eventService service.IEventService
}

func NewEventController(eventService service.IEventService) *EventController {
	return &EventController{
		eventService: eventService,
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

// ConfirmNewPhotoRequest representa a confirmação de upload de foto (apenas para Swagger)
type ConfirmNewPhotoRequest struct {
	Key string `json:"key" validate:"required"`
}

// CreateEvent godoc
// @Summary Criar evento
// @Description Cria um novo evento
// @Tags Events
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param event body input.CreateEventInput true "Dados do evento"
// @Success 200 {object} response.PresginedEventPhotoResponse "Evento criado com sucesso"
// @Failure 400 {object} ValidationErrorResponse "Erro na validação dos dados"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Router /api/event [post]
func (c *EventController) CreateEvent(ctx *fiber.Ctx) error {
	var data input.CreateEventInput
	if err := ctx.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body request")
	}

	userId := ctx.Locals("userId").(uint)

	response, err := c.eventService.CreateEvent(userId, &data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

// ConfirmNewPhoto godoc
// @Summary Confirmar upload de foto do evento
// @Description Confirma o upload de foto após upload na URL pré-assinada
// @Tags Events
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param eventId path int true "ID do evento"
// @Param confirmation body ConfirmNewPhotoRequest true "Confirmação do upload"
// @Success 204 "Foto confirmada com sucesso"
// @Failure 400 {object} ValidationErrorResponse "Erro na validação dos dados"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Failure 403 {object} ErrorMessageResponse "Sem permissão para atualizar o evento"
// @Failure 404 {object} ErrorMessageResponse "Evento não encontrado"
// @Router /api/event/photo/{eventId} [patch]
func (c *EventController) ConfirmNewPhoto(ctx *fiber.Ctx) error {
	var data input.ConfirmNewPhoto
	if err := ctx.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body request")
	}

	eventParam := ctx.Params("eventId")
	eventIdUint64, err := strconv.ParseUint(eventParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid event ID",
		})
	}
	eventId := uint(eventIdUint64)

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	if err := c.eventService.ConfirmNewPhoto(userId, eventId, &data); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

// GetEvents godoc
// @Summary Listar eventos
// @Description Retorna a lista de eventos do usuário autenticado
// @Tags Events
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param filters query input.GetEventsFilters false "Filtros de busca"
// @Success 200 {array} response.GetEventResponse "Lista de eventos"
// @Failure 400 {object} ValidationErrorResponse "Erro na validação dos parâmetros"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Router /api/event [get]
func (c *EventController) GetEvents(ctx *fiber.Ctx) error {
	var data input.GetEventsFilters
	if err := ctx.QueryParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid query parameters")
	}

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	response, err := c.eventService.GetEvents(userId, &data)
	if err != nil {
		return err
	}

	return ctx.JSON(response)
}

// UpdateEvent godoc
// @Summary Atualizar evento
// @Description Atualiza os dados de um evento
// @Tags Events
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param eventId path int true "ID do evento"
// @Param event body input.UpdateEventInput true "Dados do evento para atualização"
// @Success 204 "Evento atualizado com sucesso"
// @Failure 400 {object} ValidationErrorResponse "Erro na validação dos dados"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Failure 403 {object} ErrorMessageResponse "Sem permissão para atualizar o evento"
// @Failure 404 {object} ErrorMessageResponse "Evento não encontrado"
// @Router /api/event/{eventId} [patch]
func (c *EventController) UpdateEvent(ctx *fiber.Ctx) error {
	var data input.UpdateEventInput
	if err := ctx.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid query parameters")
	}

	eventParam := ctx.Params("eventId")
	eventIdUint64, err := strconv.ParseUint(eventParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid event ID",
		})
	}
	eventId := uint(eventIdUint64)

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	if err := c.eventService.UpdateEvent(userId, eventId, &data); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

// RequestNewPhoto godoc
// @Summary Solicitar upload de foto do evento
// @Description Solicita uma URL pré-assinada para upload de foto do evento
// @Tags Events
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param eventId path int true "ID do evento"
// @Param request body input.RequestNewPhotoInput true "Dados da requisição de foto"
// @Success 201 {object} response.PresginedEventPhotoResponse "URL pré-assinada para upload"
// @Failure 400 {object} ValidationErrorResponse "Erro na validação dos dados"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Failure 403 {object} ErrorMessageResponse "Sem permissão para atualizar o evento"
// @Failure 404 {object} ErrorMessageResponse "Evento não encontrado"
// @Router /api/event/photo/{eventId} [post]
func (c *EventController) RequestNewPhoto(ctx *fiber.Ctx) error {
	var data input.RequestNewPhotoInput
	if err := ctx.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid query parameters")
	}

	eventParam := ctx.Params("eventId")
	eventIdUint64, err := strconv.ParseUint(eventParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid event ID",
		})
	}
	eventId := uint(eventIdUint64)

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	response, err := c.eventService.RequestNewPhoto(userId, eventId, &data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

// SubscribeIntoEvent godoc
// @Summary Inscrever-se em evento
// @Description Adiciona o usuário autenticado como participante de um evento
// @Tags Events
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param eventId path int true "ID do evento"
// @Success 204 "Usuário inscrito no evento"
// @Failure 400 {object} ErrorMessageResponse "ID inválido"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Failure 404 {object} ErrorMessageResponse "Evento não encontrado"
// @Router /api/event/{eventId}/subscriber [post]
func (c *EventController) SubscribeIntoEvent(ctx *fiber.Ctx) error {
	eventParam := ctx.Params("eventId")
	eventIdUint64, err := strconv.ParseUint(eventParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid event ID",
		})
	}
	eventId := uint(eventIdUint64)

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	if err := c.eventService.SubscribeToEvent(userId, eventId); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

// RevokeEventSubscription godoc
// @Summary Cancelar inscrição em evento
// @Description Remove o usuário autenticado de um evento
// @Tags Events
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param eventId path int true "ID do evento"
// @Success 204 "Inscrição cancelada com sucesso"
// @Failure 400 {object} ErrorMessageResponse "ID inválido"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Failure 404 {object} ErrorMessageResponse "Evento ou inscrição não encontrada"
// @Router /api/event/{eventId}/subscriber [delete]
func (c *EventController) RevokeEventSubscription(ctx *fiber.Ctx) error {
	eventParam := ctx.Params("eventId")
	eventIdUint64, err := strconv.ParseUint(eventParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid event ID",
		})
	}
	eventId := uint(eventIdUint64)

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	if err := c.eventService.RevokeEventSubscription(userId, eventId); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

// InviteUserToEvent godoc
// @Summary Convidar usuário para evento
// @Description Envia um convite para um usuário participar de um evento
// @Tags Events
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param eventId path int true "ID do evento"
// @Param invitedId path int true "ID do usuário a ser convidado"
// @Success 204 "Convite enviado com sucesso"
// @Failure 400 {object} ErrorMessageResponse "ID inválido"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Failure 403 {object} ErrorMessageResponse "Sem permissão para convidar"
// @Failure 404 {object} ErrorMessageResponse "Evento ou usuário não encontrado"
// @Router /api/event/{eventId}/invite/{invitedId} [post]
func (c *EventController) InviteUserToEvent(ctx *fiber.Ctx) error {
	eventParam := ctx.Params("eventId")
	eventIdUint64, err := strconv.ParseUint(eventParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid event ID",
		})
	}
	eventId := uint(eventIdUint64)

	invitedParam := ctx.Params("invitedId")
	invitedIdUint64, err := strconv.ParseUint(invitedParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid invite target ID",
		})
	}
	invitedId := uint(invitedIdUint64)

	eventAdminId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	if err := c.eventService.InviteUserToEvent(eventAdminId, eventId, invitedId); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

// GetPendingInvites godoc
// @Summary Listar convites pendentes
// @Description Retorna os convites de eventos pendentes do usuário autenticado
// @Tags Events
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param filters query input.GetPendingInvitesFilters false "Filtros de busca"
// @Success 200 {array} response.GetPendingInvitesResponse "Lista de convites pendentes"
// @Failure 400 {object} ValidationErrorResponse "Erro na validação dos parâmetros"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Router /api/event/invite [get]
func (c *EventController) GetPendingInvites(ctx *fiber.Ctx) error {
	var filters input.GetPendingInvitesFilters

	if err := ctx.QueryParser(&filters); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid query parameters")
	}

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	response, err := c.eventService.GetPendingInvites(userId, &filters)
	if err != nil {
		return err
	}

	return ctx.JSON(response)
}

// AnswerPendingInvite godoc
// @Summary Responder convite pendente
// @Description Aceita ou rejeita um convite de evento pendente
// @Tags Events
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param eventId path int true "ID do evento"
// @Param answer body input.PendingInviteAnswer true "Resposta ao convite (accept/reject)"
// @Success 204 "Resposta processada com sucesso"
// @Failure 400 {object} ValidationErrorResponse "Erro na validação dos dados"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Failure 404 {object} ErrorMessageResponse "Convite não encontrado"
// @Router /api/event/{eventId}/invite [patch]
func (c *EventController) AnswerPendingInvite(ctx *fiber.Ctx) error {
	var answer input.PendingInviteAnswer

	if err := ctx.BodyParser(&answer); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body request")
	}

	eventParam := ctx.Params("eventId")
	eventIdUint64, err := strconv.ParseUint(eventParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid event ID",
		})
	}
	eventId := uint(eventIdUint64)

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	if err := c.eventService.AnswerPendingInvite(userId, eventId, &answer); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

// RemoveSubscriber godoc
// @Summary Remover participante do evento
// @Description Remove um participante de um evento (apenas administradores)
// @Tags Events
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param eventId path int true "ID do evento"
// @Param subscriberId path int true "ID do participante a ser removido"
// @Success 204 "Participante removido com sucesso"
// @Failure 400 {object} ErrorMessageResponse "ID inválido"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Failure 403 {object} ErrorMessageResponse "Sem permissão para remover participante"
// @Failure 404 {object} ErrorMessageResponse "Evento ou participante não encontrado"
// @Router /api/event/{eventId}/subscriber/{subscriberId} [delete]
func (c *EventController) RemoveSubscriber(ctx *fiber.Ctx) error {
	eventParam := ctx.Params("eventId")
	eventIdUint64, err := strconv.ParseUint(eventParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid event ID",
		})
	}
	eventId := uint(eventIdUint64)

	subscriberParam := ctx.Params("subscriberId")
	subscriberIdUint64, err := strconv.ParseUint(subscriberParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid subscriber ID",
		})
	}
	subscriberId := uint(subscriberIdUint64)

	eventAdminId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	if err := c.eventService.RemoveSubscriber(eventAdminId, eventId, subscriberId); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
