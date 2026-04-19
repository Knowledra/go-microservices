package routes

import (
	"user/controllers"

	"github.com/gin-gonic/gin"
	"github.com/sushantpardhi/shared/middleware"
)

func adminroutes(app *gin.RouterGroup) {
	profile := app.Group("/profile")
	profile.POST("/create/admin", controllers.CreateAdmin)
	profile.DELETE("/delete/admin/:user_id", middleware.AuthMiddleware(), controllers.DeleteAdmin)
	profile.PUT("/update/admin", middleware.AuthMiddleware(), controllers.UpdateAdmin)
}
