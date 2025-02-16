package provider

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"

	"github.com/alexpts/edu-go/pkg/zerolog/transport"
)

var MicroSec = "2006-01-02T15:04:05.999999Z07:00"

type zeroLogLevelHook struct{}

func (h zeroLogLevelHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	value := -1

	// https://en.wikipedia.org/wiki/Syslog#Severity_level
	switch level {
	case zerolog.NoLevel:
		return
	case zerolog.TraceLevel:
		value = 8
	case zerolog.DebugLevel:
		value = 7
	case zerolog.InfoLevel:
		value = 6
	case zerolog.WarnLevel:
		value = 4
	case zerolog.ErrorLevel:
		value = 3
	case zerolog.FatalLevel:
		value = 2
	case zerolog.PanicLevel:
		value = 0
	default:
		return
	}

	e.Int("sl_level", value)
}

func ProvideZeroLogger() *zerolog.Logger {
	zerolog.TimeFieldFormat = MicroSec
	zerolog.SetGlobalLevel(zerolog.TraceLevel)
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack // error must implement `StackTrace() errors.StackTrace`

	//zerolog.LevelFieldName = "" // skip std string level
	zerolog.TimestampFieldName = "ts"
	zerolog.MessageFieldName = "msg"
	zerolog.ErrorStackFieldName = "_stack"
	zerolog.ErrorFieldName = "err"

	zerolog.CallerSkipFrameCount = 2

	consoleDest := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: MicroSec,
	}

	udpDest, err := transport.UdpWriter("127.0.0.1:12201")
	if err != nil {
		fmt.Println(err)
	}

	dest := zerolog.MultiLevelWriter(
		consoleDest,
		udpDest,
		os.Stdout,
	)

	logger := zerolog.New(dest).
		With().
		Timestamp().
		Caller(). // @todo cut prefix path to git root
		Logger()

	//logger = logger.Hook(zeroLogLevelHook{})
	return &logger
}
