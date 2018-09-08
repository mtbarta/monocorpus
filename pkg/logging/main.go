package logging

import (
	"log"

	zap "go.uber.org/zap"
)

type MCLogger struct {
	*zap.SugaredLogger
	version string
}

/*
 * NewProductionLogger creates a production logger
 */
func NewProductionLogger() *zap.SugaredLogger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	sugar := logger.Sugar()
	defer sugar.Sync()

	return sugar
}
