package log

import (
	"buynow.api/config"
	"go.uber.org/zap"
)

const prodEnv = "prod"

func MustInitLogger(cfg *config.Config) *zap.Logger {
	if cfg.App.Environment != prodEnv {
		logger, err := zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
		return logger
	}

	return zap.NewNop()
}
