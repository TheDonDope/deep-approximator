package logs

import (
	"log"

	"go.uber.org/zap"
)

// Logger provides the application logger
var Logger = newLogger()

// NewLogger returns a new Logger
func newLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()
	return logger
}
