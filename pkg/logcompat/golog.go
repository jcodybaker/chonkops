package logcompat

import (
	golog "log"
	"regexp"
	"sync"
	"time"

	"github.com/rs/zerolog"
)

var mdnsLog = regexp.MustCompile(`^(\d\d\d\d/\d\d/\d\d \d\d:\d\d:\d\d) \[(INFO|ERR)\] (\w+): (.*)`)

var goLogFlagsWarning sync.Once

func (l *LogWriter) Write(b []byte) (int, error) {
	if golog.Flags() != golog.LstdFlags {
		goLogFlagsWarning.Do(func() {
			l.log.Warn().Msg("non-standard go log flags mean stdlib go logging will not be parsed to zerolog")
		})
		return l.log.Write(b)
	}
	if matches := mdnsLog.FindSubmatch(b); len(matches) == 5 {
		golog.Flags()
		var e *zerolog.Event
		switch string(matches[2]) {
		case "INFO":
			e = l.log.Info()
		case "ERR":
			e = l.log.Error()
		default:
			e = l.log.WithLevel(zerolog.NoLevel)
		}
		if t, err := time.Parse("2006/01/02 15:04:05.999999", string(matches[1])); err == nil {
			e = e.Time(zerolog.TimestampFieldName, t)
		} else if t, err := time.Parse("2006/01/02 15:04:05", string(matches[1])); err == nil {
			e = e.Time(zerolog.TimestampFieldName, t)
		}
		if len(matches[3]) > 0 {
			e = e.Str("component", string(matches[3]))
		}
		e.Msg(string(matches[4]))
		return len(b), nil
	}
	return l.log.Write(b)
}
