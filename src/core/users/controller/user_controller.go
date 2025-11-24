package controller

import (
	"strconv"

	"github.com/g3techlabs/revit-api/src/core/users/input"
	_ "github.com/g3techlabs/revit-api/src/core/users/response"
	"github.com/g3techlabs/revit-api/src/core/users/service"
	"github.com/g3techlabs/revit-api/src/response/generics"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{userService: userService}
}

// ValidationErrorResponse representa a resposta de erro de validação
// @Description Resposta retornada quando há erros de validação nos dados enviados
type ValidationErrorResponse struct {
	Errors map[string]string `json:"errors"`
}

// ErrorMessageResponse representa uma resposta de erro simples com apenas uma mensagem
// @Description Resposta retornada quando ocorre um erro que não requer detalhes adicionais
type ErrorMessageResponse struct {
	Message string `json:"message"`
}

// UpdateUser godoc
// @Summary Atualizar dados do usuário
// @Description Atualiza os dados do usuário autenticado
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user body input.UpdateUser true "Dados do usuário para atualização"
// @Success 204 "Usuário atualizado com sucesso"
// @Failure 400 {object} ValidationErrorResponse "Erro na validação dos dados. Possíveis campos: name, birthdate, removeProfilePic"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Router /api/user [patch]
func (uc *UserController) UpdateUser(ctx *fiber.Ctx) error {
	input := new(input.UpdateUser)

	if err := ctx.BodyParser(input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body: " + err.Error(),
		})
	}

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	if err := uc.userService.Update(userId, input); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

// RequestProfilePicUpdate godoc
// @Summary Solicitar atualização de foto de perfil
// @Description Solicita uma URL pré-assinada para upload de nova foto de perfil
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body input.RequestProfilePicUpdate true "Dados da requisição de atualização de foto"
// @Success 201 {object} response.ProfilePicPresignedURL "URL pré-assinada para upload"
// @Failure 400 {object} ValidationErrorResponse "Erro na validação dos dados"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Router /api/user/profile-pic/ [post]
func (uc *UserController) RequestProfilePicUpdate(ctx *fiber.Ctx) error {
	input := new(input.RequestProfilePicUpdate)

	if err := ctx.BodyParser(input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body: " + err.Error(),
		})
	}

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	response, err := uc.userService.RequestProfilePicUpdate(userId, input)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

// ConfirmNewProfilePic godoc
// @Summary Confirmar nova foto de perfil
// @Description Confirma o upload de uma nova foto de perfil após o upload na URL pré-assinada
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param confirmation body input.ConfirmNewProfilePic true "Confirmação do upload"
// @Success 204 "Foto de perfil confirmada com sucesso"
// @Failure 400 {object} ValidationErrorResponse "Erro na validação dos dados"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Router /api/user/profile-pic [patch]
func (uc *UserController) ConfirmNewProfilePic(ctx *fiber.Ctx) error {
	input := new(input.ConfirmNewProfilePic)

	if err := ctx.BodyParser(input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body: " + err.Error(),
		})
	}

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	if err := uc.userService.ConfirmNewProfilePic(userId, input); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

// GetUsers godoc
// @Summary Listar usuários
// @Description Retorna uma lista de usuários com filtros opcionais
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param query query input.GetUsersQuery false "Parâmetros de filtro e paginação"
// @Success 200 {array} response.GetUserResponse "Lista de usuários"
// @Failure 400 {object} ValidationErrorResponse "Erro na validação dos parâmetros"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Router /api/user [get]
func (uc *UserController) GetUsers(ctx *fiber.Ctx) error {
	query := new(input.GetUsersQuery)

	if err := ctx.QueryParser(query); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid query parameters: " + err.Error(),
		})
	}

	users, err := uc.userService.GetUsers(query)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(users)
}

// GetUser godoc
// @Summary Obter usuário por ID
// @Description Retorna os dados de um usuário específico
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID do usuário"
// @Success 200 {object} response.GetUserResponse "Dados do usuário"
// @Failure 400 {object} ErrorMessageResponse "ID inválido"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Failure 404 {object} ErrorMessageResponse "Usuário não encontrado"
// @Router /api/user/{id} [get]
func (uc *UserController) GetUser(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")

	idUint64, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}
	id := uint(idUint64)

	response, err := uc.userService.GetUser(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

// GetFriends godoc
// @Summary Listar amigos
// @Description Retorna a lista de amigos do usuário autenticado
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param query query input.GetUsersQuery false "Parâmetros de filtro e paginação"
// @Success 200 {array} response.Friend "Lista de amigos"
// @Failure 400 {object} ValidationErrorResponse "Erro na validação dos parâmetros"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Router /api/user/friendship [get]
func (uc *UserController) GetFriends(ctx *fiber.Ctx) error {
	query := new(input.GetUsersQuery)

	if err := ctx.QueryParser(query); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid query parameters: " + err.Error(),
		})
	}

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	friends, err := uc.userService.GetFriends(userId, query)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(friends)
}

