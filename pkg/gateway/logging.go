package gateway

import (
	"os"

	"github.com/go-kit/kit/log"
)

var logger = newLogger(log.NewJSONLogger(os.Stdout))

var klogger = newLogger(log.NewLogfmtLogger(os.Stdout))

func newLogger(init log.Logger) log.Logger {
	logger := init
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	return logger
}
