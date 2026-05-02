package routes

import (
	"email/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init(app *gin.Engine) {
	app.GET("/", func(C *gin.Context) {
		C.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Email Service API",
		})
	})

	app.POST("/send-email", controller.SendEmail)

	api := app.Group("/api/v1")
	healthRoutes(api)

}
