package main

import (
	"github.com/gin-gonic/gin"
	"task-5-vix-btpns-mumaralfajar/controllers"
	"task-5-vix-btpns-mumaralfajar/database"
	"task-5-vix-btpns-mumaralfajar/middlewares"
	"task-5-vix-btpns-mumaralfajar/repositories"
	"task-5-vix-btpns-mumaralfajar/services"
)

var (
	db              = database.SetupDatabaseConnection()
	userRepository  = repositories.NewUserRepository(db)
	photoRepository = repositories.NewPhotoRepository(db)
	jwtService      = services.NewJwtService()
	authService     = services.NewAuthService(userRepository)
	userService     = services.NewUserService(userRepository)
	photoService    = services.NewPhotoService(photoRepository)
	authController  = controllers.NewAuthController(authService, jwtService)
	userController  = controllers.NewUserController(userService, jwtService)
	photoController = controllers.NewPhotoController(photoService, jwtService)
)

func main() {
	defer database.CloseDatabaseConnection(database.SetupDatabaseConnection())
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("api/user", middlewares.AuthMiddleware(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}
}
