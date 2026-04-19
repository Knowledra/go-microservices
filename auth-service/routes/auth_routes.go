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

	authentication.POST("/create/admin", controllers.CreateAdmin)
	authentication.DELETE("/delete/admin", middleware.AuthMiddleware(), controllers.DeleteAdmin)
}
