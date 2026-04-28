package logger

import (
	"context"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

// ---------- writer (daily rotation aligned to midnight) ----------
func getWriter(pattern string) zapcore.WriteSyncer {
	writer, err := rotatelogs.New(
		pattern, // ../logs/app-%Y-%m-%d.log

		// rotate every 24h
		rotatelogs.WithRotationTime(24*time.Hour),

		// align rotation to local midnight
		rotatelogs.WithClock(rotatelogs.Local),

		// keep logs for 7 days
		rotatelogs.WithMaxAge(7*24*time.Hour),
	)
	if err != nil {
		panic(err)
	}
	return zapcore.AddSync(writer)
}

// ---------- init ----------
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

	// daily rotating files
	allWriter := getWriter("../logs/" + filename + "-%Y-%m-%d.log")

	errorFilename := "error"
	if filename == "test" {
		errorFilename = "test_error"
	}
	errorWriter := getWriter("../logs/" + errorFilename + "-%Y-%m-%d.log")

	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
		zapcore.AddSync(os.Stdout),
		zapcore.DebugLevel,
	)

	allCore := zapcore.NewCore(encoder, allWriter, zapcore.InfoLevel)
	errorCore := zapcore.NewCore(encoder, errorWriter, zapcore.ErrorLevel)

	core := zapcore.NewTee(allCore, errorCore, consoleCore)

	log = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

// ---------- log helpers ----------
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

// ---------- context logger (request_id support) ----------
func C(c context.Context) *zap.Logger {
	if c == nil {
		return log
	}

	var reqID string

	if gc, ok := c.(*gin.Context); ok {
		reqID = gc.GetString("X-Request-Id")
	} else if val, ok := c.Value("X-Request-Id").(string); ok {
		reqID = val
	}

	if reqID != "" {
		return log.With(zap.String("request_id", reqID))
	}
	return log
}
