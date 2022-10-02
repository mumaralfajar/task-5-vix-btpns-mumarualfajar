package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strconv"
	"task-5-vix-btpns-mumaralfajar/dto"
	"task-5-vix-btpns-mumaralfajar/helpers"
	"task-5-vix-btpns-mumaralfajar/services"
)

type UserController interface {
	Update(c *gin.Context)
	Profile(c *gin.Context)
}

type userController struct {
	userService services.UserService
	jwtService  services.JwtService
}

func NewUserController(userService services.UserService, jwtService services.JwtService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *userController) Update(ctx *gin.Context) {
	var userUpdateDTO dto.UserDto
	errDTO := ctx.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		res := helpers.APIErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	userUpdateDTO.ID = id
	u := c.userService.Update(userUpdateDTO)
	res := helpers.APIResponse(true, "OK!", u)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) Profile(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	user := c.userService.Profile(id)
	res := helpers.APIResponse(true, "OK", user)
	context.JSON(http.StatusOK, res)
}
