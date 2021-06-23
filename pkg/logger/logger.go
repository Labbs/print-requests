package logger

import (
	"os"

	"github.com/rs/zerolog"
)

// Logger - logger object
var Logger zerolog.Logger

// InitLogger - configure the logger with automatic fields
func InitLogger(version string) {
	host, _ := os.Hostname()
	Logger = zerolog.
		New(os.Stderr).
		With().
		Caller().
		Timestamp().
		Str("host", host).
		Str("version", version).
		Logger()
}
