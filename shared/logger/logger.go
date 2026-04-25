package logger

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var log *zap.Logger

func Init(filename string) {
	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:      "time",
		LevelKey:     "level",
		MessageKey:   "msg",
		CallerKey:    "caller",
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		EncodeLevel:  zapcore.LowercaseLevelEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	})

	// ✅ Common log file (info + warn + error)
	allWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   `../logs/` + filename + `.log`,
		MaxSize:    1,
		MaxBackups: 7,
		MaxAge:     7,
		Compress:   true,
	})

	// ✅ Error-only log file
	errorFilename := "error"
	if filename == "test" {
		errorFilename = "test_error"
	}
	errorWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   `../logs/` + errorFilename + `.log`,
		MaxSize:    1,
		MaxBackups: 7,
		MaxAge:     7,
		Compress:   true,
	})

	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
		zapcore.AddSync(os.Stdout),
		zapcore.DebugLevel,
	)

	// Core for all logs
	allCore := zapcore.NewCore(encoder, allWriter, zapcore.InfoLevel)

	// Core only for errors
	errorCore := zapcore.NewCore(encoder, errorWriter, zapcore.ErrorLevel)

	// Combine both
	core := zapcore.NewTee(allCore, errorCore, consoleCore)

	log = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

func Info(msg string, fields ...zap.Field) {
	log.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	log.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	log.Error(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	log.Debug(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	log.Fatal(msg, fields...)
}

func C(c context.Context) *zap.Logger {
	if c == nil {
		return log
	}

	var reqID string

	// Check if it's a gin Context
	if c, ok := c.(*gin.Context); ok {
		reqID = c.GetString("X-Request-Id")
	} else if val, ok := c.Value("X-Request-Id").(string); ok {
		reqID = val
	}

	if reqID != "" {
		return log.With(zap.String("request_id", reqID))
	}
	return log
}
