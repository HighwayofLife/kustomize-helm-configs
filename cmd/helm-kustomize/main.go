package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
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

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "download",
				Aliases: []string{"d"},
				Usage:   "Download helm charts from manifest file",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:    "template",
				Aliases: []string{"t"},
				Usage:   "Run helm template on downloaded charts",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println("Example")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.Fatalw("Failed to run app", "error", err.Error())
	}
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
