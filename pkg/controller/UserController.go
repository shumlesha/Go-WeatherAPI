package controller

import (
	"WeatherfForecast/dto/user"
	"WeatherfForecast/middleware"
	"WeatherfForecast/pkg/service"
	"WeatherfForecast/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

const (
	Users      = "/users"
	MyProfile  = "/my-profile"
	ChangeCity = "/change-city"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (c *UserController) RegisterRoutes(r *gin.Engine, middleware *middleware.AuthMiddleware) {
	userGroup := r.Group(Users).Use(middleware.AuthChecked())
	{
		userGroup.GET(MyProfile, c.GetMyProfile)
		userGroup.PATCH(ChangeCity, c.ChangeCity)
	}
}

func (c *UserController) GetMyProfile(context *gin.Context) {
	username, ok := context.Get(Username)

	if !ok {
		context.JSON(
			http.StatusUnauthorized,
			util.NewErrorResponse("Unauthorized", "User not found in context"))
		return
	}

	retrievedUser, err := c.userService.GetMyProfile(username.(string))

	if err != nil {
		context.JSON(http.StatusInternalServerError, util.NewErrorResponse("Error fetching user", err.Error()))
		return
	}

	context.JSON(http.StatusOK, retrievedUser.ToUserDto())
}

func (c *UserController) ChangeCity(context *gin.Context) {
	var changeCityRequest user.ChangeCityRequest

	username, ok := context.Get(Username)

	if !ok {
		context.JSON(
			http.StatusUnauthorized,
			util.NewErrorResponse("Unauthorized", "User not found in context"))
		return
	}

	if err := context.ShouldBindJSON(&changeCityRequest); err != nil {
		context.JSON(http.StatusBadRequest, util.NewErrorResponse("Invalid request", err.Error()))
		return
	}

	validate := validator.New()
	if err := validate.Struct(&changeCityRequest); err != nil {
		context.JSON(http.StatusBadRequest, util.NewValidationErrorResponse(err))
		return
	}

	updatedUser, err := c.userService.ChangeCity(username.(string), changeCityRequest)

	if err != nil {
		context.JSON(http.StatusBadRequest, util.NewErrorResponse("City change failed", err.Error()))
		return
	}

	context.JSON(http.StatusOK, updatedUser.ToUserDto())
}
