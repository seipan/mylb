package utils

import (
	"go.uber.org/zap"
)

func newlogger() zap.Logger {
	logger := zap.NewExample()
	defer logger.Sync()
	return *logger
}

func Info(message string, fields ...zap.Field) {
	log := newlogger()
	log.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	log := newlogger()
	log.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	log := newlogger()
	log.Error(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
	log := newlogger()
	log.Fatal(message, fields...)
}

func Warn(message string, fields ...zap.Field) {
	log := newlogger()
	log.Warn(message, fields...)
}

func Panic(message string, fields ...zap.Field) {
	log := newlogger()
	log.Panic(message, fields...)
}
