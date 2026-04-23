package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init(app *gin.Engine) {
	app.GET("/", func(C *gin.Context) {
		C.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Auth Service API",
		})
	})

	api := app.Group("/api/v1/auth")
	healthRoutes(api)
	authRoutes(api)
}
