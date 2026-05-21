package routes

import (
	"user/controllers"

	"github.com/gin-gonic/gin"
	"github.com/sushantpardhi/shared/middleware"
)

func userRoutes(app *gin.RouterGroup) {
	app.PUT("/update/user/:role/:user_id", middleware.AuthMiddleware(), middleware.AdminOnly(), controllers.UpdateUser)
	app.PUT("/update/self", middleware.AuthMiddleware(), controllers.UpdateSelf)
}
