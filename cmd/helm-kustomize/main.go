package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.SugaredLogger

	cfg config
)

func main() {
	logger = InitLogger()

	cfg.loadConfigs()
	defer logger.Sync()
}

// InitLogger - initilize zap logger
func InitLogger() *zap.SugaredLogger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	corelogger, _ := config.Build()

	return corelogger.Sugar()
}
