package main

import (
	"os"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestMain(m *testing.M) {
	logger = InitTestLogger()
	cfg.loadConfigs()
	rc := m.Run()
	os.Exit(rc)
}

func TestInitLogger(t *testing.T) {
	if logger == nil {
		t.Fatal("expected logger to never be nil")
	}
	logger.Infof("Logger Initialized")
}

// InitLogger - initilize zap logger
func InitTestLogger() *zap.SugaredLogger {
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{"stdout"}
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	corelogger, _ := config.Build()

	return corelogger.Sugar()
}
