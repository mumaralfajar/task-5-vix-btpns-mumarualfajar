package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"task-5-vix-btpns-mumaralfajar/dto"
	"task-5-vix-btpns-mumaralfajar/helpers"
	"task-5-vix-btpns-mumaralfajar/models"
	"task-5-vix-btpns-mumaralfajar/services"
)

type AuthController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type authController struct {
	authService services.AuthService
	jwtService  services.JwtService
}

func NewAuthController(authService services.AuthService, jwtService services.JwtService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDto
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helpers.APIErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(models.User); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = generatedToken
		response := helpers.APIResponse(true, "OK!", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helpers.APIErrorResponse("Please check again your credential", "Invalid Credential", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDto
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helpers.APIErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helpers.APIErrorResponse("Failed to process request", "Duplicate email", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		createdUser := c.authService.CreateUser(registerDTO)
		token := c.jwtService.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
		createdUser.Token = token
		response := helpers.APIResponse(true, "OK!", createdUser)
		ctx.JSON(http.StatusCreated, response)
	}
}
