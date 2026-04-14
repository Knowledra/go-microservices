package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sushantpardhi/shared/logger"
	"github.com/sushantpardhi/shared/response"
)

func Logout(c *gin.Context) {
	// Delete cookie
	c.SetCookie("access_token", "", -1, "/", "", true, true)
	logger.Info("Logout Successful")
	response.Success(c, http.StatusOK, "Logout successful", nil)
}
