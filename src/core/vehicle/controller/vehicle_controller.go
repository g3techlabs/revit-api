package controller

import (
	"strconv"

	"github.com/g3techlabs/revit-api/src/core/vehicle/input"
	_ "github.com/g3techlabs/revit-api/src/core/vehicle/response"
	"github.com/g3techlabs/revit-api/src/core/vehicle/service"
	"github.com/g3techlabs/revit-api/src/response/generics"
	"github.com/gofiber/fiber/v2"
)

type VehicleController struct {
	vehicleService service.IVehicleService
}

func NewVehicleController(vehicleService service.IVehicleService) *VehicleController {
	return &VehicleController{vehicleService: vehicleService}
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

// CreateVehicle godoc
// @Summary Criar veículo
// @Description Cria um novo veículo para o usuário autenticado
// @Tags Vehicles
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param vehicle body input.CreateVehicle true "Dados do veículo"
// @Success 201 {object} response.PresignedPhotoInfo "URL pré-assinada para upload da foto principal"
// @Failure 400 {object} ValidationErrorResponse "Erro na validação dos dados"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Router /api/vehicle [post]
func (c *VehicleController) CreateVehicle(ctx *fiber.Ctx) error {
	input := new(input.CreateVehicle)

	if err := ctx.BodyParser(input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body request: "+err.Error())
	}

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	response, err := c.vehicleService.CreateVehicle(userId, input)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

// GetVehicles godoc
// @Summary Listar veículos
// @Description Retorna a lista de veículos do usuário autenticado
// @Tags Vehicles
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param query query input.GetVehiclesParams false "Parâmetros de filtro"
// @Success 200 {array} response.Vehicle "Lista de veículos"
// @Failure 400 {object} ValidationErrorResponse "Erro na validação dos parâmetros"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Router /api/vehicle [get]
func (c *VehicleController) GetVehicles(ctx *fiber.Ctx) error {
	query := new(input.GetVehiclesParams)

	if err := ctx.QueryParser(query); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid query parameters")
	}

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	vehicles, err := c.vehicleService.GetVehicles(userId, query)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(vehicles)
}

// UpdateVehicleInfo godoc
// @Summary Atualizar informações do veículo
// @Description Atualiza as informações de um veículo
// @Tags Vehicles
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param vehicleId path int true "ID do veículo"
// @Param vehicle body input.UpdateVehicleInfo true "Dados do veículo para atualização"
// @Success 204 "Veículo atualizado com sucesso"
// @Failure 400 {object} ValidationErrorResponse "Erro na validação dos dados"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Failure 403 {object} ErrorMessageResponse "Sem permissão para atualizar o veículo"
// @Failure 404 {object} ErrorMessageResponse "Veículo não encontrado"
// @Router /api/vehicle/{vehicleId} [patch]
func (c *VehicleController) UpdateVehicleInfo(ctx *fiber.Ctx) error {
	vehicleParam := ctx.Params("vehicleId")

	vehicleIdUint64, err := strconv.ParseUint(vehicleParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid vehicle ID",
		})
	}
	vehicleId := uint(vehicleIdUint64)

	input := new(input.UpdateVehicleInfo)
	if err := ctx.BodyParser(input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	if err := c.vehicleService.UpdateVehicleInfo(vehicleId, input); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

// RequestPhotoUpsert godoc
// @Summary Solicitar upload de foto do veículo
// @Description Solicita uma URL pré-assinada para upload de foto do veículo
// @Tags Vehicles
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param vehicleId path int true "ID do veículo"
// @Param request body input.RequestPhotoUpsert true "Dados da requisição de foto"
// @Success 200 {object} response.PresignedPhotoInfo "URL pré-assinada para upload"
// @Failure 400 {object} ValidationErrorResponse "Erro na validação dos dados"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Failure 403 {object} ErrorMessageResponse "Sem permissão para atualizar o veículo"
// @Failure 404 {object} ErrorMessageResponse "Veículo não encontrado"
// @Router /api/vehicle/photo/{vehicleId} [post]
func (c *VehicleController) RequestPhotoUpsert(ctx *fiber.Ctx) error {
	vehicleParam := ctx.Params("vehicleId")

	vehicleIdUint64, err := strconv.ParseUint(vehicleParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid vehicle ID",
		})
	}
	vehicleId := uint(vehicleIdUint64)

	input := new(input.RequestPhotoUpsert)
	if err := ctx.BodyParser(input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	response, err := c.vehicleService.RequestPhotoUpsert(userId, vehicleId, input)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

// ConfirmNewPhoto godoc
// @Summary Confirmar upload de foto do veículo
// @Description Confirma o upload de foto após upload na URL pré-assinada
// @Tags Vehicles
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param vehicleId path int true "ID do veículo"
// @Param confirmation body ConfirmNewPhotoRequest true "Confirmação do upload"
// @Success 204 "Foto confirmada com sucesso"
// @Failure 400 {object} ValidationErrorResponse "Erro na validação dos dados"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Failure 403 {object} ErrorMessageResponse "Sem permissão para atualizar o veículo"
// @Failure 404 {object} ErrorMessageResponse "Veículo não encontrado"
// @Router /api/vehicle/photo/{vehicleId} [patch]
func (c *VehicleController) ConfirmNewPhoto(ctx *fiber.Ctx) error {
	vehicleParam := ctx.Params("vehicleId")

	vehicleIdUint64, err := strconv.ParseUint(vehicleParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid vehicle ID",
		})
	}
	vehicleId := uint(vehicleIdUint64)

	input := new(input.ConfirmNewPhoto)
	if err := ctx.BodyParser(input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	if err := c.vehicleService.ConfirmNewPhoto(userId, vehicleId, input); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

// DeleteVehicle godoc
// @Summary Deletar veículo
// @Description Remove um veículo do usuário autenticado
// @Tags Vehicles
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param vehicleId path int true "ID do veículo"
// @Success 204 "Veículo deletado com sucesso"
// @Failure 400 {object} ErrorMessageResponse "ID inválido"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Failure 403 {object} ErrorMessageResponse "Sem permissão para deletar o veículo"
// @Failure 404 {object} ErrorMessageResponse "Veículo não encontrado"
// @Router /api/vehicle/{vehicleId} [delete]
func (c *VehicleController) DeleteVehicle(ctx *fiber.Ctx) error {
	vehicleParam := ctx.Params("vehicleId")

	vehicleIdUint64, err := strconv.ParseUint(vehicleParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid vehicle ID",
		})
	}
	vehicleId := uint(vehicleIdUint64)

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	if err := c.vehicleService.DeleteVehicle(userId, vehicleId); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

