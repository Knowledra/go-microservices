package routes

import (
	"user/controllers"

	"github.com/gin-gonic/gin"
	"github.com/sushantpardhi/shared/middleware"
)

func adminroutes(app *gin.RouterGroup) {
	app.POST("/create/admin", controllers.CreateAdmin)
	app.POST("/create/teacher", controllers.CreateTeacher)
	app.POST("/create/student", controllers.CreateStudent)
	app.POST("/create/parent", controllers.CreateParent)

	app.PUT("/update/user/:role/:user_id", middleware.AuthMiddleware(), middleware.AdminOnly(), controllers.UpdateUser)
}
