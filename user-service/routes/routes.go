package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init(app *gin.Engine) {
	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome to User Service API",
		})
	})

	api := app.Group("/api/v1")
	healthRoutes(api)
	adminroutes(api)
}