// RemoveMainPhoto godoc
// @Summary Remover foto principal do veículo
// @Description Remove a foto principal de um veículo
// @Tags Vehicles
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param vehicleId path int true "ID do veículo"
// @Success 204 "Foto principal removida com sucesso"
// @Failure 400 {object} ErrorMessageResponse "ID inválido"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Failure 403 {object} ErrorMessageResponse "Sem permissão para atualizar o veículo"
// @Failure 404 {object} ErrorMessageResponse "Veículo não encontrado"
// @Router /api/vehicle/main-photo/{vehicleId} [delete]
func (c *VehicleController) RemoveMainPhoto(ctx *fiber.Ctx) error {
	vehicleParam := ctx.Params("vehicleId")

	vehicleIdUint64, err := strconv.ParseUint(vehicleParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid vehicle ID",
		})
	}
	vehicleId := uint(vehicleIdUint64)

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	if err := c.vehicleService.RemoveMainPhoto(userId, vehicleId); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

// RemovePhoto godoc
// @Summary Remover foto do veículo
// @Description Remove uma foto específica de um veículo
// @Tags Vehicles
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param vehicleId path int true "ID do veículo"
// @Param photoId path int true "ID da foto"
// @Success 204 "Foto removida com sucesso"
// @Failure 400 {object} ErrorMessageResponse "ID inválido"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Failure 403 {object} ErrorMessageResponse "Sem permissão para atualizar o veículo"
// @Failure 404 {object} ErrorMessageResponse "Veículo ou foto não encontrado"
// @Router /api/vehicle/{vehicleId}/photo/{photoId} [delete]
func (c *VehicleController) RemovePhoto(ctx *fiber.Ctx) error {
	vehicleParam := ctx.Params("vehicleId")

	vehicleIdUint64, err := strconv.ParseUint(vehicleParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid vehicle ID",
		})
	}
	vehicleId := uint(vehicleIdUint64)

	photoParam := ctx.Params("photoId")

	photoIdUint64, err := strconv.ParseUint(photoParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid photo ID",
		})
	}
	photoId := uint(photoIdUint64)

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	if err := c.vehicleService.RemovePhoto(userId, vehicleId, photoId); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
