package routes

import (
	"user/controllers"

	"github.com/gin-gonic/gin"
)

func adminroutes(app *gin.RouterGroup) {
	profile := app.Group("/profile")
	profile.POST("/create/admin", controllers.CreateAdmin)
	profile.DELETE("/delete/admin/:user_id", controllers.DeleteAdmin)
}
