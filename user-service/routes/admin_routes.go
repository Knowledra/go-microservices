package routes

import (
	"user/controllers"

	"github.com/gin-gonic/gin"
	"github.com/sushantpardhi/shared/middleware"
)

func adminroutes(app *gin.RouterGroup) {
	create := app.Group("/create")
	create.POST("/super-admin", controllers.CreateSuperAdmin)
	create.POST("/admin", middleware.AuthMiddleware(), middleware.SuperAdminOnly(), controllers.CreateAdmin)
	create.POST("/teacher", middleware.AuthMiddleware(), middleware.AdminOnly(), controllers.CreateTeacher)
	create.POST("/student", middleware.AuthMiddleware(), middleware.AdminOnly(), controllers.CreateStudent)
	create.POST("/parent", middleware.AuthMiddleware(), middleware.AdminOnly(), controllers.CreateParent)

	app.PUT("/update/user/:role/:user_id", middleware.AuthMiddleware(), middleware.AdminOnly(), controllers.UpdateUser)
}
