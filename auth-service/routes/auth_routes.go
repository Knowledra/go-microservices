package routes

import (
	"auth/controllers"

	"github.com/gin-gonic/gin"
	"github.com/sushantpardhi/shared/middleware"
)

func authRoutes(app *gin.RouterGroup) {
	app.POST("/login", controllers.Login)
	app.POST("/logout", middleware.AuthMiddleware(), controllers.Logout)

	app.POST("/create/super-admin", controllers.CreateSuperAdmin)
	app.POST("/create/user", middleware.AuthMiddleware(), controllers.CreateUser)
	app.DELETE("/delete/user/:user_id", middleware.AuthMiddleware(), middleware.AdminOnly(), controllers.DeleteUser)
	app.POST("/reset-password", middleware.AuthMiddleware(), controllers.ResetPassword)
}
