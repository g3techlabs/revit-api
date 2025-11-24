package controller

import (
	"github.com/g3techlabs/revit-api/src/core/auth/input"
	"github.com/g3techlabs/revit-api/src/core/auth/middleware"
	_ "github.com/g3techlabs/revit-api/src/core/auth/response"
	"github.com/g3techlabs/revit-api/src/core/auth/services"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	AuthService services.IAuthService
}

func NewAuthController(authService services.IAuthService) *AuthController {
	return &AuthController{
		AuthService: authService,
	}
}

// ValidationErrorResponse representa a resposta de erro de validação
// @Description Resposta retornada quando há erros de validação nos dados enviados. O campo "errors" contém um mapa onde a chave é o nome do campo (em minúsculas) e o valor é a mensagem de erro.
type ValidationErrorResponse struct {
	// Mapa de erros de validação, onde a chave é o nome do campo (em minúsculas) e o valor é a mensagem de erro
	Errors map[string]string `json:"errors"`
}

// ErrorMessageResponse representa uma resposta de erro simples com apenas uma mensagem
// @Description Resposta retornada quando ocorre um erro que não requer detalhes adicionais, apenas uma mensagem descritiva.
type ErrorMessageResponse struct {
	// Mensagem descritiva do erro
	Message string `json:"message"`
}

// ConflictErrorResponse representa a resposta de erro de conflito, onde o nickname ou email já estão em uso.
// @Description Resposta retornada quando o nickname ou email já estão em uso.
type ConflictErrorResponse struct {
	// Apelido único do usuário
	Nickname *string `json:"nickname" example:"nickname already in use"`
	// Email único do usuário
	Email *string `json:"email" example:"email already in use"`
}

// RegisterUser godoc
// @Summary Registrar novo usuário
// @Description Cria uma nova conta de usuário no sistema
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body input.CreateUser true "Dados do usuário para registro"
// @Success 201 {object} response.UserCreatedResponse "Usuário criado com sucesso"
// @Failure 400 {object} ValidationErrorResponse "Erro na validação dos dados. Possíveis campos: name, nickname, email, password, birthdate"
// @Failure 409 {object} ConflictErrorResponse "Alguns dos campos únicos já estão em uso"
// @Router /api/auth/register [post]
func (c *AuthController) RegisterUser(ctx *fiber.Ctx) error {
	input := new(input.CreateUser)

	if err := ctx.BodyParser(input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body:" + err.Error(),
		})
	}

	user, err := c.AuthService.RegisterUser(input)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"user": user,
	})
}

// Login godoc
// @Summary Autenticar usuário
// @Description Realiza login do usuário e retorna tokens de acesso e refresh
// @Tags Auth
// @Accept json
// @Produce json
// @Param credentials body input.LoginCredentials true "Credenciais de login (email ou nickname + senha + tipo de identificador)"
// @Success 200 {object} response.AuthTokensResponse "Tokens de autenticação"
// @Failure 400 {object} ValidationErrorResponse "Erro na validação dos dados. Possíveis campos: identifier, password, identifierType"
// @Failure 401 {object} ErrorMessageResponse "Credenciais inválidas"
// @Router /api/auth/login [post]
func (c AuthController) Login(ctx *fiber.Ctx) error {
	input := new(input.LoginCredentials)

	if err := ctx.BodyParser(input); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body" + err.Error(),
		})
	}

	authTokens, err := c.AuthService.Login(input)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(authTokens)
}

// RefreshTokens godoc
// @Summary Atualizar tokens de autenticação
// @Description Gera novos access token e refresh token, usando o refresh token atual
// @Tags Auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer {refreshToken}"
// @Success 200 {object} response.AuthTokensResponse "Novos tokens de autenticação"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Router /api/auth/refresh [post]
func (c AuthController) RefreshTokens(ctx *fiber.Ctx) error {
	refreshToken, err := middleware.ExtractBearerToken(ctx)
	if err != nil {
		return err
	}

	authTokens, err := c.AuthService.RefreshTokens(refreshToken)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(authTokens)
}

// SendPassResetEmail godoc
// @Summary Solicitar redefinição de senha
// @Description Envia email com token para redefinição de senha
// @Tags Auth
// @Accept json
// @Produce json
// @Param identifier body input.Identifier true "Email ou nickname do usuário"
// @Success 204 "Email enviado com sucesso"
// @Failure 400 {object} ValidationErrorResponse "Erro na validação dos dados. Possíveis campos: identifier, identifierType"
// @Failure 404 {object} ErrorMessageResponse "Usuário não encontrado"
// @Router /api/auth/password [post]
func (c AuthController) SendPassResetEmail(ctx *fiber.Ctx) error {
	input := new(input.Identifier)

	if err := ctx.BodyParser(input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body: " + err.Error(),
		})
	}

	err := c.AuthService.SendPassResetEmail(input)
	if err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

// ResetPassword godoc
// @Summary Redefinir senha
// @Description Redefine a senha do usuário usando o token de redefinição
// @Tags Auth
// @Accept json
// @Produce json
// @Param resetData body input.ResetPassword true "Token de redefinição e nova senha"
// @Success 204 "Senha redefinida com sucesso"
// @Failure 400 {object} ValidationErrorResponse "Erro na validação dos dados. Possíveis campos: resettoken, newpassword"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Router /api/auth/password [patch]
func (c AuthController) ResetPassword(ctx *fiber.Ctx) error {
	input := new(input.ResetPassword)

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body: " + err.Error(),
		})
	}

	if err := c.AuthService.ResetPassword(input); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c *AuthController) CheckIfEmailAvailable(ctx *fiber.Ctx) error {
	var input input.EmailInput

	if err := ctx.QueryParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body request")
	}

	isEmailAvailable, err := c.AuthService.CheckIfEmailAvailable(&input)
	if err != nil {
		return err
	}

	if !isEmailAvailable {
		return fiber.NewError(fiber.StatusConflict, "Email already taken")
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c *AuthController) CheckIfNicknameAvailable(ctx *fiber.Ctx) error {
	var input input.NicknameInput

	if err := ctx.QueryParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body request")
	}

	isEmailAvailable, err := c.AuthService.CheckIfNicknameAvailable(&input)
	if err != nil {
		return err
	}

	if !isEmailAvailable {
		return fiber.NewError(fiber.StatusConflict, "Nickname already taken")
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
