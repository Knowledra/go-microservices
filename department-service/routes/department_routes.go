package routes

import (
	"department/controllers"

	"github.com/gin-gonic/gin"
	"github.com/sushantpardhi/shared/middleware"
)

func departmentRoutes(app *gin.RouterGroup) {
	app.POST("/create", middleware.AuthMiddleware(), middleware.AdminOnly(), controllers.CreateDepartment)
}