// RequestFriendship godoc
// @Summary Solicitar amizade
// @Description Envia uma solicitação de amizade para outro usuário
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param destinataryId path int true "ID do usuário destinatário"
// @Success 204 "Solicitação de amizade enviada"
// @Failure 400 {object} ErrorMessageResponse "ID inválido"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Failure 404 {object} ErrorMessageResponse "Usuário destinatário não encontrado"
// @Router /api/user/friendship/{destinataryId} [post]
func (uc *UserController) RequestFriendship(ctx *fiber.Ctx) error {
	destinataryParam := ctx.Params("destinataryId")

	destinataryUint64, err := strconv.ParseUint(destinataryParam, 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid destinatary ID")
	}
	destinataryId := uint(destinataryUint64)

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	if err := uc.userService.RequestFriendship(userId, destinataryId); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

// AnswerFriendshipRequest godoc
// @Summary Responder solicitação de amizade
// @Description Aceita ou rejeita uma solicitação de amizade recebida
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param requesterId path int true "ID do usuário que enviou a solicitação"
// @Param answer body input.FriendshipRequestAnswer true "Resposta à solicitação (accept/reject)"
// @Success 204 "Resposta processada com sucesso"
// @Failure 400 {object} ValidationErrorResponse "Erro na validação dos dados"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Failure 404 {object} ErrorMessageResponse "Solicitação não encontrada"
// @Router /api/user/friendship/{requesterId} [patch]
func (uc *UserController) AnswerFriendshipRequest(ctx *fiber.Ctx) error {
	requesterParam := ctx.Params("requesterId")

	requesterUint64, err := strconv.ParseUint(requesterParam, 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid requester ID")
	}
	requesterId := uint(requesterUint64)

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	answer := new(input.FriendshipRequestAnswer)
	if err := ctx.BodyParser(answer); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	response, err := uc.userService.AnswerFriendshipRequest(userId, requesterId, answer)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusNoContent).JSON(response)
}

// RemoveFriendship godoc
// @Summary Remover amizade
// @Description Remove a amizade entre o usuário autenticado e outro usuário
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param friendId path int true "ID do amigo a ser removido"
// @Success 204 "Amizade removida com sucesso"
// @Failure 400 {object} ErrorMessageResponse "ID inválido"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Failure 404 {object} ErrorMessageResponse "Amizade não encontrada"
// @Router /api/user/friendship/{friendId} [delete]
func (uc *UserController) RemoveFriendship(ctx *fiber.Ctx) error {
	friendParam := ctx.Params("friendId")

	friendUint64, err := strconv.ParseUint(friendParam, 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid friend ID")
	}
	friendId := uint(friendUint64)

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	if err := uc.userService.RemoveFriendship(userId, friendId); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

// GetFriendshipRequests godoc
// @Summary Listar solicitações de amizade
// @Description Retorna as solicitações de amizade pendentes do usuário autenticado
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param query query input.GetFriendshipRequestsQuery false "Parâmetros de filtro"
// @Success 200 {array} response.FriendshipRequest "Lista de solicitações de amizade"
// @Failure 400 {object} ValidationErrorResponse "Erro na validação dos parâmetros"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Router /api/user/friendship/requests [get]
func (uc *UserController) GetFriendshipRequests(ctx *fiber.Ctx) error {
	query := new(input.GetFriendshipRequestsQuery)

	if err := ctx.QueryParser(query); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid query parameters")
	}

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	response, err := uc.userService.GetFriendshipRequests(userId, query)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

// CheckIfEmailAvailable godoc
// @Summary Verificar disponibilidade de email
// @Description Verifica se um email está disponível para uso
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param email query string true "Email a ser verificado"
// @Success 204 "Email disponível"
// @Failure 400 {object} ValidationErrorResponse "Erro na validação do email"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Failure 409 {object} ErrorMessageResponse "Email já está em uso"
// @Router /api/user/email-available [get]
func (c *UserController) CheckIfEmailAvailable(ctx *fiber.Ctx) error {
	var input input.EmailInput

	if err := ctx.QueryParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body request")
	}

	isEmailAvailable, err := c.userService.CheckIfEmailAvailable(&input)
	if err != nil {
		return err
	}

	if !isEmailAvailable {
		return fiber.NewError(fiber.StatusConflict, "Email already taken")
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

// CheckIfNicknameAvailable godoc
// @Summary Verificar disponibilidade de nickname
// @Description Verifica se um nickname está disponível para uso
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param nickname query string true "Nickname a ser verificado"
// @Success 204 "Nickname disponível"
// @Failure 400 {object} ValidationErrorResponse "Erro na validação do nickname"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Failure 409 {object} ErrorMessageResponse "Nickname já está em uso"
// @Router /api/user/nickname-available [get]
func (c *UserController) CheckIfNicknameAvailable(ctx *fiber.Ctx) error {
	var input input.NicknameInput

	if err := ctx.QueryParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body request")
	}

	isEmailAvailable, err := c.userService.CheckIfNicknameAvailable(&input)
	if err != nil {
		return err
	}

	if !isEmailAvailable {
		return fiber.NewError(fiber.StatusConflict, "Nickname already taken")
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
