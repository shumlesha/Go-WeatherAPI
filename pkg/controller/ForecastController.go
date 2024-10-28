package controller

import (
	"WeatherfForecast/middleware"
	"WeatherfForecast/pkg/service"
	"WeatherfForecast/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	Forecasts = "/forecasts"
	Current   = "/current"
	Username  = "username"
)

type ForecastController struct {
	forecastService service.ForecastService
}

func NewForecastController(forecastService service.ForecastService) *ForecastController {
	return &ForecastController{
		forecastService: forecastService,
	}
}

func (c *ForecastController) RegisterRoutes(r *gin.Engine, middleware *middleware.AuthMiddleware) {
	forecast := r.Group(Forecasts).Use(middleware.AuthChecked())
	{
		forecast.GET(Current, c.GetCurrentForecast)
	}
}

func (c *ForecastController) GetCurrentForecast(context *gin.Context) {
	username, ok := context.Get(Username)

	if !ok {
		context.JSON(
			http.StatusUnauthorized,
			util.NewErrorResponse("Unauthorized", "User not found in context"))
		return
	}

	forecast, err := c.forecastService.GetCurrentForecast(username.(string))

	if err != nil {
		context.JSON(http.StatusInternalServerError, util.NewErrorResponse("Error fetching forecast", err.Error()))
		return
	}

	context.JSON(http.StatusOK, forecast)
}
