package controllers

import (
	"github.com/gin-gonic/gin"
	"task-5-vix-btpns-mumaralfajar/services"
)

type PhotoController interface {
	Insert(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	FindByID(ctx *gin.Context)
}

type photoController struct {
	photoService services.PhotoService
	jwtService   services.JwtService
}

func NewPhotoController(photoService services.PhotoService, jwtService services.JwtService) PhotoController {
	return &photoController{
		photoService: photoService,
		jwtService:   jwtService,
	}
}

func (c *photoController) Insert(ctx *gin.Context) {
	panic("implement me")
}

func (c *photoController) Update(ctx *gin.Context) {
	panic("implement me")
}

func (c *photoController) Delete(ctx *gin.Context) {
	panic("implement me")
}

func (c *photoController) FindAll(ctx *gin.Context) {
	panic("implement me")
}

func (c *photoController) FindByID(ctx *gin.Context) {
	panic("implement me")
}
