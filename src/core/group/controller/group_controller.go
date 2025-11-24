package controller

import (
	"strconv"

	"github.com/g3techlabs/revit-api/src/core/group/input"
	_ "github.com/g3techlabs/revit-api/src/core/group/response"
	"github.com/g3techlabs/revit-api/src/core/group/service"
	"github.com/g3techlabs/revit-api/src/response/generics"
	"github.com/gofiber/fiber/v2"
)

type GroupController struct {
	groupService service.IGroupService
}

func NewGroupController(groupService service.IGroupService) *GroupController {
	return &GroupController{groupService: groupService}
}

// ValidationErrorResponse representa a resposta de erro de validação
type ValidationErrorResponse struct {
	Errors map[string]string `json:"errors"`
}

// ErrorMessageResponse representa uma resposta de erro simples
type ErrorMessageResponse struct {
	Message string `json:"message"`
}

// CreateGroup godoc
// @Summary Criar grupo
// @Description Cria um novo grupo
// @Tags Groups
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param group body input.CreateGroup true "Dados do grupo"
// @Success 201 {object} response.PresignedGroupPhotosInfo "Grupo criado com sucesso"
// @Failure 400 {object} ValidationErrorResponse "Erro na validação dos dados"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Router /api/group [post]
func (c *GroupController) CreateGroup(ctx *fiber.Ctx) error {
	data := new(input.CreateGroup)

	if err := ctx.BodyParser(data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body request")
	}

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	response, err := c.groupService.CreateGroup(userId, data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

// ConfirmNewPhotos godoc
// @Summary Confirmar upload de fotos do grupo
// @Description Confirma o upload de fotos após upload nas URLs pré-assinadas
// @Tags Groups
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param groupId path int true "ID do grupo"
// @Param confirmation body input.ConfirmNewPhotos true "Confirmação do upload"
// @Success 204 "Fotos confirmadas com sucesso"
// @Failure 400 {object} ValidationErrorResponse "Erro na validação dos dados"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Failure 403 {object} ErrorMessageResponse "Sem permissão para atualizar fotos"
// @Failure 404 {object} ErrorMessageResponse "Grupo não encontrado"
// @Router /api/photos/{groupId} [patch]
func (c *GroupController) ConfirmNewPhotos(ctx *fiber.Ctx) error {
	data := new(input.ConfirmNewPhotos)

	if err := ctx.BodyParser(data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body request")
	}

	groupParam := ctx.Params("groupId")

	groupIdUint64, err := strconv.ParseUint(groupParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid group ID",
		})
	}
	groupId := uint(groupIdUint64)

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	if err := c.groupService.ConfirmNewPhotos(userId, groupId, data); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

// GetGroups godoc
// @Summary Listar grupos
// @Description Retorna a lista de grupos do usuário autenticado
// @Tags Groups
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param query query input.GetGroupsQuery false "Parâmetros de filtro"
// @Success 200 {array} map[string]interface{} "Lista de grupos"
// @Failure 400 {object} ValidationErrorResponse "Erro na validação dos parâmetros"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Router /api/group [get]
func (c *GroupController) GetGroups(ctx *fiber.Ctx) error {
	query := new(input.GetGroupsQuery)

	if err := ctx.QueryParser(query); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid query parameters")
	}

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	response, err := c.groupService.GetGroups(userId, query)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

// UpdateGroup godoc
// @Summary Atualizar grupo
// @Description Atualiza os dados de um grupo
// @Tags Groups
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param groupId path int true "ID do grupo"
// @Param group body input.UpdateGroup true "Dados do grupo para atualização"
// @Success 204 "Grupo atualizado com sucesso"
// @Failure 400 {object} ValidationErrorResponse "Erro na validação dos dados"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Failure 403 {object} ErrorMessageResponse "Sem permissão para atualizar o grupo"
// @Failure 404 {object} ErrorMessageResponse "Grupo não encontrado"
// @Router /api/group/{groupId} [patch]
func (c *GroupController) UpdateGroup(ctx *fiber.Ctx) error {
	data := new(input.UpdateGroup)

	if err := ctx.BodyParser(data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body request")
	}

	groupParam := ctx.Params("groupId")
	groupIdUint64, err := strconv.ParseUint(groupParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid group ID",
		})
	}
	groupId := uint(groupIdUint64)

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	if err := c.groupService.UpdateGroup(userId, groupId, data); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

// RequestNewGroupPhotos godoc
// @Summary Solicitar upload de fotos do grupo
// @Description Solicita URLs pré-assinadas para upload de fotos do grupo
// @Tags Groups
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param groupId path int true "ID do grupo"
// @Param request body input.RequestNewGroupPhotos true "Dados da requisição de fotos"
// @Success 201 {object} response.PresignedGroupPhotosInfo "URLs pré-assinadas para upload"
// @Failure 400 {object} ValidationErrorResponse "Erro na validação dos dados"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Failure 403 {object} ErrorMessageResponse "Sem permissão para atualizar fotos"
// @Failure 404 {object} ErrorMessageResponse "Grupo não encontrado"
// @Router /api/photos/{groupId} [put]
func (c *GroupController) RequestNewGroupPhotos(ctx *fiber.Ctx) error {
	data := new(input.RequestNewGroupPhotos)

	if err := ctx.BodyParser(data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body request")
	}

	groupParam := ctx.Params("groupId")
	groupIdUint64, err := strconv.ParseUint(groupParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid group ID",
		})
	}
	groupId := uint(groupIdUint64)

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	response, err := c.groupService.RequestNewGroupPhotos(userId, groupId, data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

// JoinGroup godoc
// @Summary Entrar em grupo
// @Description Adiciona o usuário autenticado a um grupo público
// @Tags Groups
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param groupId path int true "ID do grupo"
// @Success 204 "Usuário adicionado ao grupo"
// @Failure 400 {object} ErrorMessageResponse "ID inválido"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Failure 403 {object} ErrorMessageResponse "Grupo privado ou sem permissão"
// @Failure 404 {object} ErrorMessageResponse "Grupo não encontrado"
// @Router /api/group/{groupId}/member [post]
func (c *GroupController) JoinGroup(ctx *fiber.Ctx) error {
	groupParam := ctx.Params("groupId")
	groupIdUint64, err := strconv.ParseUint(groupParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid group ID",
		})
	}
	groupId := uint(groupIdUint64)

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	if err := c.groupService.JoinGroup(userId, groupId); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

