package log

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
	"os"
)

var Module = fx.Options(
	fx.Invoke(InitLogger),
)

// InitLogger configures logger.
func InitLogger() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}
