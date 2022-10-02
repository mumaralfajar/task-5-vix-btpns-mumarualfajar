package controller

import "github.com/gin-gonic/gin"

type AuthController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type authController struct {
	gin.Context
}

func NewAuthController() AuthController {
	return &authController{}
}

func (c *authController) Login(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "hello login",
	})
}

func (c *authController) Register(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "hello register",
	})
}