// QuitGroup godoc
// @Summary Sair do grupo
// @Description Remove o usuário autenticado de um grupo
// @Tags Groups
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param groupId path int true "ID do grupo"
// @Success 204 "Usuário removido do grupo"
// @Failure 400 {object} ErrorMessageResponse "ID inválido"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Failure 404 {object} ErrorMessageResponse "Grupo ou membro não encontrado"
// @Router /api/group/{groupId}/member [delete]
func (c *GroupController) QuitGroup(ctx *fiber.Ctx) error {
	groupParam := ctx.Params("groupId")
	groupIdUint64, err := strconv.ParseUint(groupParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid group ID",
		})
	}
	groupId := uint(groupIdUint64)

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	if err := c.groupService.QuitGroup(userId, groupId); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

// InviteUser godoc
// @Summary Convidar usuário para grupo
// @Description Envia um convite para um usuário entrar no grupo
// @Tags Groups
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param groupId path int true "ID do grupo"
// @Param invitedId path int true "ID do usuário a ser convidado"
// @Success 204 "Convite enviado com sucesso"
// @Failure 400 {object} ErrorMessageResponse "ID inválido"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Failure 403 {object} ErrorMessageResponse "Sem permissão para convidar"
// @Failure 404 {object} ErrorMessageResponse "Grupo ou usuário não encontrado"
// @Router /api/group/{groupId}/invite/{invitedId} [post]
func (c *GroupController) InviteUser(ctx *fiber.Ctx) error {
	groupParam := ctx.Params("groupId")
	groupIdUint64, err := strconv.ParseUint(groupParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid group ID",
		})
	}
	groupId := uint(groupIdUint64)

	invitedParam := ctx.Params("invitedId")
	invitedIdUint64, err := strconv.ParseUint(invitedParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid invited user ID",
		})
	}
	invitedId := uint(invitedIdUint64)

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	if err := c.groupService.InviteUser(userId, groupId, invitedId); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

// GetPendingInvites godoc
// @Summary Listar convites pendentes
// @Description Retorna os convites de grupos pendentes do usuário autenticado
// @Tags Groups
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param query query input.GetPendingInvites false "Parâmetros de filtro"
// @Success 200 {array} map[string]interface{} "Lista de convites pendentes"
// @Failure 400 {object} ValidationErrorResponse "Erro na validação dos parâmetros"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Router /api/group/invite [get]
func (c *GroupController) GetPendingInvites(ctx *fiber.Ctx) error {
	var query input.GetPendingInvites

	if err := ctx.QueryParser(&query); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid query parameters")
	}

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	response, err := c.groupService.GetPendingInvites(userId, &query)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

// AnswerPendingInvite godoc
// @Summary Responder convite pendente
// @Description Aceita ou rejeita um convite de grupo pendente
// @Tags Groups
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param groupId path int true "ID do grupo"
// @Param answer body input.AnswerPendingInvite true "Resposta ao convite (accept/reject)"
// @Success 204 "Resposta processada com sucesso"
// @Failure 400 {object} ValidationErrorResponse "Erro na validação dos dados"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Failure 404 {object} ErrorMessageResponse "Convite não encontrado"
// @Router /api/group/{groupId}/invite [patch]
func (c *GroupController) AnswerPendingInvite(ctx *fiber.Ctx) error {
	var answer input.AnswerPendingInvite

	if err := ctx.BodyParser(&answer); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body request")
	}

	groupParam := ctx.Params("groupId")
	groupIdUint64, err := strconv.ParseUint(groupParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid group ID",
		})
	}
	groupId := uint(groupIdUint64)

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	if err := c.groupService.AnswerPendingInvite(userId, groupId, &answer); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

// RemoveMember godoc
// @Summary Remover membro do grupo
// @Description Remove um membro de um grupo (apenas administradores)
// @Tags Groups
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param groupId path int true "ID do grupo"
// @Param memberId path int true "ID do membro a ser removido"
// @Success 204 "Membro removido com sucesso"
// @Failure 400 {object} ErrorMessageResponse "ID inválido"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Failure 403 {object} ErrorMessageResponse "Sem permissão para remover membro"
// @Failure 404 {object} ErrorMessageResponse "Grupo ou membro não encontrado"
// @Router /api/group/{groupId}/member/{memberId} [delete]
func (c *GroupController) RemoveMember(ctx *fiber.Ctx) error {
	groupParam := ctx.Params("groupId")
	groupIdUint64, err := strconv.ParseUint(groupParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid group ID",
		})
	}
	groupId := uint(groupIdUint64)

	memberParam := ctx.Params("memberId")
	memberIdUint64, err := strconv.ParseUint(memberParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid member ID",
		})
	}
	memberId := uint(memberIdUint64)

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	if err := c.groupService.RemoveMember(userId, groupId, memberId); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
