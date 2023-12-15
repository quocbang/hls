package logging

import (
	"log"

	"go.uber.org/zap"
)

func init() {
	var (
		logger *zap.Logger
		err    error
	)

	logger, err = zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to initialize logger, err: ")
	}

	zap.ReplaceGlobals(logger)
	zap.RedirectStdLog(logger)
}
