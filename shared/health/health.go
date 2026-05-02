package health

import (
	"context"
	"net/http"
	"os"

	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sushantpardhi/shared/db"
	"github.com/sushantpardhi/shared/logger"
	"github.com/sushantpardhi/shared/response"
	"github.com/sushantpardhi/shared/server"
)

func GetHealth(c *gin.Context) {
	start := time.Now()

	service := os.Getenv("SERVICE_NAME")

	C, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	dbStatus := "skipped"
	if service != "email-service" {
		dbStatus = "up"
		sqlDB, err := db.DB.DB()
		if err != nil {
			dbStatus = "down"
			logger.C(c).Error("DB connection error: " + err.Error())
		} else if err := sqlDB.PingContext(C); err != nil {
			dbStatus = "down"
			logger.C(c).Error("DB ping failed: " + err.Error())
		}
	}

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	uptime := time.Since(server.StartTime)

	status := "up"
	if dbStatus == "down" {
		status = "degraded"
		logger.C(c).Error("Health degraded: database down")
	}

	logger.C(c).Info("Health check completed in " + time.Since(start).String() + ", status: " + status)

	response.Success(c, http.StatusOK, "Health check completed", gin.H{
		"status":  status,
		"service": service,
		"checks": gin.H{
			"database": dbStatus,
		},
		"system": gin.H{
			"uptime":       uptime.String(),
			"goroutines":   runtime.NumGoroutine(),
			"memory_alloc": m.Alloc,
			"memory_total": m.TotalAlloc,
			"memory_sys":   m.Sys,
		},
		"response_time": time.Since(start).String(),
	})
}
