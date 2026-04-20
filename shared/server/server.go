package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sushantpardhi/shared/logger"
	"go.uber.org/zap"
)

var StartTime = time.Now()

type Server struct {
	httpServer *http.Server
}

func New(handler http.Handler, port string) *Server {
	// Create a new server
	return &Server{
		httpServer: &http.Server{
			Addr:    ":" + port,
			Handler: handler,
		},
	}
}

func (s *Server) Start() {
	// Start server in a goroutine
	go func() {
		logger.Info("Server starting", zap.String("addr", s.httpServer.Addr))
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Server listen error", zap.Error(err))
		}
	}()

	// Graceful shutdown
	s.gracefulShutdown()
}

func (s *Server) gracefulShutdown() {
	// Create a channel to receive signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit

	logger.Info("Server shutdown initiated", zap.String("signal", sig.String()))

	// Create a context with a timeout
	C, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown the server
	if err := s.httpServer.Shutdown(C); err != nil {
		logger.Error("Server shutdown error", zap.Error(err))
	} else {
		logger.Info("Server shutdown completed")
	}

	logger.Info("Server exited")
}
