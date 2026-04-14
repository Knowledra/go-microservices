package response

import (
	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code,omitempty"`
	Data    any    `json:"data"`
}

func Success(c *gin.Context, code int, message string, data any) {
	c.JSON(200, SuccessResponse{
		Success: true,
		Message: message,
		Code:    code,
		Data:    data,
	})
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code,omitempty"`
	Error   error  `json:"error"`
}

func Error(c *gin.Context, code int, msg string, err error) {
	c.JSON(code, ErrorResponse{
		Success: false,
		Code:    code,
		Message: msg,
		Error:   err,
	})
}
