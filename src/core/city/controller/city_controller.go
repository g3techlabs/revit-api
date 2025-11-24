package controller

import (
	"github.com/g3techlabs/revit-api/src/core/city/input"
	_ "github.com/g3techlabs/revit-api/src/core/city/response"
	"github.com/g3techlabs/revit-api/src/core/city/service"
	"github.com/gofiber/fiber/v2"
)

type CityController struct {
	cityService service.ICityService
}

func NewCityController(cityService service.ICityService) *CityController {
	return &CityController{cityService: cityService}
}

// ValidationErrorResponse representa a resposta de erro de validação
type ValidationErrorResponse struct {
	Errors map[string]string `json:"errors"`
}

// ErrorMessageResponse representa uma resposta de erro simples
type ErrorMessageResponse struct {
	Message string `json:"message"`
}

// GetCities godoc
// @Summary Listar cidades
// @Description Retorna uma lista de cidades filtradas por nome, com suporte a paginação
// @Tags Cities
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param name query string true "Nome da cidade para busca"
// @Param page query int false "Número da página (opcional)"
// @Param limit query int false "Limite de resultados por página (opcional)"
// @Success 200 {array} response.GetCityReponse "Lista de cidades"
// @Failure 400 {object} ValidationErrorResponse "Erro na validação dos parâmetros. Possíveis campos: name, page, limit"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Router /api/city [get]
func (c *CityController) GetCities(ctx *fiber.Ctx) error {
	var query input.GetCitiesFilters

	if err := ctx.QueryParser(&query); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid query parameters")
	}

	response, err := c.cityService.GetCities(&query)
	if err != nil {
		return err
	}

	return ctx.JSON(response)
}

// GetNearbyCities godoc
// @Summary Listar cidades próximas
// @Description Retorna uma lista de cidades próximas a uma localização geográfica (latitude e longitude)
// @Tags Cities
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param latitude query number true "Latitude da localização (-90 a 90)"
// @Param longitude query number true "Longitude da localização (-180 a 180)"
// @Success 200 {array} response.GetCityReponse "Lista de cidades próximas"
// @Failure 400 {object} ValidationErrorResponse "Erro na validação dos parâmetros. Possíveis campos: latitude, longitude"
// @Failure 401 {object} ErrorMessageResponse "Token inválido ou expirado"
// @Router /api/city/nearby [get]
func (c *CityController) GetNearbyCities(ctx *fiber.Ctx) error {
	var query input.GetNearbyCitiesFilters

	if err := ctx.QueryParser(&query); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid query parameters")
	}

	response, err := c.cityService.GetNearbyCities(&query)
	if err != nil {
		return err
	}

	return ctx.JSON(response)
}
