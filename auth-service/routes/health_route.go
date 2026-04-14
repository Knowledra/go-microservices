package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sushantpardhi/shared/health"
)

func healthRoutes(app *gin.RouterGroup) {
	// app.GET("/health", middleware.AuthMiddleware(), middleware.AdminOnly(), health.GetHealth)
	app.GET("/health", health.GetHealth)
}
