package logger

import (
	"file-cleaner/internal/lib/tracinghook"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var Logger zerolog.Logger

func init() {
	Logger = zerolog.New(os.Stdout).With().
		Timestamp().Logger().Hook(tracinghook.TracingHook{})

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// Defined configured logger as default
	log.Logger = Logger
}

func Info() *zerolog.Event {
	return Logger.Info()
}

func Warn() *zerolog.Event {
	return Logger.Warn()
}

func Error() *zerolog.Event {
	return Logger.Error()
}

func Debug() *zerolog.Event {
	return Logger.Debug()
}

func Fatal() *zerolog.Event {
	return Logger.Fatal()
}

func Panic() *zerolog.Event {
	return Logger.Panic()
}
