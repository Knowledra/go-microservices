package routes

import (
	"user/controllers"

	"github.com/gin-gonic/gin"
	"github.com/sushantpardhi/shared/middleware"
)

func adminroutes(app *gin.RouterGroup) {
	profile := app.Group("/profile")
	profile.POST("/create/admin", controllers.CreateAdmin)
	profile.POST("/create/teacher", controllers.CreateTeacher)
	profile.POST("/create/student", controllers.CreateStudent)
	profile.POST("/create/parent", controllers.CreateParent)

	profile.PUT("/update/admin", middleware.AuthMiddleware(), controllers.UpdateAdmin)
}
