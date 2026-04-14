package routes

import (
	"auth/controllers"
	middleware "auth/middlewares"

	"github.com/gin-gonic/gin"
)

func authRoutes(app *gin.RouterGroup) {
	authentication := app.Group("/auth")
	authentication.POST("/login", controllers.Login)
	authentication.POST("/logout", middleware.AuthMiddleware(), controllers.Logout)

	authentication.POST("/create/admin", controllers.CreateAdmin)

}
