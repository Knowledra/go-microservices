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
	logger.Info("Getting health status")
	start := time.Now()

	service := os.Getenv("SERVICE_NAME")

	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	dbStatus := "up"
	sqlDB, err := db.DB.DB()
	if err != nil {
		dbStatus = "down"
		logger.Error("DB connection error: " + err.Error())
	} else if err := sqlDB.PingContext(ctx); err != nil {
		dbStatus = "down"
		logger.Error("DB ping failed: " + err.Error())
	}

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	uptime := time.Since(server.StartTime)

	status := "up"
	if dbStatus == "down" {
		status = "degraded"
		logger.Error("Health degraded: database down")
	}

	logger.Info("Health check completed in " + time.Since(start).String())

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
