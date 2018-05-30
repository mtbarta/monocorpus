package util

import (
	"os"

	"github.com/go-kit/kit/log"
)

func CreateNewLogFmtLogger() log.Logger {
	klogger := log.NewLogfmtLogger(os.Stdout)
	klogger = log.With(klogger, "ts", log.DefaultTimestampUTC)
	klogger = log.With(klogger, "caller", log.DefaultCaller)

	return klogger
}

func CreateNewJSONLogger() log.Logger {
	klogger := log.NewJSONLogger(os.Stdout)
	klogger = log.With(klogger, "ts", log.DefaultTimestampUTC)
	klogger = log.With(klogger, "caller", log.DefaultCaller)

	return klogger
}
