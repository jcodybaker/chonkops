package logcompat

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var DefaultLogWriter *LogWriter = &LogWriter{
	log: &log.Logger,
}

type LogWriter struct {
	log *zerolog.Logger
}
