package routes

import (
	"auth/controllers"

	"github.com/gin-gonic/gin"
	"github.com/sushantpardhi/shared/middleware"
)

func authRoutes(app *gin.RouterGroup) {
	authentication := app.Group("/auth")
	authentication.POST("/login", controllers.Login)
	authentication.POST("/logout", middleware.AuthMiddleware(), controllers.Logout)

	authentication.POST("/create/user", controllers.CreateUser)
	authentication.DELETE("/delete/user", middleware.AuthMiddleware(), controllers.DeleteUser)
}
