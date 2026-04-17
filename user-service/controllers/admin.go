package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sushantpardhi/shared/response"
)

func CreateAdmin(ctx *gin.Context) {
	response.Success(ctx, http.StatusOK, "Admin created successfully", nil)
}
