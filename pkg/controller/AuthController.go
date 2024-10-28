package controller

import (
	"WeatherfForecast/dto/user"
	"WeatherfForecast/pkg/service"
	"WeatherfForecast/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

const (
	Auth     = "/auth"
	Register = "/register"
	Login    = "/login"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (c *AuthController) RegisterRoutes(r *gin.Engine) {
	auth := r.Group(Auth)
	{
		auth.POST(Register, c.Register)
		auth.POST(Login, c.Login)
	}
}

func (c *AuthController) Register(context *gin.Context) {
	var registerUserRequest user.RegisterUserRequest

	if err := context.ShouldBindJSON(&registerUserRequest); err != nil {
		context.JSON(http.StatusBadRequest, util.NewErrorResponse("Invalid request", err.Error()))
		return
	}

	validate := validator.New()
	if err := validate.Struct(&registerUserRequest); err != nil {
		context.JSON(http.StatusBadRequest, util.NewValidationErrorResponse(err))
		return
	}

	registeredUser, err := c.authService.Register(registerUserRequest)

	if err != nil {
		context.JSON(http.StatusBadRequest, util.NewErrorResponse("Registration failed", err.Error()))
		return
	}

	context.JSON(http.StatusCreated, registeredUser.ToUserDto())
}

func (c *AuthController) Login(context *gin.Context) {
	var loginUserRequest user.LoginUserRequest
	if err := context.ShouldBindJSON(&loginUserRequest); err != nil {
		context.JSON(http.StatusBadRequest, util.NewErrorResponse("Invalid request", err.Error()))
		return
	}

	validate := validator.New()
	if err := validate.Struct(&loginUserRequest); err != nil {
		context.JSON(http.StatusBadRequest, util.NewValidationErrorResponse(err))
		return
	}

	token, err := c.authService.Login(loginUserRequest)
	if err != nil {
		context.JSON(http.StatusBadRequest, util.NewErrorResponse("Login failed", err.Error()))
		return
	}

	context.JSON(http.StatusOK, user.NewTokenDto(token))
}
