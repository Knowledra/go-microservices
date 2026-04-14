package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init(app *gin.Engine) {
	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome to EduPlatform API",
		})
	})

	api := app.Group("/api/v1")
	// adminRoutes(api)
	healthRoutes(api)
	authRoutes(api)
}
